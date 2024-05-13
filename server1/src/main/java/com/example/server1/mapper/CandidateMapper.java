package com.example.server1.mapper;

import com.example.server1.model.Candidate;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CandidateMapper {

    Candidate selectByCandidate(@Param("candidate") Candidate candidate);


    List<Candidate> selectByPage(@Param("index") int index,
                                 @Param("limit") int limit,
                                 @Param("candidateName") String candidateName);

    void insert(@Param("candidate") Candidate candidate);


    void deleteById(Integer candidateId);
    void updateDetails(@Param("candidateId") Integer candidateId,@Param("candidateDetails") String candidateDetails);

}