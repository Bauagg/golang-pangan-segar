package controllerkonsumen

import (
	"fmt"
	"os"
	"pangan-segar/config"
	"pangan-segar/databases"
	modelglobal "pangan-segar/model/global"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileKonsumen(ctx *gin.Context) {
	var user modelglobal.Users
	user_id, _ := ctx.Get("userID")

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

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "data Profile user success",
		"data":    user,
	})
}

func UpdateProfileKonsumen(ctx *gin.Context) {
	var input modelglobal.InputUpdateProfile
	var user modelglobal.Users
	user_id, _ := ctx.Get("userID")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
	}

	file, _ := ctx.FormFile("profile")

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

	if input.Email != nil {

		regexEmaill := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !regexEmaill.MatchString(*input.Email) {
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid email format",
			})

			return
		}

		if user.Email == nil || *user.Email != *input.Email {
			err := databases.DB.Table("users").Where("email = ?", input.Email).First(&user).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Internal server error",
				})
				return
			}

			if err == nil {
				// Jika tidak ada error dan user ditemukan, email sudah terdaftar
				ctx.JSON(400, gin.H{
					"error":   true,
					"message": "Email sudah terdaftar.",
				})
				return
			}

			user.Email = input.Email
		} else if user.Email == input.Email {
			user.Email = input.Email
		}
	}

	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Phone != "" {
		if user.Phone != input.Phone {
			err := databases.DB.Table("users").Where("phone = ?", input.Phone).First(&user).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Internal server error",
				})
				return
			}

			if err == nil {
				// Jika tidak ada error dan user ditemukan, email sudah terdaftar
				ctx.JSON(400, gin.H{
					"error":   true,
					"message": "Phone sudah terdaftar.",
				})
				return
			}

			user.Phone = input.Phone
		} else if user.Phone == input.Phone {
			user.Phone = input.Phone
		}
	}

	if file != nil {
		imageDir := "./public/profile-user"

		if user.Profile != nil {
			fileName := filepath.Base(*user.Profile)
			oldFilePath := filepath.Join(imageDir, fileName)

			if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Failed to delete old image: " + err.Error(),
				})
				return
			}
		}

		// Create directory if not exists
		if _, err := os.Stat(imageDir); os.IsNotExist(err) {
			err = os.MkdirAll(imageDir, os.ModePerm)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Failed to create image directory: " + err.Error(),
				})
				return
			}
		}

		// Save new profile image
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := filepath.Join(imageDir, fileName)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to save image: " + err.Error(),
			})
			return
		}

		// Update profile field with new image URL
		urlHost := config.URL_HOST_SERVER + "/profile-user/" + fileName
		user.Profile = &urlHost
	}

	if err := databases.DB.Table("users").Save(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Databes User Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update profile User success",
		"data":    user,
	})
}
