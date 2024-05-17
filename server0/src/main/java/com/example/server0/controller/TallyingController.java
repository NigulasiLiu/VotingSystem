package com.example.server0.controller;


import com.example.server0.model.TallyResult;
import com.example.server0.model.VoteData;
import com.example.server0.service.ResultService;
import com.example.server0.service.TallyingService;
import com.example.server0.service.VerifyService;
import com.example.server0.service.VoteCandidateService;
import com.example.server0.service.eval.Eval;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import static com.example.server0.service.eval.Eval.eval;

@RestController

public class TallyingController {

    private List<VoteData> kList = new ArrayList<>();
    @Autowired
    private ResultService resultService;
    @Autowired
    private VerifyService verifyService;
    @Autowired
    private TallyingService tallyingService;


    public static int lambda=256;
    @Autowired
    private VoteCandidateService voteCandidateService;

    @PostMapping("/receiveVote")
    public String receiveVote(@RequestBody VoteData voteData) {
        int voteId = voteData.getVoteId();
        String k = voteData.getK();

        // 在这里进行你的业务逻辑处理
        // ...

        int nums=voteCandidateService.countCandidatesByVoteId(voteId);
        // 调用 EVAL 算法计算选民的 outputs0 和 VW0
        Eval evalResult = Eval.eval(k, nums,lambda);
        byte[] outputs0 = evalResult.getOutputs0();
        String VW0 = evalResult.getVW0();

        resultService.insertResult(voteId, outputs0, VW0);


        //kList.add(voteData); // 将 k 添加到列表中

        // 返回响应或其他操作
        return "Success!!!";
    }
    @PostMapping("/verifyAndTally")
    public ResponseEntity<String> verifyAndTally(@RequestBody Integer voteId) {
        List<String> vw0List = resultService.vw0ByVoteId(voteId);
        String[] vw0 = vw0List.toArray(new String[0]);

        List<String> vw1List =resultService.vw1ByVoteId(voteId);
        String[] vw1 = vw1List.toArray(new String[0]);

        List<byte[]> outputs0=resultService.outputs0ByVoteId(voteId);


        // 调用验证函数
        String[] nHonest = verifyService.verify(vw0,vw1,lambda);
        // 调用计票函数
        List<byte[]> outs0 = tallyingService.computeTally(outputs0, nHonest);

        //这里要将outs0输送给EA


        // 计票成功，返回成功响应
        return ResponseEntity.ok("Tally successful!");
    }

    // 返回所有接收到的 k 列表
    @GetMapping("/kList")
    public List<VoteData> getList() {
        return kList;
    }

}

