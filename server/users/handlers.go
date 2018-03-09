package users

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type updateUserAdminCommand struct {
	isAdmin bool `json:"is_admin"`
}

type userDto struct {
	userId  string `json:"user_id"`
	name    string `json:"name"`
	email   string `json:"email"`
	isAdmin bool   `json:"is_admin"`
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
			name:    u.Name,
			isAdmin: u.IsAdmin,
			userId:  u.Id,
			email:   u.Email,
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
	userId := c.GetString("userId")
	c.ShouldBindWith(&command, binding.JSON)

	for _, u := range users {
		if u.Id == userId {
			if err := config.DB.Model(&u).UpdateColumn("is_admin", command.isAdmin).Error; err != nil {
				c.JSON(http.StatusInternalServerError, "could not update user")
				return
			}
			c.JSON(http.StatusOK, command.isAdmin)
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
