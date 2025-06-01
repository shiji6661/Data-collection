package handler

import (
	"Api/client"
	"article_srv/proto_article/article"
	"context"
)

// 文章分类添加
func CateAdd(ctx context.Context, i *article.CateAddReq) (*article.CateAddResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		add, err := in.CateAdd(ctx, i)
		if err != nil {
			return nil, err
		}
		return add, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.CateAddResp), nil
}

// 文章和内容发表
func PostArticle(ctx context.Context, i *article.PostArticleReq) (*article.PostArticleResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		add, err := in.PostArticle(ctx, i)
		if err != nil {
			return nil, err
		}
		return add, nil
	})
	return articleClient.(*article.PostArticleResp), err
}

// 文章内容修改
func EditArticle(ctx context.Context, i *article.UpdateArticleReq) (*article.UpdateArticleResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		updateArticle, err := in.UpdateArticle(ctx, i)
		if err != nil {
			return nil, err
		}
		return updateArticle, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.UpdateArticleResp), nil
}

// 文章内容删除
func Deleted(ctx context.Context, i *article.DeleteArticleReq) (*article.DeleteArticleResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		deleteArticle, err := in.DeleteArticle(ctx, i)
		if err != nil {
			return nil, err
		}
		return deleteArticle, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.DeleteArticleResp), nil
}

// 文章分类查询
func CateList(ctx context.Context, i *article.CateListReq) (*article.CateListResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		list, err := in.CateList(ctx, i)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.CateListResp), nil
}

// 文章评论
func ArtComment(ctx context.Context, i *article.ArtCommentReq) (*article.ArtCommentResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		artComment, err := in.ArtComment(ctx, i)
		if err != nil {
			return nil, err
		}
		return artComment, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.ArtCommentResp), nil
}

// 评论回复
func ReplyComment(ctx context.Context, i *article.ReplyCommentReq) (*article.ReplyCommentResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		replyComment, err := in.ReplyComment(ctx, i)
		if err != nil {
			return nil, err
		}
		return replyComment, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.ReplyCommentResp), nil
}

// 删除评论
func DeletedComment(ctx context.Context, i *article.DelCommentReq) (*article.DelCommentResp, error) {
	articleClient, err := client.ArticleClient(ctx, func(ctx context.Context, in article.ArticleClient) (interface{}, error) {
		comment, err := in.DelComment(ctx, i)
		if err != nil {
			return nil, err
		}
		return comment, nil
	})
	if err != nil {
		return nil, err
	}
	return articleClient.(*article.DelCommentResp), nil
}
