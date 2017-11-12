package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

func AddPoll(c *gin.Context) {
	var p Poll

	if err := c.BindJSON(&p); err != nil {
		c.String(http.StatusInternalServerError, "internal server error, could not parse json")
		return
	}

	p.CreatedByUserId = "peter@stratsys.se"
	p.Id = uuid.NewV4().String()
	err := config.DB.Create(&p).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Success bool   `json:"success"`
			Reason  string `json:"reason"`
		}{false, err.Error()})
		return
	}
	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Poll    Poll `json:"poll"`
	}{true, p})
}

func UpdatePoll(c *gin.Context) {
	var p Poll

	if err := c.BindJSON(&p); err != nil {
		c.String(http.StatusInternalServerError, "internal server error, could not parse json")
		return
	}

	id := c.Param("id")
	p.Id = id

	err := config.DB.Save(p).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Success bool   `json:"success"`
			Reason  string `json:"reason"`
		}{false, err.Error()})
		return
	}
	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Poll    Poll `json:"poll"`
	}{true, p})
}

func GetPolls(c *gin.Context) {
	var polls []Poll
	config.DB.Find(&polls)
	c.JSON(http.StatusOK, polls)
}
func GetPoll(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func GetOptions(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func GetOption(c *gin.Context) {
	c.String(http.StatusOK, "")
}
func AddOption(c *gin.Context) {
	c.String(http.StatusOK, "")
}
