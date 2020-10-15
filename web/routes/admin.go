package routes

import "github.com/gofiber/fiber/v2"

// AddAdminRoutes : will add all the route under admin
func AddAdminRoutes(admin fiber.Router) {
	userRoute := admin.Group("/user")

	userRoute.Post("/", CreateNewUser)
	userRoute.Get("/", GetAllUsers)
	userRoute.Get("/:user_id", GetUserByID)
	userRoute.Put("/:user_id", UpdateUser)
	userRoute.Delete("/:user_id", DeleteUser)

}

// CreateNewUser : Based on the request payload user will be created
func CreateNewUser(c *fiber.Ctx) error {

	return c.SendStatus(200)
}

// GetAllUsers : Get All User in System
func GetAllUsers(c *fiber.Ctx) error {

	return c.SendStatus(200)
}

// GetUserByID : Get a User in System by its ID
func GetUserByID(c *fiber.Ctx) error {

	return c.SendStatus(200)
}

// UpdateUser : Update a User in System by its ID
func UpdateUser(c *fiber.Ctx) error {

	return c.SendStatus(200)
}

// DeleteUser : Delete a User in System by its ID
func DeleteUser(c *fiber.Ctx) error {

	return c.SendStatus(200)
}
