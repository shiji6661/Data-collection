package server

import (
	"article_srv/internal/logic"
	"article_srv/proto_article/article"
	"context"
)

type ServerArticle struct {
	article.UnimplementedArticleServer
}

func (s ServerArticle) CateAdd(ctx context.Context, in *article.CateAddReq) (*article.CateAddResp, error) {
	add, err := logic.CategoryAdd(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}

func (s ServerArticle) PostArticle(ctx context.Context, in *article.PostArticleReq) (*article.PostArticleResp, error) {
	postArticle, err := logic.PostArticle(in)
	if err != nil {
		return nil, err
	}
	return postArticle, nil
}

func (s ServerArticle) UpdateArticle(ctx context.Context, in *article.UpdateArticleReq) (*article.UpdateArticleResp, error) {
	editArticle, err := logic.EditArticle(in)
	if err != nil {
		return nil, err
	}
	return editArticle, nil
}

func (s ServerArticle) DeleteArticle(ctx context.Context, in *article.DeleteArticleReq) (*article.DeleteArticleResp, error) {
	resp, err := logic.Delete(in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s ServerArticle) CateList(ctx context.Context, in *article.CateListReq) (*article.CateListResp, error) {
	list, err := logic.CateList(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (s ServerArticle) ArtComment(ctx context.Context, in *article.ArtCommentReq) (*article.ArtCommentResp, error) {
	comment, err := logic.ArtComment(in)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (s ServerArticle) ReplyComment(ctx context.Context, in *article.ReplyCommentReq) (*article.ReplyCommentResp, error) {
	reply, err := logic.Reply(in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func (s ServerArticle) DelComment(ctx context.Context, in *article.DelCommentReq) (*article.DelCommentResp, error) {
	comment, err := logic.DeleteComment(in)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
