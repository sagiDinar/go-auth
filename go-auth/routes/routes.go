package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post(path:"/api/register", controllers.Register)
	app.Post(path:"/api/login", controllers.Login)
	app.Get(path:"/api/User", controllers.User)
}
