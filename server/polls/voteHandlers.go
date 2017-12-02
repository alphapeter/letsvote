package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetVotes(c *gin.Context) {
	user := c.MustGet("user").(users.User)
	var votes []Vote
	config.DB.Find(&votes).Where("user_id = ?", user.Id)
	var dtos []VoteDto
	for _, v := range votes {
		dtos = append(dtos, VoteDto{
			PollId: v.PollId,
			Score1: v.Score1OptionId.String,
			Score2: v.Score2OptionId.String,
			Score3: v.Score3OptionId.String,
		})
	}
	c.JSON(http.StatusOK, dtos)
}

func HandleVote(c *gin.Context) {
	user := c.MustGet("user").(users.User)
	v := Vote{}
	pollId := c.Params.ByName("pollId")

	var poll Poll
	err := config.DB.
		Preload("Options").
		First(&poll, "id = ?", pollId).Error

	if err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	update := VoteDto{}
	c.ShouldBindJSON(&update)

	if poll.Status > VOTING {
		errorResponse(c, "Votes cannot be added to a poll that has ended", http.StatusBadRequest)
		return
	}
	if poll.Status < VOTING {
		errorResponse(c, "Votes cannot be added to a poll that has not started", http.StatusBadRequest)
		return
	}
	if !allVotesAreUniqueOrEmpty(update) {
		errorResponse(c, "Bad state, there cannot be two votes for the same option!", http.StatusBadRequest)
		return
	}

	if !allVoteOptionsInPoll(update, poll.Options) {
		errorResponse(c, "Bad state, option voted for is not included in the poll", http.StatusBadRequest)
		return
	}

	err = config.DB.First(&v, "user_id = ? and poll_id = ?", user.Id, pollId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	v.PollId = pollId
	v.UserId = user.Id

	v.Score1OptionId = NullString(update.Score1)
	v.Score2OptionId = NullString(update.Score2)
	v.Score3OptionId = NullString(update.Score3)
	config.DB.Save(&v)

	UserVoted(user.Id, pollId)
	c.JSON(http.StatusOK, struct {
		Success bool
	}{Success: true})
}

func allVoteOptionsInPoll(v VoteDto, options []Option) bool {
	return (len(v.Score1) == 0 || optionExists(options, v.Score1)) &&
		(len(v.Score2) == 0 || optionExists(options, v.Score2)) &&
		(len(v.Score3) == 0 || optionExists(options, v.Score3))
}

func allVotesAreUniqueOrEmpty(v VoteDto) bool {
	return (len(v.Score1) == 0 || v.Score1 != v.Score2 && v.Score1 != v.Score3) &&
		(len(v.Score3) == 0 || v.Score2 != v.Score3)
}

func optionExists(options []Option, optionId string) bool {
	for _, o := range options {
		if o.Id == optionId {
			return true
		}
	}
	return false
}
