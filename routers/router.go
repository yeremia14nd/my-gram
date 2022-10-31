package routers

import (
	"my-gram/controllers"
	"my-gram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), controllers.DeleteUser)
		//delete
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", middlewares.PhotoPostAuthorization(), controllers.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.GET("/", middlewares.PhotoPostAuthorization(), controllers.GetPhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", middlewares.CommentPostAuthorization(), controllers.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.GET("/", middlewares.CommentPostAuthorization(), controllers.GetComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", middlewares.SocialMediaPostAuthorization(), controllers.CreateSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.GET("/", middlewares.SocialMediaPostAuthorization(), controllers.GetSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
