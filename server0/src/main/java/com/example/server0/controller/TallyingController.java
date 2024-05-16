package com.example.server0.controller;
import com.example.server0.model.VoteData;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;

@RestController
public class TallyingController {

    private List<VoteData> kList = new ArrayList<>();

    @PostMapping("/receiveVote")
    public String receiveVote(@RequestBody VoteData voteData) {
        int voteId = voteData.getVoteId();
        int voterId = voteData.getVoterId();
        String k = voteData.getK();
        int lambda = voteData.getLambda();

        // 在这里进行你的业务逻辑处理
        // ...

        kList.add(voteData); // 将 k 添加到列表中

        // 返回响应或其他操作
        return "Success!!!";
    }

    // 返回所有接收到的 k 列表
    @GetMapping("/kList")
    public List<VoteData> getList() {
        return kList;
    }

}

