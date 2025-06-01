package trigger

import (
	"Api/internal/handler"
	"Api/internal/request"
	"Api/internal/response"
	"article_srv/proto_article/article"
	"github.com/gin-gonic/gin"
)

// 分类添加
func CateAdd(c *gin.Context) {
	var data request.CateAddReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	add, err := handler.CateAdd(c, &article.CateAddReq{
		Title: data.Title,
		Intr:  data.Intr,
		Img:   data.Img,
		Pid:   data.Pid,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, add)
}

// 文章发布
func PostArticle(c *gin.Context) {
	var data request.PostArticleReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	add, err := handler.PostArticle(c, &article.PostArticleReq{
		Username:   data.Username,
		CId:        data.CId,
		ImageInput: data.ImageInput,
		Content:    data.Content,
		Title:      data.Title,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, add)
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var data request.EditArticleReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	editArticle, err := handler.EditArticle(c, &article.UpdateArticleReq{
		Username:   data.Username,
		CId:        int64(data.CId),
		ImageInput: data.ImageInput,
		Content:    data.Content,
		Title:      data.Title,
		ArticleId:  int64(data.ArticleId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, editArticle)
}

// 文章删除
func Delete(c *gin.Context) {
	var data request.DeletedReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	deleted, err := handler.Deleted(c, &article.DeleteArticleReq{
		ArticleId: int64(data.ArticleId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, deleted)
}

// 分类列表
func CateList(c *gin.Context) {
	list, err := handler.CateList(c, &article.CateListReq{})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, list)
}

// 文章评论
func ArtComment(c *gin.Context) {
	var data request.ArticleCommentReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	articleComment, err := handler.ArtComment(c, &article.ArtCommentReq{
		Username:  data.Username,
		ArticleId: int64(data.ArticleId),
		Content:   data.Content,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, articleComment)
}

// 评论回复
func ReplyComment(c *gin.Context) {
	var data request.ReplyCommentReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	replyComment, err := handler.ReplyComment(c, &article.ReplyCommentReq{
		Username:  data.Username,
		ArticleId: int64(data.ArticleId),
		Content:   data.Content,
		Pid:       int64(data.Pid),
		CommentId: int64(data.CommentId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, replyComment)
}

// 删除评论
func DeleteComment(c *gin.Context) {
	var data request.DeleteCommentReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	comment, err := handler.DeletedComment(c, &article.DelCommentReq{
		CommentId: int64(data.CommentId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, comment)
}
