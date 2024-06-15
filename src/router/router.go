package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/article-admin/src/api"
)

/* 注册路由 */
func RegistRouter(app *gin.Engine) {
	// auth
	app.POST("/api/login", api.Login)

	// article
	app.POST("/api/addArticle", api.AddArticle)
	app.POST("/api/updateArticle", api.UpdateArticle)
	app.POST("/api/deleteArticle", api.DeleteArticle)
	app.POST("/api/getArticle", api.GetArticle)
	app.POST("/api/getArticlePage", api.GetArticlePage)
	app.POST("/api/updateArticleStatus", api.UpdateArticleStatus)

}
