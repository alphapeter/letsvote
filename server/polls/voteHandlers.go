package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
)

func HandleVote(c *gin.Context) {
	user := c.MustGet("user").(users.User)
	v := Vote{}
	pollId := c.Params.ByName("pollId")

	//todo assert pollId and poll in correct status

	err := config.DB.First(&v, "user_id = ? and poll_id = ?", user.Id, pollId).Error
	if err!= nil && err != gorm.ErrRecordNotFound {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	update := VoteDto{}
	c.ShouldBindJSON(&update)

	v.PollId = pollId
	v.UserId = user.Id

	//todo assert options are for the correct poll
	v.Score1OptionId = NullString(update.Score1OptionId)
	v.Score2OptionId = NullString(update.Score2OptionId)
	v.Score3OptionId = NullString(update.Score3OptionId)
	config.DB.Save(&v)

	c.JSON(http.StatusOK, struct {
		Success bool
	}{Success:true})
}