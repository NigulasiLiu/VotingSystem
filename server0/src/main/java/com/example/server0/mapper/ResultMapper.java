package com.example.server0.mapper;

import com.example.server0.model.Candidate;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Result;
import org.apache.ibatis.annotations.Results;
import org.apache.ibatis.annotations.Select;

import java.util.List;

public interface ResultMapper {
    void insertResult(@Param("voteId") Integer voteId,@Param("outputs0") byte[] outputs0,@Param("vw0") String vw0);

    @Results(id = "vwResultMap0", value = {
            @Result(property = "vw0", column = "vw0")
    })
    @Select("SELECT ARRAY_AGG(vw0) AS vw0s FROM results WHERE voteId = #{voteId}")
    List<String> vw0ByVoteId(@Param("voteId") Integer voteId);

    @Results(id = "vwResultMap1", value = {
            @Result(property = "vw1", column = "vw1")
    })
    @Select("SELECT ARRAY_AGG(vw1) AS vw1s FROM results WHERE voteId = #{voteId}")
    List<String> vw1ByVoteId(@Param("voteId") Integer voteId);

    List<byte[]> outputs0ByVoteId(@Param("voteId") Integer voteId);
}
