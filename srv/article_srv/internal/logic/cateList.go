package logic

import (
	"article_srv/proto_article/article"
	"common/global"
	"errors"
	"models/model_article/model_mongodb"
)

func CateList(in *article.CateListReq) (*article.CateListResp, error) {
	category, err := model_mongodb.FindArticleCategory(global.Database, global.CollectionName)
	if err != nil {
		return nil, errors.New("查询mongodb分类表失败")
	}
	var cateList []*article.CateList
	for _, articleCategory := range category {
		lists := article.CateList{
			Title: articleCategory.Title,
			Intr:  articleCategory.Intr,
			Img:   articleCategory.Image,
			Pid:   int64(articleCategory.Pid),
			Id:    int64(articleCategory.Id),
		}
		cateList = append(cateList, &lists)
	}
	return &article.CateListResp{List: cateList}, nil
}
