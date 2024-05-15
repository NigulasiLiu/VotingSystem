package com.example.server0.service.eval;

import com.example.server0.model.Vote;
import com.example.server0.model.VoteVoter;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Random;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

import static com.example.server0.service.eval.ConvertG.Convert;
import static com.example.server0.service.eval.Getbit.getbit;
import static com.example.server0.service.eval.Hash.hashBit;
import static com.example.server0.service.eval.HashPrime.hash;
import static com.example.server0.service.eval.HashPrime.hashPrime;
import static com.example.server0.service.eval.PG.pg;

public class Eval {
    private BigInteger[] outputs0;
    private String VW0;

    public Eval(BigInteger[] outputs0, String VW0) {
        this.outputs0 = outputs0;
        this.VW0 = VW0;
    }

    public BigInteger[] getOutputs0() {
        return outputs0;
    }

    public String getVW0() {
        return VW0;
    }

    // Eval 函数
    public static Eval eval(String votes, int[] value,int lamda) {
        int eta = value.length;
        // 初始化输出数组、验证参数和值
        BigInteger[] outputs0 = new BigInteger[eta];
        BigInteger val = BigInteger.ZERO;
        BigInteger pi = BigInteger.ZERO;
        int n = (int) Math.ceil(Math.log(eta) / Math.log(2));
        ParseK parse=ParseK.parseVotes(votes,lamda);
        String s= parse.getS0();
        String t= parse.getT0();
        String[] CW=parse.getCWArray();
        String cs= parse.getCs();
        String pos= parse.getPos();
        int posS=pos.charAt(0);


        // 遍历每个输入值
        for (int omega = 0; omega < eta; omega++) {
            // 获取当前输入值 x
            BigInteger x = new BigInteger(String.valueOf(value[omega]));
            String xx = x.toString(2);

            // 对于每个修正参数 CW(j)，执行循环
            for (int j = 1; j < n; j++) {
                // 解析 CW(j)，获取 sCW 和 tCW
                String sCW = CW[j-1].substring(0, lamda);
                String tCWL = CW[j-1].substring(lamda, lamda+1);
                String tCWR = CW[j-1].substring(lamda+1, lamda+2);

                //将s扩展为2λ+2位
                s=pg(s);
                String sL = s.substring(0, lamda);
                String tL = s.substring(lamda,lamda+1);
                String sR = s.substring(lamda+1,lamda+lamda+1);
                String tR = s.substring(lamda+lamda+1,lamda+lamda+2);

                // 将字符串转换为 BigInteger 对象
                BigInteger sl = new BigInteger(sL, 2);
                BigInteger tl = new BigInteger(tL, 2);
                BigInteger scw = new BigInteger(sCW, 2);
                BigInteger tcwl=new BigInteger(tCWL,2);
                BigInteger tcwr=new BigInteger(tCWR,2);
                BigInteger sb=new BigInteger(s,2);
                BigInteger tb=new BigInteger(t,2);
                BigInteger sr=new BigInteger(sR,2);
                BigInteger tr=new BigInteger(tR,2);



                if(xx.charAt(j)==0){
                    // 计算 s = sL 异或 (t 按位与 sCW)
                     sb= sl.xor(tb.and(scw));
                     tb=tl.xor(tb.and(tcwl));
                }
                else{

                    sb=sr.xor(tb.and(scw));
                    tb=tr.xor(tb.and(tcwr));

                }
                s=sb.toString();
                t=tb.toString();
            }
            //现在是得出y
            String sN=Convert(s);
            BigInteger sn=new BigInteger(sN,2);
            BigInteger tn=new BigInteger(t,2);
            String cwN=CW[CW.length-1];
            BigInteger cwn=new BigInteger(cwN,2);
            BigInteger y=sn.add(tn.and(cwn));

            String xn=xx+s;
            String piPrime=hashBit(xn,n);
            int m=getbit(s,posS);

            String veR=hashPrime(cs);
            BigInteger ver=new BigInteger(veR,2);

            if(m==0){
                String p=hashPrime(piPrime);
                BigInteger P=new BigInteger(p,2);
                ver=ver.xor(P);
            }else{
                BigInteger pPrime=new BigInteger(piPrime,2);
                BigInteger cS=new BigInteger(cs,2);
                BigInteger r=pPrime.xor(cS);
                String R=r.toString();
                String q=hashPrime(R);
                BigInteger Q=new BigInteger(q,2);

                ver=ver.xor(Q);
            }


            outputs0[omega]=y;
            val=val.add(y);
            pi=pi.xor(ver);



        }
        String VW0=(pi.toString())+(val.toString());
        return new Eval(outputs0, VW0);
    }

    public static String[] parseVotes(String votes, int lambda) {
        int cwLength = lambda + 2; // 每个 CW 的长度
        int csLength = 4 * lambda; // cs 的长度
        int posLength = 1; // pos 的长度

        // 计算 s 和 t 的长度
        int sLength = lambda;
        int tLength = 1;

        // 计算 CW 的个数
        int numCWs = (votes.length() - sLength - tLength - csLength - posLength) / cwLength;

        // 初始化结果数组
        String[] result = new String[numCWs + 4]; // 加上 s、t、cs 和 pos

        // 解析 s 和 t
        result[0] = votes.substring(0, sLength); // s
        result[1] = votes.substring(sLength, sLength + tLength); // t

        // 解析 CWs
        int start = sLength + tLength;
        for (int i = 0; i < numCWs; i++) {
            result[i + 2] = votes.substring(start + i * cwLength, start + (i + 1) * cwLength);
        }

        // 解析 cs
        result[numCWs + 2] = votes.substring(votes.length() - csLength - posLength , votes.length() - posLength);

        // 解析 pos
        result[numCWs + 3] = votes.substring(votes.length() - posLength);

        return result;
    }

}