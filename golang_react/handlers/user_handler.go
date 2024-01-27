package handlers

import (
	"super_admin/config"
	"super_admin/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(ctx *fiber.Ctx) error{
	var users []models.User
	config.DB.Find(&users)
	ctx.Status(200).JSON(users)
	return nil
}


func DeleteUser(ctx *fiber.Ctx) error{
	userId := ctx.Params("userId")
	
	var user models.User
	config.DB.Where("id=?", userId).Delete(&user)
	
	ctx.Status(200).JSON(user)
	return nil
	
}


func GetAllClients(ctx *fiber.Ctx) error {
	var clients []models.Client
	config.DB.Find(&clients)
	ctx.JSON(clients)
	return nil
	
}


func DeleteClient(ctx *fiber.Ctx) error{
	clientId := ctx.Params("clientId")
	
	var client models.Client
	config.DB.Where("id=?", clientId).Delete(&client)

	var user models.User
	config.DB.Where("client_id=?", clientId).Delete(&user)
	
	ctx.JSON("awerawer")
	return nil
	
}