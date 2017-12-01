package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"github.com/jinzhu/gorm"
)

func errorResponse(c *gin.Context, reason string, responseCode int) {
	c.JSON(responseCode, struct {
		Success bool   `json:"success"`
		Reason  string `json:"reason"`
	}{false, reason})
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

func HandleVote(c *gin.Context) {
	user := c.MustGet("user").(users.User)
	v := Vote{}
	pollId := c.Params.ByName("pollId")

	//todo assert pollId and poll in correct status

	err := config.DB.First(&v, "userId = ?, pollId = ?", user.Id, pollId).Error
	if err!= nil && err != gorm.ErrRecordNotFound {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	update := Vote{}
	c.ShouldBindJSON(&update)

	v.PollId = pollId
	v.UserId = user.Id

	//todo assert options are for the correct poll
	v.Score1OptionId = update.Score1OptionId
	v.Score2OptionId = update.Score2OptionId
	v.Score3OptionId = update.Score3OptionId
	config.DB.Save(&v)

	c.JSON(http.StatusOK, struct {
		Success bool
	}{Success:true})
}

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

func UpdateOption(c *gin.Context) {
	var p map[string]string
	if err := c.BindJSON(&p); err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	option, err:= fetchOptionForEdit(c, "optionId")
	if err != nil {
		errorResponse(c, err.message(), err.responseCode())
	}

	hasBeenUpdated := false
	if name, ok := p["name"]; ok {
		option.Name = name
		hasBeenUpdated = true
	}

	if description, ok := p["description"]; ok {
		option.Description = description
		hasBeenUpdated = true
	}
	if !hasBeenUpdated{
		errorResponse(c, "No valid fields for patch", http.StatusBadRequest)
	}

	if err := config.DB.Model(&option).Update(p).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Option    Option `json:"poll"`
	}{true, option})
	PollUpdated(option.Id)
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

func ActivatePoll (c *gin.Context){

}

func EndPoll (c *gin.Context){

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

func fetchOptionForEdit(c *gin.Context, idParameterName string) (Option, fetchError) {
	id := c.Param(idParameterName)
	user := c.MustGet("user").(users.User)

	var option Option
	if err := config.DB.First(&option, "id = ?", id).Error; err != nil {
		return option, unsuccessfulFetch{msg: err.Error(), code: http.StatusInternalServerError}
	}
	if (!hasPermissionToEdit(option, user)){
		return option, unautorizedFetch{}
	}
	return option, nil
}

func hasPermissionToEdit(created userCreated, user users.User) bool {
	return created.getUserId() == user.Id
}






type fetchError interface {
	message() string
	responseCode() int
}

type unautorizedFetch struct {

}

func (unautorizedFetch) message() string{
	return "unauthorized"
}
func (unautorizedFetch) responseCode() int{
	return http.StatusUnauthorized
}

type unsuccessfulFetch struct {
	msg string
	code int
}

func (f unsuccessfulFetch) message() string{
	return f.msg
}
func (f unsuccessfulFetch) responseCode() int{
	return f.code
}




