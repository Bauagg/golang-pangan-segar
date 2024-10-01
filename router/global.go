package router

import (
	controlerglobal "pangan-segar/controller/global"
	"pangan-segar/middleware"

	"github.com/gin-gonic/gin"
)

func RouterGlobal(app *gin.Engine) {
	router := app

	router.POST("/api/pangan-segar/v-1/register", controlerglobal.Register)
	router.POST("/api/pangan-segar/v-1/login", controlerglobal.LoginUser)

	// Router pin
	router.POST("/api/pangan-segar/v-1/lupa-pin", controlerglobal.OtpLupaPin)
	router.PUT("/api/pangan-segar/v-1/pin", controlerglobal.UpdatePinUser)
	router.POST("/api/pangan-segar/v-1/pin", controlerglobal.LoginPin)

	// Router OTP
	router.GET("/api/pangan-segar/v-1/otp/:id", controlerglobal.SendOtpPhone)
	router.POST("/api/pangan-segar/v-1/otp", controlerglobal.VerifyOtp)

	// Router Address
	router.GET("/api/pangan-segar/v-1/address", middleware.AuthMiddleware(), controlerglobal.ListAllAddress)
	router.GET("/api/pangan-segar/v-1/address/:id", middleware.AuthMiddleware(), controlerglobal.DetailAllAddress)
	router.POST("/api/pangan-segar/v-1/address", middleware.AuthMiddleware(), controlerglobal.CreateAddress)
}
