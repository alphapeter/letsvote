package polls

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/tap"
)

func PollCreated(id string) {
	poll, _ := FetchPoll(id)
	tap.Broadcast("POLL_CREATED", poll)
}

func PollUpdated(update map[string]interface{}) {
	tap.Broadcast("POLL_UPDATED", update)
}

func OptionUpdated(update map[string]interface{}) {
	tap.Broadcast("OPTION_UPDATED", update)
}

func PollDeleted(id string) {
	tap.Broadcast("POLL_DELETED", id)
}

func OptionCreated(id string) {
	option, _ := FetchOption(id)
	tap.Broadcast("OPTION_CREATED", option)
}
func OptionDeleted(pollId string, optionId string) {
	payload := struct {
		OptionId string `json:"option_id"`
		PollId   string `json:"poll_id"`
	}{PollId: pollId, OptionId: optionId}
	tap.Broadcast("OPTION_DELETED", payload)
}

func ScoreCountFailed(pollId string, error string) {
	payload := struct {
		PollId string `json:"poll_id"`
		Error  string `json:"message"`
	}{PollId: pollId, Error: error}
	tap.Broadcast("POLL_SCORECOUNTFAILED", payload)
}

func UserVoted(userId string, pollId string) {
	vote := struct {
		UserId string `json:"user_id"`
		PollId string `json:"poll_id"`
	}{userId, pollId}
	tap.Broadcast("USER_VOTED", vote)
}

func FetchPoll(id string) (Poll, error) {
	var poll Poll
	err := config.DB.
		Preload("Options").
		Preload("CreatedBy").
		Preload("Winner").
		Preload("Votes").
		Preload("Options.CreatedBy").
		First(&poll, "id = ?", id).Error
	return poll, err
}

func FetchOption(id string) (Option, error) {
	var option Option
	err := config.DB.
		Preload("CreatedBy").
		First(&option, "id = ?", id).Error
	return option, err
}
