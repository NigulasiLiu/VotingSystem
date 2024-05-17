package com.example.server0.service;

import com.example.server0.model.Vote;
import com.example.server0.model.TallyResult;
import java.util.List;

public  interface TallyingService {

    // 实现 TallyingService 接口中的 computeTally 方法
   List<byte[]> computeTally(List<byte[]> outputs0,String[] nHonest);
}
