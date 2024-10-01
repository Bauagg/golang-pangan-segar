package router

import (
	controllerkonsumen "pangan-segar/controller/konsumen"
	"pangan-segar/middleware"

	"github.com/gin-gonic/gin"
)

func RouterKonsumen(app *gin.Engine) {
	router := app

	router.GET("/api/pangan-segar/v-1/profile", middleware.AuthMiddleware(), controllerkonsumen.ProfileKonsumen)
	router.PUT("/api/pangan-segar/v-1/profile", middleware.AuthMiddleware(), controllerkonsumen.UpdateProfileKonsumen)
}
