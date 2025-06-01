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
func EditArticle(in *article.UpdateArticleReq) (*article.UpdateArticleResp, error) {
	find, _ := model_mongodb.FindArticleById(global.Database, "article", int(in.ArticleId))

	if find.Id == 0 {
		return nil, errors.New("该文章不存在")
	}

	pid, err := model_mongodb.FindArticleCategoryPid(global.Database, global.CollectionName, int(in.CId))
	if err != nil {
		return nil, errors.New("查询文章分类失败")
	}
	if pid.Id == 0 {
		return nil, errors.New("文章分类不存在")
	}
	a := model_mysql.Article{
		Cid:        int32(in.CId),
		Title:      in.Title,
		Author:     in.Username,
		ImageInput: in.ImageInput,
		UpdatedAt:  time.Now(),
	}
	a.Id = uint32(in.ArticleId)
	err = model_mongodb.EditArticle(global.Database, global.Article, int(in.ArticleId), a)
	if err != nil {
		return nil, errors.New("文章修改失败")
	}
	ac := model_mysql.ArticleContent{
		Nid:       a.Id,
		Content:   in.Content,
		UpdatedAt: time.Now(),
	}
	err = model_mongodb.EditArticleContent(global.Database, global.ArticleContent, int(a.Id), ac)
	if err != nil {
		return nil, errors.New("文章内容修改失败")
	}
	return &article.UpdateArticleResp{Greet: "文章修改成功"}, nil
}
