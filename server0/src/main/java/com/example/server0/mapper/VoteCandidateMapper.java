package com.example.server0.mapper;

import org.apache.ibatis.annotations.MapKey;
import org.apache.ibatis.annotations.Param;

import java.util.Map;

public interface VoteCandidateMapper {

    int countCandidatesByVoteId(@Param("voteId") Integer voteId);
    int[] candidateIdsByVoteId(@Param("voteId") Integer voteId);
}
