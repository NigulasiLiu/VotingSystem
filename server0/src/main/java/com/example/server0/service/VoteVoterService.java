package com.example.server0.service;

import com.example.server0.mapper.VoteCandidateMapper;
import com.example.server0.mapper.VoteVoterMapper;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service
public class VoteVoterService {
    @Resource
    private VoteVoterMapper voteVoterMapper;

    public  int countVotersByVoteId(Integer voteId) {
        // 执行查询操作
        int numVoters = voteVoterMapper.countVotersByVoteId(voteId);
        return numVoters;
    }
    public int[] voterIdsByVoteId(Integer voteId) {
        // 调用 Mapper 层的方法查询指定 voteId 对应的所有 candidateId
        int[] voterIds = voteVoterMapper.voterIdsByVoteId(voteId);

        return voterIds;
    }
}
