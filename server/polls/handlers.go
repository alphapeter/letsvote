package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
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

func UpdatePoll(c *gin.Context) {
	var p Poll
	if err := c.BindJSON(&p); err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	id := c.Param("id")
	user := c.MustGet("user").(users.User)

	p.Id = id

	var poll Poll
	if err := config.DB.First(&poll, "id = ?", id).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if poll.CreatedByUserId != user.Id {
		errorResponse(c, "Unauthorized", http.StatusUnauthorized)
		return
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

func DeletePoll(c *gin.Context) {
	id := c.Param("pollId")
	user := c.MustGet("user").(users.User)
	var poll Poll
	if err := config.DB.First(&poll, "id = ?", id).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if poll.CreatedByUserId != user.Id {
		errorResponse(c, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
	pollId := c.Param("id")

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
	user := c.MustGet("user").(users.User);

	optionId := c.Param("optionId")

	option, err := FetchOption(optionId)
	if err != nil {
		errorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if option.CreatedByUserId != user.Id {
		errorResponse(c, "Can only be delete by user whom created the option", http.StatusUnauthorized)
		return
	}
	pollId := option.PollId
	config.DB.Delete(&option, "id = ?", option.Id)
	c.JSON(http.StatusOK, struct {
		Success 		bool `json:"success"`
	}{true })
	OptionDeleted(pollId, optionId)
}


