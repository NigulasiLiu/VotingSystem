package com.example.server0.mapper;

import com.example.server0.model.Voter;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface VoterMapper {
    Voter selectByVoter(@Param("voter") Voter voter);

    int selectVoterCount(@Param("voterName") String voterName,
                       @Param("voterType") String voterType);

    void insertVoter(@Param("voter") Voter voter);

    void deleteByVoterId(Integer voterId);

    void updateVoterType(@Param("voterId")Integer voterId,
                         @Param("voterType")String voterType);

}

