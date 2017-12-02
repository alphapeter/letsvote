package polls

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/satori/go.uuid"
	"github.com/alphapeter/letsvote/server/config"
)

func UpdateOption(c *gin.Context) {
	var o map[string]string
	if err := c.BindJSON(&o); err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	option, err:= fetchOptionForEdit(c, "optionId")
	if err != nil {
		errorResponse(c, err.message(), err.responseCode())
	}

	updated := map[string]interface{}{}
	if name, ok := o["name"]; ok {
		option.Name = name
		updated["name"] = name
	}

	if description, ok := o["description"]; ok {
		option.Description = description
		updated["description"] = description
	}
	if len(updated) == 0{
		errorResponse(c, "No valid fields for patch", http.StatusBadRequest)
	}

	if err := config.DB.Model(&option).Update(o).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	updated["option_id"] = option.Id
	updated["poll_id"] = option.PollId

	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Option    Option `json:"poll"`
	}{true, option})
	OptionUpdated(updated)
}

func AddOption(c *gin.Context) {
	var o Option

	if err := c.BindJSON(&o); err != nil {
		errorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	user := c.MustGet("user").(users.User)
	pollId := c.Param("pollId")

	var poll Poll
	if err := config.DB.First(&poll, "id = ?", pollId).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	o.PollId = poll.Id
	o.CreatedByUserId = user.Id
	o.Id = uuid.NewV1().String()
	err := config.DB.Create(&o).Error

	if err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Option    Option `json:"option"`
	}{true, o})
	OptionCreated(o.Id)
}
func DeleteOption(c *gin.Context){

	option, err := fetchOptionForEdit(c, "optionId")
	if err != nil {
		errorResponse(c, err.message(), err.responseCode())
		return
	}

	optionId := option.Id

	pollId := option.PollId
	config.DB.Delete(&option, "id = ?", option.Id)
	c.JSON(http.StatusOK, struct {
		Success 		bool `json:"success"`
	}{true })
	OptionDeleted(pollId, optionId)
}



func fetchOptionForEdit(c *gin.Context, idParameterName string) (Option, fetchError) {
	id := c.Param(idParameterName)
	user := c.MustGet("user").(users.User)

	var option Option
	if err := config.DB.First(&option, "id = ?", id).Error; err != nil {
		return option, unsuccessfulFetch{msg: err.Error(), code: http.StatusInternalServerError}
	}
	if !hasPermissionToEdit(option, user) {
		return option, unautorizedFetch{}
	}
	return option, nil
}