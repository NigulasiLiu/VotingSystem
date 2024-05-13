package com.example.server1.mapper;

import com.example.server1.model.Voter;
import org.apache.ibatis.annotations.Param;

public interface VoterMapper {
    Voter selectByVoter(@Param("voter") Voter voter);

    Voter login(@Param("voter") Voter voter);

    Voter register(@Param("voter") Voter voter);

    void deleteByVoterId(Integer voterId);

    void updateVoterType(@Param("voterId")Integer voterId,
                         @Param("voterType")String voterType);

}

