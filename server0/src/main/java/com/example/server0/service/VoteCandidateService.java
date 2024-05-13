package com.example.server0.service;

import com.example.server0.mapper.VoteCandidateMapper;
import com.example.server0.mapper.VoteMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.Map;

@Service
public class VoteCandidateService {
    @Resource
    private  VoteCandidateMapper voteCandidateMapper;

    public  int countCandidatesByVoteId(Integer voteId) {
        // 执行查询操作
        int numCandidates = voteCandidateMapper.countCandidatesByVoteId(voteId);
        return numCandidates;
    }
}
