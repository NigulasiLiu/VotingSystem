package com.example.server0.service.eval;
import java.util.Arrays;

public class Hash {
    public static String hash(String bitArray) {
        int lambda = bitArray.length();
        int targetLength = 4 * lambda;
        if (bitArray.length() >= targetLength) {
            return bitArray.substring(0, targetLength);
        }
        int repeatedBits = (int) Math.ceil((double) targetLength / bitArray.length());
        StringBuilder extendedBitArray = new StringBuilder();
        for (int i = 0; i < repeatedBits; i++) {
            extendedBitArray.append(bitArray);
        }
        return extendedBitArray.substring(0, targetLength);
    }


    public static String hashBit(String input,int n) {

        StringBuilder a=new StringBuilder();
        int num=input.length();
        int num1=num-n;
        for(int i=0;i<num1;i++) {
            a.append(input.charAt(i));
        }
        String input1=a.toString();
        String hashedOutput = hash(input1); // 计算哈希值
        return hashedOutput;
    }
}




