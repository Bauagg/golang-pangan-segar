package controlerglobal

import (
	"fmt"
	"math/rand"
	"pangan-segar/databases"
	modelglobal "pangan-segar/model/global"
	"pangan-segar/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VerifyOtp(ctx *gin.Context) {
	var input modelglobal.InputOtp
	var otp modelglobal.Otps
	var user modelglobal.Users

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	if err := databases.DB.Table("users").Where("id = ?", input.UserId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid User not Found.",
			})
			return

		}
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "An unexpected error occurred while retrieving user data. Please try again later or contact support if the issue persists.",
		})
		return
	}

	if err := databases.DB.Table("otps").Where("number_otp = ? AND user_id = ? AND expires_at > ?", input.NumberOtp, input.UserId, time.Now()).First(&otp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "The OTP code you entered is not found. Please check if it is correct or request a new OTP.",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "An unexpected error occurred while verifying the OTP. Please try again later.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "success Verifikasi OTP",
	})
}

func SendOtpPhone(ctx *gin.Context) {
	var user modelglobal.Users

	if err := databases.DB.Table("users").Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid User not Found.",
			})
			return

		}
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Databes User Error",
		})
		return
	}

	// Generate a random 4-digit OTP
	rand.Seed(time.Now().UnixNano())
	randomOTP := rand.Intn(9000) + 1000

	// Save OTP to the database
	otp := modelglobal.Otps{
		NumberOtp: uint64(randomOTP),
		UserId:    uint64(user.ID),
		ExpiresAt: time.Now().Add(1 * time.Minute), // OTP expires in 1 minutes
	}

	errCreateOtp := databases.DB.Table("otps").Where("user_id = ?", user.ID).Updates(&otp).Error
	if errCreateOtp != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to generate OTP.",
		})
		return
	}

	result, status, err := utils.SendPhoneOtp(user.Phone, uint64(randomOTP))
	if err != nil {
		fmt.Println("Failed to send OTP:", err)
		ctx.JSON(status, gin.H{
			"error":   true,
			"message": "Failed to send OTP. Please try again later.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "send OTP success",
		"result":  result,
	})
}
