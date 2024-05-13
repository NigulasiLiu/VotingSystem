package com.example.server0.controller;

import com.example.server0.model.Vote;
import com.example.server0.model.VoteVoter;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/votes")
public class VoteController {

    // 模拟存储所有选票的列表
    private List<VoteVoter> votesList = new ArrayList<>();

    // POST 请求，用于接收前端传来的选票数据并保存到列表中
    @PostMapping("/add")
    public void addVote(@RequestBody VoteVoter vote) {
        votesList.add(vote);
    }

    // GET 请求，用于获取所有选票数据
    @GetMapping("/all")
    public List<VoteVoter> getAllVotes() {
        return votesList;
    }

    // GET 请求，用于获取特定ID的选票数据
    @GetMapping("/{id}")
    public VoteVoter getVoteById(@PathVariable("id") Integer id) {
        for (VoteVoter vote : votesList) {
            if (vote.getVote_Id().equals(id)) {
                return vote;
            }
        }
        return null;
    }
}
