package controlerglobal

import (
	"fmt"
	"pangan-segar/databases"
	modelglobal "pangan-segar/model/global"
	"pangan-segar/utils"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var input modelglobal.InputRegister
	var user modelglobal.Users

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Check if the phone number already exists
	err := databases.DB.Table("users").Where("phone = ?", input.Phone).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		// Log internal errors for debugging purposes
		fmt.Println("Error querying user:", err)
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err == nil {
		// Phone number is already registered
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Phone sudah terdaftar.",
		})
		return
	}

	user.Name = input.Name
	user.Phone = input.Phone
	user.Active = false
	user.Role = input.Role
	user.Coin = 0
	user.Profile = nil

	if err := databases.DB.Table("users").Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create user account. Please try again or contact support if the issue persists.",
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

	errCreateOtp := databases.DB.Table("otps").Create(&otp).Error
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

	// Successfully created user and sent OTP
	ctx.JSON(200, gin.H{
		"error":       false,
		"message":     "User account created successfully, OTP sent.",
		"data":        user,
		"otpResponse": result, // Include the OTP response in the response data
	})
}

func UpdatePinUser(ctx *gin.Context) {
	var input modelglobal.InputPinUser
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
			"message": "Databes User Error",
		})
		return
	}

	user.Pin = utils.HashPassword(input.Pin)
	if err := databases.DB.Table("users").Save(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databes user sedang error",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update data pin user success",
		"data":    user,
	})
}

func LoginUser(ctx *gin.Context) {
	var input modelglobal.InputLogi
	var user modelglobal.Users
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   false,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("users").Where("phone = ?", input.Phone).First(&user).Error; err != nil {
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

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "login success",
		"data":    user,
	})
}

func OtpLupaPin(ctx *gin.Context) {
	var input modelglobal.InputLogi
	var user modelglobal.Users
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   false,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("users").Where("phone = ?", input.Phone).First(&user).Error; err != nil {
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
		"message": "login success",
		"data":    user,
		"result":  result,
	})
}

func LoginPin(ctx *gin.Context) {
	var input modelglobal.InputPinUser
	var user modelglobal.Users
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   false,
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
			"message": "Databes User Error",
		})
		return
	}

	err := utils.VerifikasiHashPassword(input.Pin, user.Pin)
	if err != nil {
		ctx.JSON(401, gin.H{ // Status 401 for Unauthorized
			"error":   true,
			"message": "Invalid pin",
		})
		return
	}

	token, err := utils.SignToken(input.UserId, user.Phone, string(user.Role))
	if err != nil {
		ctx.JSON(500, gin.H{ // Status 500 for Internal Server Error
			"error":   true,
			"message": "Failed to generate token.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "confirm pin success",
		"token":   token,
	})
}
