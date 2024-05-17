package dto

import "test/model"

type VoteDto struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
	Num      int    `json:"num"`
	Count    int    `json:"count"`
	State    int    `json:"state"`
}

func ToVoteDto(vote model.Vote, participate int) VoteDto {
	return VoteDto{
		ID:       vote.ID,
		Name:     vote.Name,
		Deadline: vote.Deadline,
		Num:      vote.Num,
		Count:    participate,
		State:    vote.State,
	}
}
