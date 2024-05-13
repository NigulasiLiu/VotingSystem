package com.example.server0.mapper;

import org.apache.ibatis.annotations.Param;

public interface VoteVoterMapper {
    int countVotersByVoteId(@Param("voteId") Integer voteId);
    int[] voterIdsByVoteId(@Param("voteId") Integer voteId);
}
