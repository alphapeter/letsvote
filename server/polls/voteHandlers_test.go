package polls

import "testing"

func TestVoteOptionValidationAllOk(t *testing.T) {

	vote := VoteDto{Score1: "a", Score2: "b", Score3: "c"}
	options := []Option{
		{Id: "a"},
		{Id: "b"},
		{Id: "c"},
	}

	if !allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected true")
	}
}

func TestVoteOptionValidationOption1IsMissing(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "b", Score3: "c"}
	options := []Option{
		{Id: "b"},
		{Id: "c"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}

func TestVoteOptionValidationOption2IsMissing(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "b", Score3: "c"}
	options := []Option{
		{Id: "a"},
		{Id: "c"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}

func TestVoteOptionValidationOption3IsMissing(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "b", Score3: "c"}
	options := []Option{
		{Id: "a"},
		{Id: "b"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}

func TestVoteOptionValidationOptionsAreEmpty(t *testing.T) {
	vote := VoteDto{Score1: "", Score2: "", Score3: ""}
	options := []Option{
		{Id: "a"},
		{Id: "b"},
		{Id: "c"},
	}

	if !allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected true")
	}
}

func TestVoteOptionValidationOptions1IsEmpty(t *testing.T) {
	vote := VoteDto{Score1: "", Score2: "b", Score3: "c"}
	options := []Option{
		{Id: "a"},
		{Id: "c"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}

func TestVoteOptionValidationOptions1IsWrong(t *testing.T) {
	vote := VoteDto{Score1: "d", Score2: "b", Score3: "c"}
	options := []Option{
		{Id: "a"},
		{Id: "b"},
		{Id: "c"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}
func TestVoteOptionValidationOptions2IsWrong(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "d", Score3: "c"}
	options := []Option{
		{Id: "a"},
		{Id: "b"},
		{Id: "c"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}
func TestVoteOptionValidationOptions3IsWrong(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "b", Score3: "d"}
	options := []Option{
		{Id: "a"},
		{Id: "b"},
		{Id: "c"},
	}

	if allVoteOptionsInPoll(vote, options) {
		t.Error("Excpected false")
	}
}

func TestVotesAreUniqueOrEmpty_1_equal_2__False(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "a", Score3: "d"}
	if allVotesAreUniqueOrEmpty(vote) {
		t.Error("Excpected false")
	}
}
func TestVotesAreUniqueOrEmpty_2_equal_3__False(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "b", Score3: "b"}
	if allVotesAreUniqueOrEmpty(vote) {
		t.Error("Excpected false")
	}
}
func TestVotesAreUniqueOrEmpty_1_equal_3__False(t *testing.T) {
	vote := VoteDto{Score1: "a", Score2: "b", Score3: "a"}
	if allVotesAreUniqueOrEmpty(vote) {
		t.Error("Excpected false")
	}
}

func TestVotesAreUniqueOrEmpty_all_empty__True(t *testing.T) {
	vote := VoteDto{Score1: "", Score2: "", Score3: ""}
	if !allVotesAreUniqueOrEmpty(vote) {
		t.Error("Excpected true")
	}
}
