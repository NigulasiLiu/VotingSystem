package com.example.server0.service.eval;

import com.example.server0.model.Vote;
import com.example.server0.model.VoteVoter;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Random;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public class Eval {
    private double[] outputs0;
    private byte[] VW0;

    public Eval(double[] outputs0, byte[] VW0) {
        this.outputs0 = outputs0;
        this.VW0 = VW0;
    }

    public double[] getOutputs0() {
        return outputs0;
    }

    public byte[] getVW0() {
        return VW0;
    }

    // Eval 函数
    public static Eval eval(String votes, int[] value) {
        // 初始化输出数组、验证参数和值
        BigInteger[] outputs = new BigInteger[eta];
        BigInteger val = BigInteger.ZERO;
        BigInteger pi = BigInteger.ZERO;

        // 遍历每个输入值
        for (int omega = 0; omega < eta; omega++) {
            // 获取当前输入值 x
            BigInteger x = new BigInteger(xValues[omega]);

            // 初始化 s^(0) 和 t^(0)
            STValues sT = new STValues(BigInteger.ONE, BigInteger.ZERO);

            // 对于每个修正参数 CW(j)，执行循环
            for (int j = 0; j < CWs.length - 1; j++) {
                // 解析 CW(j)，获取 sCW 和 tCW
                String[] parts = CWs[j].split("\\|");
                String sCW = parts[0];
                String tCW = parts[1];

                // 执行群操作 G，得到 s^(j) 和 t^(j)
                sT = G(sT, sCW, tCW);

                // 根据 x 的第 j 位判断取 sL/tL 还是 sR/tR
                BigInteger sL = sT.s;
                BigInteger tL = sT.t;
                BigInteger sR, tR; // 获取 sR 和 tR 的值

                BigInteger s, t;
                if (x.testBit(j)) {
                    s = sR.xor(tL.multiply(new BigInteger(sCW)));
                    t = tR.xor(tL.multiply(new BigInteger(tCW)));
                } else {
                    s = sL.xor(tL.multiply(new BigInteger(sCW)));
                    t = tL.xor(tL.multiply(new BigInteger(tCW)));
                }

                // 更新 s^(j) 和 t^(j)
                sT = new STValues(s, t);
            }

            // 计算 y = f(x) 产生的输出
            BigInteger y = Convert(sT.s.add(sT.t.multiply(new BigInteger(CWs[CWs.length - 1]))));
            y = y.mod(BigInteger.valueOf(2));

            // 计算 π' 和 t
            byte[] xBytes = x.toByteArray();
            byte[] sNBytes = sT.s.toByteArray();
            byte[] piPrimeData = new byte[xBytes.length + sNBytes.length];
            System.arraycopy(xBytes, 0, piPrimeData, 0, xBytes.length);
            System.arraycopy(sNBytes, 0, piPrimeData, xBytes.length, sNBytes.length);

            BigInteger piPrime = H(piPrimeData);
            int pos = posValues[omega];
            BigInteger tValue = GB(sT.s, pos);

            // 计算 ver
            BigInteger ver = cs;
            if (tValue.equals(BigInteger.ZERO)) {
                ver = ver.xor(HPrime(piPrime.toByteArray()));
            } else {
                ver = ver.xor(HPrime(piPrime.xor(cs).toByteArray()));
            }

            // 将 y 存入输出数组，并累加到 val 中
            outputs[omega] = y;
            val = val.add(y);

            // 更新 π
            pi = pi.xor(ver);
        }

        // 拼接验证参数和 val，得到 VW
        BigInteger[] VW = Arrays.copyOf(pi.toByteArray(), pi.toByteArray().length + 1);
        VW[VW.length - 1] = val;
    }
    return new Eval(outputs0, VW0);
}