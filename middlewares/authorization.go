package middlewares

import (
	"my-gram/database"
	"my-gram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoPostAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Select("user_id").First(&User, uint(userID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "User Not Found",
				"message": "User doesn't exist",
			})
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to post data",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
func CommentPostAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err := db.Select("user_id").First(&Comment, uint(userID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Comment Not Found",
				"message": "Comment doesn't exist",
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to comment this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
func SocialMediaPostAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err := db.Select("user_id").First(&SocialMedia, uint(userID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Social media Not Found",
				"message": "Social media doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err = db.Select("user_id").First(&User, uint(userId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
