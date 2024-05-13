package com.example.server0.service;

import com.example.server0.mapper.VoteMapper;
import com.example.server0.model.Vote;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
@Service
public class VoteService {
    @Resource
    private VoteMapper voteMapper;



    public Vote selectByVote(Vote vote) {
        return voteMapper.selectByVote(vote);
    }



    public Map<String ,Object> selectByVotePage(int page,int limit,String voteName,String voteResult,String voteType){
        int index=(page-1)*limit;
        List<Vote>  votelist=voteMapper.selectByVotePage(index,limit,voteName,voteResult,voteType);
        Map<String,Object> map = new HashMap<>();
        map.put("code", 0);
        map.put("msg", "查询成功");
        map.put("data",votelist);
        return  map;
    }

    public Map<String ,Object> insert(Vote vote){




        voteMapper.insertVote(vote);
        Map<String,Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "新增投票成功");

        return  map;
    }


    public Map<String, Object> deleteByVoteId(Integer voteId) {
        voteMapper.deleteByVoteId(voteId);
        Map<String, Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "删除投票成功");
        return map;
    }

    public Map<String, Object> updateVoteTypeAndVoteDdlAndVoteResult(Integer voteId, String voteType) {
        voteMapper.updateVoteTypeAndVoteDdlAndVoteResult(voteId, voteType);
        Map<String, Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "更新投票成功");
        return map;
    }

}
