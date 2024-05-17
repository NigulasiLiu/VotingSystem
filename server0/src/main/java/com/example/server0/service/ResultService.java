package com.example.server0.service;

import com.example.server0.mapper.CandidateMapper;
import com.example.server0.mapper.ResultMapper;
import com.example.server0.model.Candidate;
import org.apache.ibatis.annotations.Param;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class ResultService {
    @Autowired
    private ResultMapper resultMapper;
    public Map<String ,Object> insertResult(Integer voteId,byte[] outputs0,String vw0){
        resultMapper.insertResult(voteId,outputs0,vw0);
        Map<String,Object> map = new HashMap<>();
        map.put("code", 200);
        return  map;
    }

    public List<String> vw0ByVoteId( Integer voteId){
        List<String> vw0s =resultMapper.vw0ByVoteId(voteId);

        return vw0s;
    }

    public List<String> vw1ByVoteId( Integer voteId){
        List<String> vw1s =resultMapper.vw1ByVoteId(voteId);

        return vw1s;
    }

    public List<byte[]> outputs0ByVoteId( Integer voteId){
        List<byte[]> outputs0s =resultMapper.outputs0ByVoteId(voteId);
        return outputs0s;
    }

}
