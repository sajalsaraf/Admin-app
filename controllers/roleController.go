package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/database"
	"github.com/sajalsaraf/Admin-app.git/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

type RoleCreateDTO struct {
	name        string
	permissions []string
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}
	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.DB.Preload("Permissions").Find(&role)
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	// var result interface{}

	// database.DB.Table("role_permissions").Where("role_id = ?", uint(id)).Delete(&result)
	database.DB.Exec("DELETE FROM role_permissions where role_id = ?", id)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role) // this will update in "roles" table(using &role) with the value of role(last in Updates(role))
	return c.JSON(role)
}
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.DB.Delete(&role)
	return nil
}
