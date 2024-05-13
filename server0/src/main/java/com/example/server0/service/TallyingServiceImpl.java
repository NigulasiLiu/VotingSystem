package com.example.server0.service;

import com.example.server0.controller.VoteController;
import com.example.server0.model.TallyResult;
import com.example.server0.model.Vote;
import com.example.server0.model.VoteVoter;
import com.example.server0.service.eval.Eval;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class TallyingServiceImpl implements TallyingService {

    // 自动注入 VoteCandidateService
    @Autowired
    private VoteCandidateService voteCandidateService;

    // 自动注入 VoteController
    @Autowired
    private VoteController voteController;

    // 实现 TallyingService 接口中的 computeTally 方法
    @Override
    public TallyResult computeTally() {
        // 从投票控制器中获取投票数组
        List<VoteVoter> votesList = voteController.getAllVotes();
        // 创建一个保存计数结果的 Map
        Map<Integer, List<double[]>> outputsResult = new HashMap<>();
        Map<Integer, List<byte[]>> vwResult = new HashMap<>();

        // 遍历每一张选票
        for (VoteVoter votes : votesList) {
            // 从选票对象中获取投票ID
            Integer voteId = votes.getVote_Id();
            // 调用 VoteCandidateService 中的方法计算候选人数量
            int numCandidates = voteCandidateService.countCandidatesByVoteId(voteId);
            // 创建一个保存计数结果的 Map
            Map<String, List<double[]>> tallyResultForOutput = new HashMap<>();
            Map<String, List<byte[]>> tallyResultForVW = new HashMap<>();
            int[] value=voteCandidateService.candidateIdsByVoteId(voteId);

            // 调用 EVAL 算法计算选民的 outputs0 和 VW0
            Eval evalResult = Eval.eval(votes, value);
            double[] outputs0 = evalResult.getOutputs0();
            byte[] VW0 = evalResult.getVW0();

            // 将 outputs0 和 VW0 添加到对应的列表中

            List<double[]> outputs0List = tallyResultForOutput.getOrDefault("outputs0", new ArrayList<>());
            List<byte[]> VW0List = tallyResultForVW.getOrDefault("VW0", new ArrayList<>());
            outputs0List.add(outputs0);
            VW0List.add(VW0);
            // 将输出列表和 VW0 列表放入 tallyResultForSingleVote map 中
            tallyResultForOutput.put("outputs0", outputs0List);
            tallyResultForVW.put("VW0", VW0List);

            // 将选民ID作为键，outputs0List 和 VW0List 组成的列表作为值，添加到计数结果 Map 中
            outputsResult.put(votes.getVoter_Id(), tallyResultForOutput.get("outputs0"));
            vwResult.put(votes.getVoter_Id(), tallyResultForVW.get("VW0"));
        }

        // 创建 TallyResult 对象，并将计数结果设置为其值
        TallyResult result = new TallyResult();
        result.setVwTally(vwResult);
        result.setOutputsTally(outputsResult);

        // 返回计数结果
        return result;
    }
}
