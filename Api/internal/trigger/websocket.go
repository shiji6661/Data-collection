package trigger

import (
	"Api/global"
	"Api/internal/handler"
	"Api/internal/request"
	"Api/internal/response"
	"collection_srv/proto_collection/collection"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

// var mutex *sync.RWMutex
var OnlineUser = make(map[int64]*websocket.Conn)

// todo:此函数是 WebSocket 连接的入口，负责将 HTTP 连接升级为 WebSocket 连接，并且处理接收到的消息。
func Echo(conn *gin.Context) {
	userId := conn.GetUint("userId")
	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(conn.Writer, conn.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var req request.WsRequest
		err = json.Unmarshal(message, &req)
		if err != nil {
			response.WsResponseError(c, 1001, "非法入参")
		}

		switch req.Cmd {
		case "online":
			go OnlineFunc(c, message, conn, int64(userId))
		case "send":
			go SendFunc(c, message, conn, int64(userId))
		case "allSend":
			go AllSendFunc(c, message, conn, int64(userId))
		default:
			response.WsResponseError(c, 10000, "没有该方法")
		}
	}
}

// todo:此函数用于向所有在线用户发送消息。
func AllSendFunc(c *websocket.Conn, message []byte, conn *gin.Context, uid int64) {
	var req request.AllSendRequest
	err := json.Unmarshal(message, &req)
	if err != nil {
		response.WsResponseError(c, 10001, "allSend解析失败")
		return
	}

	msg := req.Data.Message
	for userId, _ := range OnlineUser {
		if err = response.WsResponseSuccess(OnlineUser[userId], msg); err != nil {
			fmt.Printf("向用户 %d 发送消息失败: %v", userId, err)
			// 发送失败时可以选择断开连接
			//m.removeConnection(userId)
		} else {
			//TODO:广播信息入MongoDB逻辑
			_, err = handler.InformationStore(conn, &collection.InformationStoreRequest{
				Uid:       uid,
				Tid:       0,
				Heartbeat: msg,
				Database:  global.DATABASE,
				Table:     global.BROADCAST,
			})
			if err != nil {
				response.WsResponseError(c, 10001, "广播消息入库失败")
				return
			} else {
				err = response.WsResponseSuccess(c, "发送成功")
				if err != nil {
					return
				}
			}
		}
	}
}

// todo:此函数用于向指定用户发送消息。
func SendFunc(c *websocket.Conn, message []byte, conn *gin.Context, uid int64) {
	var req request.SendRequest
	err := json.Unmarshal(message, &req)
	if err != nil {
		response.WsResponseError(c, 10001, "send解析失败")
		return
	}

	toUserId := req.Data.ToUserId
	msg := req.Data.Message

	if toUserId == 0 {
		response.WsResponseError(c, 100001, "目标用户ID不能为空")
		return
	}
	fmt.Println("2222", toUserId)
	//mutex.RLock()
	con := OnlineUser[toUserId]
	//mutex.RUnlock()
	if con == nil {
		//TODO:用户未上线时，消息存入redis，待用户上线后再发送
		_, err = handler.MessageCache(conn, &collection.MessageCacheRequest{
			Uid:       uid,
			Tid:       toUserId,
			Heartbeat: msg,
		})
		if err != nil {
			response.WsResponseError(c, 100004, err.Error())
			return
		}
		fmt.Println("3333")
		//response.WsResponseError(c, 100003, "消息已存入缓存，待用户上线后再发送")
		//response.WsResponseError(c, 100002, "目标用户未上线")
		//return
	} else {
		response.WsResponseSuccess(OnlineUser[uid], msg)
		//TODO:消息入MongoDB逻辑
		_, err = handler.InformationStore(conn, &collection.InformationStoreRequest{
			Uid:       uid,
			Tid:       toUserId,
			Heartbeat: msg,
			Database:  global.DATABASE,
			Table:     global.PRIVATE,
		})
		if err != nil {
			response.WsResponseError(c, 10001, err.Error())
			return
		} else {
			err = response.WsResponseSuccess(c, "入库成功")
			if err != nil {
				return
			}
		}
		// 只在成功时通知发送者
		err = response.WsResponseSuccess(c, "发送成功")
		if err != nil {
			return
		}

	}
}

// todo:此函数用于处理用户上线请求。
func OnlineFunc(c *websocket.Conn, message []byte, conn *gin.Context, uid int64) {
	var req request.OnlineRequest
	err := json.Unmarshal(message, &req)
	if err != nil {
		response.WsResponseError(c, 100001, "上线失败")
	}
	// 先检查用户是否已在线，如果是则断开之前的连接
	//mutex.Lock()
	if existingConn, exists := OnlineUser[uid]; exists {
		fmt.Printf("用户 %d 重复登录，断开之前的连接", uid)
		existingConn.Close()
		delete(OnlineUser, uid)
	}

	OnlineUser[uid] = c
	fmt.Printf("用户 %d 已添加到在线用户列表", uid)
	//mutex.Unlock()

	err = response.WsResponseSuccess(c, "上线成功")
	if err != nil {
		fmt.Printf("通知用户 %d 上线成功失败: %v", uid, err)
		return
	} else {
		fmt.Printf("通知用户 %d 上线成功成功", uid)
		//TODO:查看是否有未读消息，有则发送
		cache, err := handler.GetMessageCache(conn, &collection.GetMessageCacheRequest{Uid: uid})
		if err != nil {
			response.WsResponseError(c, 10001, err.Error())
			return
		} else {
			for _, m := range cache.List {
				fmt.Printf("用户 %d 有未读消息: %s", cache.Uid, m.Heartbeat)
				if err = response.WsResponseSuccess(OnlineUser[cache.Uid], m.Heartbeat); err != nil {
					fmt.Printf("向用户 %d 发送消息失败: %v", cache.Uid, err)
				}
			}
		}
	}

	//TODO:上线信息入MongoDB逻辑,可以用来实现用户实时状态查看
	fmt.Printf("用户 %d 上线成功，当前在线用户数: %d", uid, len(OnlineUser))
}
