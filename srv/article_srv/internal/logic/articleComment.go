package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"errors"
	"models/model_article/model_mongodb"
	"models/model_article/model_mysql"
	"time"
)

func ArtComment(in *article.ArtCommentReq) (*article.ArtCommentResp, error) {
	find, _ := model_mongodb.FindArticleById(global.Database, global.Database, int(in.ArticleId))

	if find.Id == 0 {
		return nil, errors.New("该文章不存在")
	}
	am := model_mysql.ArticleComment{
		ArticleId: int32(in.ArticleId),
		Username:  in.Username,
		Content:   in.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := am.CreateComment()
	if err != nil {
		return nil, errors.New("评论失败")
	}
	err = model_mongodb.InsertArticleComment(global.Database, global.ArticleComment, am)
	if err != nil {
		return nil, errors.New("评论失败")
	}
	return &article.ArtCommentResp{Greet: "评论成功"}, nil
}
