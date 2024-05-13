package com.example.server0.model;

public class VoteVoter {
    private Integer voteVoterId;
    // 投票ID
    private Integer vote_Id;
    // 候选人ID
    private Integer voter_Id;

    // 候选人对象
    private Voter voter; // 添加候选人对象属性
    private String binaryString;

    public String getBinaryString() {
        return binaryString;
    }

    public void setBinaryString(String binaryString) {
        this.binaryString = binaryString;
    }

    public Integer getVoteVoterId() {
        return voteVoterId;
    }

    public void setVoteVoterId(Integer voteVoterId) {
        this.voteVoterId = voteVoterId;
    }

    public Integer getVote_Id() {
        return vote_Id;
    }

    public void setVote_Id(Integer vote_Id) {
        this.vote_Id = vote_Id;
    }

    public Integer getVoter_Id() {
        return voter_Id;
    }

    public void setVoter_Id(Integer voter_Id) {
        this.voter_Id = voter_Id;
    }

    public Voter getVoter() {
        return voter;
    }

    public void setVoter(Voter voter) {
        this.voter = voter;
    }
}
