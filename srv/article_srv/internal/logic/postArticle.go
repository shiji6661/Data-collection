package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"context"
	"errors"
	"models/model_article/model_mongodb"
	"models/model_article/model_mysql"
	"time"
)

// 发布文章和内容
func PostArticle(in *article.PostArticleReq) (*article.PostArticleResp, error) {
	a := model_mysql.Article{
		Cid:        int32(in.CId),
		Title:      in.Title,
		Author:     in.Username,
		ImageInput: in.ImageInput,
		AddTime:    time.Now(),
		UpdatedAt:  time.Now(),
	}
	pid, err := model_mongodb.FindArticleCategoryPid(global.Database, global.CollectionName, int(in.CId))
	if err != nil {
		return nil, errors.New("查询文章分类失败")
	}
	if pid.Id == 0 {
		return nil, errors.New("文章分类不存在")
	}
	err = a.CreateArticle()
	if err != nil {
		return nil, errors.New("文章发表失败")
	}
	an := model_mysql.ArticleContent{
		Nid:         a.Id,
		Content:     in.Content,
		CreatedTime: time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = an.CreateArticleContent()
	if err != nil {
		return nil, errors.New("文章内容添加失败")
	}
	session, err := global.Client.StartSession()
	if err != nil {
		return nil, errors.New("启动 MongoDB 会话失败")
	}
	defer session.EndSession(context.TODO())

	err = session.StartTransaction()
	if err != nil {
		return nil, errors.New("启动 MongoDB 事务失败")
	}

	err = model_mongodb.Create(global.Database, global.Article, a)
	if err != nil {
		session.AbortTransaction(context.TODO())
		return nil, errors.New("文章发表失败mongodb")
	}
	err = model_mongodb.Create(global.Database, global.ArticleContent, an)
	if err != nil {
		session.AbortTransaction(context.TODO())
		return nil, errors.New("文章内容添加失败mongodb")
	}
	err = session.CommitTransaction(context.TODO())
	if err != nil {
		return nil, errors.New("提交 MongoDB 事务失败")
	}
	return &article.PostArticleResp{Greet: "文章和内容发表成功"}, nil
}
