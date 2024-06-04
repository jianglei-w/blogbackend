package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jianglei-w/blogbackend/controller"
	"github.com/jianglei-w/blogbackend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)
	//Restful style API
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/post", controller.AllPost)
	app.Get("/api/post/:id", controller.DetailPost)
	app.Put("/api/post/:id", controller.UpdatePost)
	app.Delete("/api/post/:id", controller.DeletePost)

	app.Get("/api/uniquepost", controller.UniquePost)

	app.Post("/api/upload", controller.Upload)
	app.Static("/api/uploads", "./uploads")

}
