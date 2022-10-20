package routers

import (
	"final-project/common/database"
	socialController "final-project/social/controller"
	userController "final-project/user/controller"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	db := database.ConnectDB()
	router := gin.Default()
	user := userController.NewUserController(db)
	social := socialController.NewSocialController(db)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/login", user.UserLogin)
		userGroup.POST("/register", user.CreateUser)
	}

	socialGroup := router.Group("/socials")
	{
		socialGroup.GET("/", social.CreateSocial)
		//itemGroup.GET("/:id", itemController.FindItemById)
		//itemGroup.POST("/", orderController.CreateOrder)
		//itemGroup.PUT("/:id", itemController.UpdateItem)
		//itemGroup.DELETE("/:id", itemController.DeleteItem)
	}
	//
	//orderGroup := router.Group("/orders")
	//{
	//	orderGroup.POST("/", orderController.CreateOrder)
	//	orderGroup.GET("/", orderController.FindOrders)
	//	orderGroup.GET("/:id", orderController.FindOrderById)
	//	orderGroup.PUT("/:id", orderController.UpdateOrder)
	//	orderGroup.DELETE("/:id", orderController.DeleteOrder)
	//}
	//
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
