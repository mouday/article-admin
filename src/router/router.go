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
	app.POST("/api/renderArticle", api.RenderArticle)

	// category
	app.POST("/api/addCategory", api.AddCategory)
	app.POST("/api/updateCategory", api.UpdateCategory)
	app.POST("/api/deleteCategory", api.DeleteCategory)
	app.POST("/api/getCategory", api.GetCategory)
	app.POST("/api/getCategoryList", api.GetCategoryList)
	app.POST("/api/updateCategoryStatus", api.UpdateCategoryStatus)

	// tag
	app.POST("/api/addTag", api.AddTag)
	app.POST("/api/updateTag", api.UpdateTag)
	app.POST("/api/deleteTag", api.DeleteTag)
	app.POST("/api/getTag", api.GetTag)
	app.POST("/api/getTagPage", api.GetTagPage)
	app.POST("/api/updateTagStatus", api.UpdateTagStatus)
	app.POST("/api/updateCategoryOrderValue", api.UpdateCategoryOrderValue)

}
