package polls

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/satori/go.uuid"
)

func UpdatePoll(c *gin.Context) {
	var p map[string]string
	if err := c.BindJSON(&p); err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	poll, err:= fetchPollForEdit(c, "pollId")
	if err != nil {
		errorResponse(c, err.message(), err.responseCode())
		return
	}

	hasBeenUpdated := false
	if name, ok := p["name"]; ok {
		poll.Name = name
		hasBeenUpdated = true
	}

	if description, ok := p["description"]; ok {
		poll.Description = description
		hasBeenUpdated = true
	}
	if input, ok := p["status"]; ok {
		i, err := strconv.Atoi(input)
		if err != nil {
			errorResponse(c, err.Error(), http.StatusBadRequest)
			return
		}

		var status = Status(i)
		poll.Status = status
		hasBeenUpdated = true
	}

	if !hasBeenUpdated{
		errorResponse(c, "No valid fields for patch", http.StatusBadRequest)
	}

	if err := config.DB.Model(&poll).Update(p).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Poll    Poll `json:"poll"`
	}{true, poll})
	PollUpdated(poll.Id)
}

func AddPoll(c *gin.Context) {
	var p Poll

	if err := c.BindJSON(&p); err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	user := c.MustGet("user").(users.User)

	p.CreatedByUserId = user.Id
	p.CreatedBy = user
	p.Id = uuid.NewV4().String()
	err := config.DB.Create(&p).Error

	if err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Poll    Poll `json:"poll"`
	}{true, p})
	PollCreated(p.Id)
}

func DeletePoll(c *gin.Context) {
	poll, err := fetchPollForEdit(c, "pollId")
	if err != nil {
		errorResponse(c, err.message(), err.responseCode())
		return
	}
	id := poll.Id

	if err := config.DB.Delete(&poll, "Id = ?", poll.Id).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Success 		bool `json:"success"`
	}{true })
	PollDeleted(id)
}

func GetPolls(c *gin.Context) {
	var polls []Poll
	config.DB.Preload("Options").
		Preload("CreatedBy").
		Preload("Winner").
		Preload("Votes").
		Preload("Options.CreatedBy").
		Find(&polls)
	c.JSON(http.StatusOK, polls)
}

func fetchPollForEdit(c *gin.Context, idParameterName string) (Poll, fetchError){
	id := c.Param(idParameterName)
	user := c.MustGet("user").(users.User)

	var poll Poll
	if err := config.DB.First(&poll, "id = ?", id).Error; err != nil {
		return poll, unsuccessfulFetch{msg: err.Error(), code: http.StatusInternalServerError}
	}
	if (!hasPermissionToEdit(poll, user)){
		return poll, unautorizedFetch{}
	}
	return poll, nil
}
