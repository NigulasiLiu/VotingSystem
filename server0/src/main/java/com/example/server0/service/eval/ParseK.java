package com.example.server0.service.eval;

import java.math.BigInteger;
import java.util.*;


public class ParseK {
    private String S0;
    private String T0;
    private String[] CWArray;
    private String cs;
    private String pos;

    // 构造函数
    public ParseK(String S0, String T0, String[] CWArray, String cs, String pos) {
        this.S0 = S0;
        this.T0 = T0;
        this.CWArray = CWArray;
        this.cs = cs;
        this.pos = pos;
    }

    // Getter 方法
    public String getS0() {
        return S0;
    }

    public String getT0() {
        return T0;
    }

    public String[] getCWArray() {
        return CWArray;
    }

    public String getCs() {
        return cs;
    }

    public String getPos() {
        return pos;
    }


    public static ParseK parseVotes(String votes, int lambda) {
        int cwLength = lambda + 2; // 每个 CW 的长度
        int csLength = 4 * lambda; // cs 的长度
        int posLength = 1; // pos 的长度

        // 计算 s 和 t 的长度
        int sLength = lambda;
        int tLength = 1;

        // 计算 CW 的个数
        int numCWs = (votes.length() - sLength - tLength - csLength - posLength) / cwLength;

        // 解析 s 和 t
        String S0 = votes.substring(0, sLength); // S0
        String T0 = votes.substring(sLength, sLength + tLength); // T0

        // 解析 CWs
        String[] CWArray = new String[numCWs];
        int start = sLength + tLength;
        for (int i = 0; i < numCWs; i++) {
            CWArray[i] = votes.substring(start + i * cwLength, start + (i + 1) * cwLength);
        }

        // 解析 cs
        String cs = votes.substring(votes.length() - csLength - posLength, votes.length() - posLength); // cs

        // 解析 pos
        String pos = votes.substring(votes.length() - posLength); // pos

        return new ParseK(S0, T0, CWArray, cs, pos);
    }




}
