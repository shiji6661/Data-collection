package router

import (
	"Api/internal/middleware"
	"Api/internal/trigger"
	"github.com/gin-gonic/gin"
)

func LoadArticle(r *gin.Engine) {
	r.Use(middleware.Logger())
	user := r.Group("/article")
	{
		user.POST("/cate/add", trigger.CateAdd)           // todo:分类添加
		user.POST("/post", trigger.PostArticle)           // todo:文章和内容发表
		user.POST("/edit", trigger.EditArticle)           // todo:文章修改
		user.POST("/delete", trigger.Delete)              // todo:文章删除
		user.GET("/cate/list", trigger.CateList)          // todo:文章分类列表
		user.POST("/comment", trigger.ArtComment)         // todo:文章评论
		user.POST("/reply/comment", trigger.ReplyComment) // todo: 评论回复
		user.POST("/del/comment", trigger.DeleteComment)  // todo: 评论删除
	}
}
