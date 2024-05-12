package com.example.server0.service;

import com.example.server0.mapper.VoterMapper;
import com.example.server0.model.Voter;
import jakarta.annotation.Resource;
import org.springframework.stereotype.Service;
import org.springframework.util.DigestUtils;

import java.nio.charset.StandardCharsets;
import java.sql.Timestamp;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class VoterService {
    @Resource
    private VoterMapper voterMapper;

    public Voter selectByVoter(Voter voter) {
        return voterMapper.selectByVoter(voter);
    }

    public Map<String, Object> login(Voter voter) {
        String password = voter.getPassword();
        byte[] bytes = password.getBytes(StandardCharsets.UTF_8);
        String passwordMD5 = DigestUtils.md5DigestAsHex(bytes);
        voter.setPassword(passwordMD5);

        Voter loginVoter = voterMapper.login(voter);
        Map<String, Object> map = new HashMap<>();   // 初始化一个返回数据对象
        if(loginVoter == null) {   // 代表用户名不存在
            map.put("code", 422);
            map.put("msg", "用户不存在");
        }
        else {                    // 代表用户登录成功 用户名和密码正确
            map.put("code", 200);
            map.put("msg", "登录成功");
            map.put("voter", loginVoter);
        }
        return map;
    }

    public Map<String, Object> register(Voter voter) {
        String password = voter.getPassword();
        byte[] bytes = password.getBytes(StandardCharsets.UTF_8);
        String passwordMD5 = DigestUtils.md5DigestAsHex(bytes);
        voter.setPassword(passwordMD5);

        Voter registerVoter=voterMapper.register(voter);
        Map<String, Object> map = new HashMap<>();
        if(registerVoter == null) {   // 代表用户名不存在
            map.put("code", 200);
            map.put("msg", "新增用户成功");
        }
        else {                    // 代表用户登录成功 用户名和密码正确
            map.put("code", 422);
            map.put("msg", "已注册");
        }
        return map;
    }

    public Map<String, Object> deleteByVoterId(Integer voterId) {
        voterMapper.deleteByVoterId(voterId);
        Map<String, Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "删除用户成功");
        return map;
    }

    public Map<String, Object> updateVoterType(Integer voterId, String voterType) {
        voterMapper.updateVoterType(voterId, voterType);
        Map<String, Object> map = new HashMap<>();
        map.put("code", 200);
        map.put("msg", "更新用户成功");
        return map;
    }
}


