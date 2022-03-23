package main

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/routes"
	"blog/api/service"
	"blog/infrastructure"
	"blog/models"
	mainRouter "blog/routes"
)

func main() {

	infrastructure.LoadEnv()                                    //loading env
	router := mainRouter.NewGinRouter()                         //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // databse has been initialized and configured
	postRepository := repository.NewPostRepository(db)          // repository are being setup
	postService := service.NewPostService(postRepository)       // service are being setup
	postController := controller.NewPostController(postService) // controller are being set up
	postRoute := routes.NewPostRoute(postController, router)    // post routes are initialized
	postRoute.Setup()                                           // post routes are being setup

	db.DB.AutoMigrate(&models.Post{}) // migrating Post model to datbase table
	router.Gin.Run(":8000")           //server started on 8000 port

}
