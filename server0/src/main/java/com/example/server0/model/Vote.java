package com.example.server0.model;

import java.sql.Date;
import java.sql.Timestamp;

public class Vote {
    private Integer VoteId;
    private String VoteName;
    private String VoteResult;
    private String VoteType;
    private Date VoteDdl;

    public Integer getVoteId() {
        return VoteId;
    }

    public void setVoteId(Integer voteId) {
        VoteId = voteId;
    }

    public String getVoteName() {
        return VoteName;
    }

    public void setVoteName(String voteName) {
        VoteName = voteName;
    }

    public String getVoteResult() {
        return VoteResult;
    }

    public void setVoteResult(String voteResult) {
        VoteResult = voteResult;
    }

    public String getVoteType() {
        return VoteType;
    }

    public void setVoteType(String voteType) {
        VoteType = voteType;
    }

    public Date getVoteDdl() {
        return VoteDdl;
    }

    public void setVoteDdl(Date voteDdl) {
        VoteDdl = voteDdl;
    }
}