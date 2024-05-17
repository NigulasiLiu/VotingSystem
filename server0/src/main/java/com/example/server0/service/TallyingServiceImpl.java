package com.example.server0.service;

import com.example.server0.controller.TallyingController;
import com.example.server0.controller.VoteController;
import com.example.server0.model.TallyResult;
import com.example.server0.model.Vote;
import com.example.server0.model.VoteData;
import com.example.server0.model.VoteVoter;
import com.example.server0.service.eval.Eval;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class TallyingServiceImpl implements TallyingService {

    // 自动注入 VoteCandidateService
    @Autowired
    private VoteCandidateService voteCandidateService;

    @Autowired
    private VoteController voteController;

    // 自动注入 TallyingController
    @Autowired
    private TallyingController tallyingController;

    // 实现 TallyingService 接口中的 computeTally 方法
    @Override
    public List<byte[]> computeTally(List<byte[]> outputs0,String[] nHonest) {
// 初始化结果列表
        List<byte[]> result = new ArrayList<>();

        // 确保nHonest的长度和outputs0的大小一致
        if (nHonest.length != outputs0.size()) {
            throw new IllegalArgumentException();
        }

        // 遍历nHonest数组，决定是否保留对应位的outputs0
        for (int i = 0; i < nHonest.length; i++) {
            if ("0".equals(nHonest[i])) { // 如果nHonest为0则保留对应位的outputs0
                result.add(outputs0.get(i));
            }
        }

        // 返回计数结果
        return result;
    }
}
