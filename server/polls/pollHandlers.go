package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

func UpdatePoll(c *gin.Context) {
	var p map[string]string
	if err := c.BindJSON(&p); err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	poll, err := fetchPollForEdit(c, "pollId")
	if err != nil {
		errorResponse(c, err.message(), err.responseCode())
		return
	}

	updated := map[string]interface{}{}
	shouldCountScore := false
	if name, ok := p["name"]; ok {
		poll.Name = name
		updated["name"] = name
	}

	if description, ok := p["description"]; ok {
		poll.Description = description
		updated["description"] = description
	}
	if input, ok := p["status"]; ok {
		i, err := strconv.Atoi(input)
		if err != nil {
			errorResponse(c, err.Error(), http.StatusBadRequest)
			return
		}
		status := REGISTRATING
		if i >= 10 {
			status = ENDED
		} else if i >= 8 {
			status = COUNTING
			shouldCountScore = true

		} else if i >= 5 {
			status = VOTING
		}
		poll.Status = status
		updated["status"] = status
	}

	if len(updated) == 0 {
		errorResponse(c, "No valid fields for patch", http.StatusBadRequest)
	}

	if err := config.DB.Model(&poll).Update(p).Error; err != nil {
		errorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if shouldCountScore {
		go countScore(poll.Id)
	}
	c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
		Poll    Poll `json:"poll"`
	}{true, poll})

	updated["id"] = poll.Id

	PollUpdated(updated)
}

func countScore(pollId string) {
	time.Sleep(200)
	poll, err := FetchPoll(pollId)
	if err != nil {
		ScoreCountFailed(pollId, "Could not fetch poll "+pollId+" "+err.Error())
		return
	}

	var votes []Vote
	if err = config.DB.Find(&votes, "poll_id = ?", pollId).Error; err != nil {
		ScoreCountFailed(pollId, "Could not fetch votes for pollId: "+pollId+" "+err.Error())
		return
	}

	var updates []map[string]interface{}
	var affectedOptions []Option
	for _, o := range poll.Options {
		score := getScore(votes, o.Id)
		if score != o.Score {
			o.Score = score

			update := map[string]interface{}{
				"option_id": o.Id,
				"poll_id":   poll.Id,
				"score":     score,
			}
			updates = append(updates, update)
			affectedOptions = append(affectedOptions, o)
		}
	}

	for _, o := range affectedOptions {
		if err = config.DB.Model(&o).UpdateColumn("score", o.Score).Error; err != nil {
			ScoreCountFailed(pollId, "Could save options for pollId: "+pollId+" "+err.Error())
			return
		}
	}

	poll.Status = ENDED
	if err = config.DB.Model(&poll).UpdateColumn("status", poll.Status).Error; err != nil {
		ScoreCountFailed(pollId, "Could not update poll status for pollId: "+pollId+" "+err.Error())
		return
	}

	PollUpdated(map[string]interface{}{
		"status": poll.Status,
		"id":     poll.Id,
	})
	for _, update := range updates {
		OptionUpdated(update)
	}
}

func getScore(votes []Vote, optionId string) int {
	score := 0
	for _, vote := range votes {
		if vote.Score1OptionId.Valid && vote.Score1OptionId.String == optionId {
			score += 1
		}
		if vote.Score2OptionId.Valid && vote.Score2OptionId.String == optionId {
			score += 2
		}
		if vote.Score3OptionId.Valid && vote.Score3OptionId.String == optionId {
			score += 3
		}
	}
	return score
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
	id := uuid.NewV4()
	p.Id = id.String()
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
		Success bool `json:"success"`
	}{true})
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

func fetchPollForEdit(c *gin.Context, idParameterName string) (Poll, fetchError) {
	id := c.Param(idParameterName)
	user := c.MustGet("user").(users.User)

	var poll Poll
	if err := config.DB.First(&poll, "id = ?", id).Error; err != nil {
		return poll, unsuccessfulFetch{msg: err.Error(), code: http.StatusInternalServerError}
	}
	if !hasPermissionToEdit(poll, user) {
		return poll, unautorizedFetch{}
	}
	return poll, nil
}
