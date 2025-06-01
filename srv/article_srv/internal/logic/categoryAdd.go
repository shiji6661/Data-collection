package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"errors"
	"models/model_article/model_mongodb"
	"models/model_article/model_mysql"
	"time"
)

// 分类添加
func CategoryAdd(in *article.CateAddReq) (*article.CateAddResp, error) {
	ac := model_mysql.ArticleCategory{
		Title:   in.Title,
		Intr:    in.Intr,
		Image:   in.Img,
		Pid:     int32(in.Pid),
		AddTime: time.Now(),
	}
	err := ac.CreateCategory()
	if err != nil {
		return nil, errors.New("文章分类添加失败")
	}
	err = model_mongodb.Create(global.Database, global.CollectionName, ac)
	if err != nil {
		return nil, errors.New("文章分类添加失败")
	}
	return &article.CateAddResp{Greet: "文章分类添加成功"}, nil
}
