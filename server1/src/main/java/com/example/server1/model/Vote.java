package com.example.server1.model;

import java.sql.Date;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;

public class Vote {
    private Integer voteId;
    private String voteName;
    private String voteResult;
    private String voteType;
    private Date voteDdl;

    public Integer getVoteId() {
        return voteId;
    }

    public void setVoteId(Integer voteId) {
        this.voteId = voteId;
    }

    public String getVoteName() {
        return voteName;
    }

    public void setVoteName(String voteName) {
        this.voteName = voteName;
    }

    public String getVoteResult() {
        return voteResult;
    }

    public void setVoteResult(String voteResult) {
        this.voteResult = voteResult;
    }

    public String getVoteType() {
        return voteType;
    }

    public void setVoteType(String voteType) {
        this.voteType = voteType;
    }

    public Date getVoteDdl() {
        return voteDdl;
    }

    public void setVoteDdl(Date voteDdl) {
        this.voteDdl = voteDdl;
    }


}