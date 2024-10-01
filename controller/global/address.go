package controlerglobal

import (
	"pangan-segar/databases"
	modelglobal "pangan-segar/model/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAllAddress(ctx *gin.Context) {
	var address []modelglobal.Address
	user_id, _ := ctx.Get("userID")

	if err := databases.DB.Table("addresses").Where("user_id = ?", user_id).Find(&address).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Address Error",
		})
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list data address success",
		"data":    address,
	})
}

func DetailAllAddress(ctx *gin.Context) {
	var address modelglobal.Address
	user_id, _ := ctx.Get("userID")

	if err := databases.DB.Table("addresses").Where("user_id = ? id = ?", user_id, ctx.Param("id")).First(&address).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Address Error",
		})
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list data address success",
		"data":    address,
	})
}

func CreateAddress(ctx *gin.Context) {
	var input modelglobal.Address
	var user modelglobal.Users
	user_id, _ := ctx.Get("userID")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	if err := databases.DB.Table("users").Where("id = ?", user_id).First(&user).Error; err != nil {
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

	input.UserId = uint64(user.ID)

	var validateAddress modelglobal.Address
	if err := databases.DB.Table("addresses").Where("user_id = ? AND status = ?", user.ID, "utama").First(&validateAddress).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			status := "utama"
			input.Status = &status
		}
	}

	if err := databases.DB.Table("addresses").Create(&input).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Address Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create data success",
		"data":    input,
	})
}
