package users

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type updateUserAdminCommand struct {
	IsAdmin bool `json:"is_admin"`
}

type userDto struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func GetUsers(c *gin.Context) {
	user := c.MustGet("user").(User)

	var users []User
	config.DB.Find(&users)

	if !hasAdminPermission(user, users) {
		c.JSON(http.StatusForbidden, "access denied")
	}

	var userDtos []userDto
	for _, u := range users {
		userDtos = append(userDtos, userDto{
			Name:    u.Name,
			IsAdmin: u.IsAdmin,
			Id:      u.Id,
			Email:   u.Email,
		})
	}
	c.JSON(http.StatusOK, userDtos)
}

func SetAdminPermission(c *gin.Context) {
	user := c.MustGet("user").(User)

	var users []User
	config.DB.Find(&users)

	if !hasAdminPermission(user, users) {
		c.JSON(http.StatusForbidden, "access denied")
	}

	command := updateUserAdminCommand{}
	userId := c.Params.ByName("userId")
	c.ShouldBindWith(&command, binding.JSON)

	for _, u := range users {
		if u.Id == userId {
			if err := config.DB.Model(&u).UpdateColumn("is_admin", command.IsAdmin).Error; err != nil {
				c.JSON(http.StatusInternalServerError, "could not update user")
				return
			}
			c.JSON(http.StatusOK, command.IsAdmin)
			return
		}
	}

	c.JSON(http.StatusBadRequest, "user not found")
}

func hasAdminPermission(user User, users []User) bool {
	if user.IsAdmin {
		return true
	}
	for _, u := range users {
		if u.IsAdmin {
			return false
		}
	}
	return true
}
