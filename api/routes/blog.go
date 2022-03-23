package routes

import (
	"blog/api/controller"
	"blog/routes"
)

//PostRoute -> Route for question module
type PostRoute struct {
	Controller controller.PostController
	Handler    routes.GinRouter
}

//NewPostRoute -> initializes new choice rouets
func NewPostRoute(controller controller.PostController, handler routes.GinRouter) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p PostRoute) Setup() {
	post := p.Handler.Gin.Group("/posts") //Router group
	{
		post.GET("/", p.Controller.GetPosts)
		post.POST("/", p.Controller.AddPost)
		post.GET("/:id", p.Controller.GetPost)
		post.DELETE("/:id", p.Controller.DeletePost)
		post.PUT("/:id", p.Controller.UpdatePost)
	}
}
