package com.example.server1.service;

import com.example.server1.mapper.VoteCandidateMapper;
import com.example.server1.mapper.VoteMapper;
import jakarta.annotation.Resource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
public class VoteCandidateService {
    @Resource
    private static VoteCandidateMapper voteCandidateMapper;

    public static int countCandidatesByVoteId(Integer voteId) {
        // 执行查询操作
        int numCandidates = voteCandidateMapper.countCandidatesByVoteId(voteId);
        return numCandidates;
    }
}
