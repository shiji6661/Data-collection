package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"errors"
	"models/model_article/model_mongodb"
	"models/model_article/model_mysql"
	"time"
)

func Reply(in *article.ReplyCommentReq) (*article.ReplyCommentResp, error) {
	id, err := model_mongodb.FindCommentById(global.Database, global.ArticleComment, int(in.CommentId))
	if err != nil {
		return nil, err
	}
	if id.Id == 0 {
		return nil, errors.New("该评论不存在")
	}
	am := model_mysql.ArticleComment{
		ArticleId: int32(in.ArticleId),
		Username:  in.Username,
		Content:   in.Content,
		Pid:       int32(in.Pid),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = am.CreateComment()
	if err != nil {
		return nil, errors.New("回复失败")
	}
	err = model_mongodb.InsertArticleComment(global.Database, global.ArticleComment, am)
	if err != nil {
		return nil, errors.New("回复失败")
	}
	return &article.ReplyCommentResp{Greet: "回复成功"}, nil
}
