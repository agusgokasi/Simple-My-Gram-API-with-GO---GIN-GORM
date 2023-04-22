package route

import (
	"project-final/controller"
	"project-final/middleware"
	"project-final/service"

	_ "project-final/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           MyGram API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
// @SecurityDefinitions bearerAuth
// @in header
// @name Authorization
// @TokenUrl /users/login
// @Scheme bearer
func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	controllers := controller.NewHttpServer(app)
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:photoId", controllers.GetPhotoByID)
		photoRouter.PUT("/:photoId", middleware.PhotoAuth(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuth(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/photos/:photoId/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetAllComment)
	}

	commentBaseRouter := r.Group("/comments")
	{
		commentBaseRouter.Use(middleware.Authentication())
		commentBaseRouter.GET("/:commentId", controllers.GetCommentByID)
		commentBaseRouter.PUT("/:commentId", middleware.CommentAuth(), controllers.UpdateComment)
		commentBaseRouter.DELETE("/:commentId", middleware.CommentAuth(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", controllers.GetSocialMediaByID)
		socialMediaRouter.PUT("", middleware.SocialMediaAuth(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("", middleware.SocialMediaAuth(), controllers.DeleteSocialMedia)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
	))
}
