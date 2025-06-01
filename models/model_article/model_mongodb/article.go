package model_mongodb

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"models/model_article/model_mysql"
	"time"
)

// 分类添加
func Create(dateBase, collectionName string, doc interface{}) error {
	if global.Client == nil {
		return fmt.Errorf("MongoDB client is not initialized")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := global.Client.Database(dateBase).Collection(collectionName)
	_, err := collection.InsertOne(ctx, doc)
	return err
}

// 查询文章分类表的类型id
func FindArticleCategoryPid(dateBase, collectionName string, pid int) (model_mysql.ArticleCategory, error) {
	var date model_mysql.ArticleCategory

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	collection := global.Client.Database(dateBase).Collection(collectionName)

	fist := bson.D{{"pid", pid}}

	err := collection.FindOne(ctx, fist).Decode(&date)

	if err != nil {
		return model_mysql.ArticleCategory{}, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model_mysql.ArticleCategory{}, nil
	} else if err != nil {
		return model_mysql.ArticleCategory{}, nil
	}

	return date, err
}

// 编辑文章
func EditArticle(dateBase, collectionName string, id int, date model_mysql.Article) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	// 获取集合
	collection := global.Client.Database(dateBase).Collection(collectionName)

	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", date}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil
	}
	return nil

}

// 编辑文章内容
func EditArticleContent(dateBase, collectionName string, id int, date model_mysql.ArticleContent) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	// 获取集合
	collection := global.Client.Database(dateBase).Collection(collectionName)

	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", date}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil
	}
	return nil

}

// 查询文章管理的id
func FindArticleById(dateBase, collectionName string, id int) (model_mysql.Article, error) {
	var date model_mysql.Article

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	collection := global.Client.Database(dateBase).Collection(collectionName)

	fist := bson.D{{"id", id}}

	err := collection.FindOne(ctx, fist).Decode(&date)

	if err != nil {
		return model_mysql.Article{}, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model_mysql.Article{}, nil
	} else if err != nil {
		return model_mysql.Article{}, nil
	}
	return date, err
}

// 删除文章
func DeleteArticle(dateBase, collectionName string, id int, deletedAt time.Time) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	coll := global.Client.Database(dateBase).Collection(collectionName)

	// 构造查询条件
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{
		"DeletedAt", deletedAt}}}}
	// 执行删除操作
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete articles: %w", err)
	}

	// 检查是否有文档被删除
	if result.MatchedCount == 0 {
		return fmt.Errorf("no documents found with cid %d", id)
	}
	return nil
}

// 删除文章内容
func DeleteArticleContent(dateBase, collectionName string, id int, deletedAt time.Time) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	coll := global.Client.Database(dateBase).Collection(collectionName)

	// 构造查询条件
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"DeletedAt", deletedAt}}}}
	// 执行删除操作
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete articles: %w", err)
	}

	// 检查是否有文档被删除
	if result.MatchedCount == 0 {
		return fmt.Errorf("no documents found with cid %d", id)
	}
	return nil
}

// 查询文章管理列表
func FindArticleCategory(dateBase, collectionName string) ([]model_mysql.ArticleCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := global.Client.Database(dateBase).Collection(collectionName)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%w", err)

	}
	// 文章分类表
	defer cur.Close(ctx)

	var res []model_mysql.ArticleCategory

	for cur.Next(ctx) {

		var result model_mysql.ArticleCategory

		err = cur.Decode(&result)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		res = append(res, result)
	}

	err = cur.Err()

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return res, nil
}

// 文章评论
func InsertArticleComment(dateBase, collectionName string, doc interface{}) error {
	if global.Client == nil {
		return fmt.Errorf("MongoDB client is not initialized")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := global.Client.Database(dateBase).Collection(collectionName)
	_, err := collection.InsertOne(ctx, doc)
	return err
}

// 查询评论的id
func FindCommentById(dateBase, collectionName string, id int) (model_mysql.ArticleComment, error) {
	var date model_mysql.ArticleComment

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	collection := global.Client.Database(dateBase).Collection(collectionName)

	fist := bson.D{{"id", id}}

	err := collection.FindOne(ctx, fist).Decode(&date)

	if err != nil {
		return model_mysql.ArticleComment{}, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model_mysql.ArticleComment{}, nil
	} else if err != nil {
		return model_mysql.ArticleComment{}, nil
	}
	return date, err
}

// 删除文章
func DeleteComment(dateBase, collectionName string, id int, deletedAt time.Time) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	coll := global.Client.Database(dateBase).Collection(collectionName)

	// 构造查询条件
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"deletedat", deletedAt}}}}
	// 执行删除操作
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete articles: %w", err)
	}

	// 检查是否有文档被删除
	if result.MatchedCount == 0 {
		return fmt.Errorf("no documents found with cid %d", id)
	}
	return nil
}
