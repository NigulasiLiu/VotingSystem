package com.example.server0.model;

import java.util.ArrayList;
import java.util.List;

// 假设 VoteCandidate 类表示投票和候选人之间的关系
// 导入必要的包

public class VoteCandidate {
    // 投票候选人的ID
    private Integer voteCandidateId;
    // 投票ID
    private Integer vote_Id;
    // 候选人ID
    private Integer candidate_Id;
    // 候选人对象
    private Candidate candidate; // 添加候选人对象属性

    // 获取投票候选人的ID
    public Integer getVoteCandidateId() {
        return voteCandidateId;
    }

    // 设置投票候选人的ID
    public void setVoteCandidateId(Integer voteCandidateId) {
        this.voteCandidateId = voteCandidateId;
    }

    // 获取投票ID
    public Integer getVote_Id() {
        return vote_Id;
    }

    // 设置投票ID
    public void setVote_Id(Integer vote_Id) {
        this.vote_Id = vote_Id;
    }

    // 获取候选人ID
    public Integer getCandidate_Id() {
        return candidate_Id;
    }

    // 设置候选人ID
    public void setCandidate_Id(Integer candidate_Id) {
        this.candidate_Id = candidate_Id;
    }

    public Candidate getCandidate() {
        return candidate;
    }

    public void setCandidate(Candidate candidate) {
        this.candidate = candidate;
    }

}

