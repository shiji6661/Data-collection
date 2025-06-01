package request

// 分类添加请求
type CateAddReq struct {
	Title string `json:"title" form:"title"`
	Intr  string `json:"intr" form:"intr"`
	Img   string `json:"img" form:"img"`
	Pid   int64  `json:"pid" form:"pid"`
}

// 文章发表
type PostArticleReq struct {
	Username   string `json:"username" form:"username"`
	CId        int64  `json:"cId" form:"cId"`
	ImageInput string `json:"imageInput" form:"imageInput"`
	Content    string `json:"content" form:"content"`
	Title      string `json:"title" form:"title"`
}

// 编辑文章
type EditArticleReq struct {
	Username   string `json:"username" form:"username"`
	CId        int    `json:"CId" form:"CId"`
	ImageInput string `json:"imageInput" form:"imageInput"`
	Content    string `json:"content" form:"content"`
	Title      string `json:"title" form:"title"`
	ArticleId  int    `json:"articleId" form:"articleId"`
}

// 删除文章
type DeletedReq struct {
	ArticleId int `json:"articleId" form:"articleId"`
}

// 文章评论
type ArticleCommentReq struct {
	ArticleId int    `json:"articleId" form:"articleId"`
	Username  string `json:"username" form:"username"`
	Content   string `json:"content" form:"content"`
}

// 评论回复
type ReplyCommentReq struct {
	ArticleId int    `json:"articleId" form:"articleId"`
	Username  string `json:"username" form:"username"`
	Content   string `json:"content" form:"content"`
	Pid       int    `json:"pid" form:"pid"`
	CommentId int    `json:"commentId" form:"commentId"`
}

// 评论删除
type DeleteCommentReq struct {
	CommentId int `json:"commentId" form:"commentId"`
}
