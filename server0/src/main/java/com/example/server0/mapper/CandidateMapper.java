package com.example.server0.mapper;

import com.example.server0.model.Candidate;
import org.apache.ibatis.annotations.Param;

public interface CandidateMapper {
    Candidate selectByCandidate(@Param("candidate") Candidate candidate);
}
