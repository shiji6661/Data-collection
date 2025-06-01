package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"errors"
	"models/model_article/model_mongodb"
	"models/model_article/model_mysql"
	"time"
)

// 删除文章和内容
func Delete(in *article.DeleteArticleReq) (*article.DeleteArticleResp, error) {
	find, _ := model_mongodb.FindArticleById(global.Database, "article", int(in.ArticleId))

	if find.Id == 0 {
		return nil, errors.New("该文章不存在")
	}

	a := model_mysql.Article{
		Id:        uint32(in.ArticleId),
		DeletedAt: time.Now(),
	}
	a.Id = uint32(in.ArticleId)
	err := model_mongodb.DeleteArticle(global.Database, global.Article, int(a.Id), a.DeletedAt)
	if err != nil {
		return nil, errors.New("删除文章失败")
	}
	ac := model_mysql.ArticleContent{
		Id:        a.Id,
		DeletedAt: time.Now(),
	}
	err = model_mongodb.DeleteArticleContent(global.Database, global.ArticleContent, int(ac.Id), ac.DeletedAt)
	if err != nil {
		return nil, errors.New("删除文章内容失败")
	}
	return &article.DeleteArticleResp{Greet: "删除文章和内容成功"}, nil
}
