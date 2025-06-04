package client

import (
	"article_srv/proto_article/article"
	"context"
	"google.golang.org/grpc"

)

type HandlerArticle func(ctx context.Context, in article.ArticleClient) (interface{}, error)

func ArticleClient(ctx context.Context, handler func(ctx context.Context, in article.ArticleClient) (interface{}, error)) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8005", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := article.NewArticleClient(dial)
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
