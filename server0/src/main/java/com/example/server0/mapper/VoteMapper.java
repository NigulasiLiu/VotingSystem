package com.example.server0.mapper;

import com.example.server0.model.Vote;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface VoteMapper {
    Vote selectByVote(@Param("vote") Vote vote);

    List<Vote> selectByVotePage(@Param("index") int index,
                                @Param("limit") int limit,
                                @Param("voteName") String voteName,
                                @Param("voteResult") String voteResult,
                                @Param("voteType") String voteType);


    int selectVoteCount(@Param("voteName") String voteName,
                        @Param("voteType") String voteType,
                        @Param("voteResult") String voteResult);

    void insertVote(@Param("vote") Vote vote);

    void deleteByVoteId(Integer voteId);

    void updateVoteTypeAndVoteDdlAndVoteResult(@Param("voterId")Integer voterId,
                                               @Param("voterType")String voterType);

}
