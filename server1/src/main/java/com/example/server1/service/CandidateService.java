package com.example.server1.service;


import com.example.server1.mapper.CandidateMapper;
import com.example.server1.model.Candidate;
import jakarta.annotation.Resource;
import org.springframework.stereotype.Service;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class CandidateService {
    @Resource
    private CandidateMapper candidateMapper;
    public Candidate selectByCandidate(Candidate candidate) {
        return candidateMapper.selectByCandidate(candidate);
    }

    public Map<String ,Object> selectByPage(int page,int limit,String candidateName){
        int index=(page-1)*limit;
        List<Candidate> candidatelist=candidateMapper.selectByPage(index,limit,candidateName);
        Map<String,Object> map = new HashMap<>();
        map.put("code", 0);
        map.put("msg", "查询成功");
        map.put("data",candidatelist);
        return  map;
    }

    public Map<String ,Object> insert(Candidate candidate){



        candidateMapper.insert(candidate);
        Map<String,Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "新增候选人成功");
        return  map;
    }


    public Map<String, Object> deleteById(Integer candidateId) {
        candidateMapper.deleteById(candidateId);
        Map<String, Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "删除候选人成功");
        return map;
    }

    public Map<String, Object> updateDetails(Integer candidateId, String candidateDetails) {
        candidateMapper.updateDetails(candidateId, candidateDetails);
        Map<String, Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "更新用户成功");
        return map;
    }
}
