package server

func Run(app *fiber.App, port string) error {
	return app.Listen(port)
}
