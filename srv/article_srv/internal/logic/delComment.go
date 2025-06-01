package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"errors"
	"models/model_article/model_mongodb"
	"time"
)

func DeleteComment(in *article.DelCommentReq) (*article.DelCommentResp, error) {
	id, err := model_mongodb.FindCommentById(global.Database, global.ArticleComment, int(in.CommentId))
	if err != nil {
		return nil, errors.New("该评论不存在")
	}
	if id.Id == 0 {
		return nil, errors.New("该评论不存在")
	}
	err = model_mongodb.DeleteComment(global.Database, global.ArticleComment, int(in.CommentId), time.Now())
	if err != nil {
		return nil, errors.New("评论删除失败")
	}
	return &article.DelCommentResp{Greet: "评论删除成功"}, nil
}
