package com.example.server0.service.eval;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public class Hash {

    // 哈希函数，将输入的λ+n位二进制字符串通过SHA-256算法计算哈希值，并截取前4λ位作为结果
    public static String hash(String input) {
        try {
            // 创建SHA-256哈希算法实例
            MessageDigest digest = MessageDigest.getInstance("SHA-256");

            // 计算哈希值
            byte[] hashBytes = digest.digest(input.getBytes());

            // 截取前4λ位作为结果
            int numbers=input.length()/2;
            byte[] resultBytes = new byte[numbers];
            System.arraycopy(hashBytes, 0, resultBytes, 0, numbers); // 复制前4λ位到结果数组

            // 将结果数组转换为二进制字符串

            StringBuilder binaryString = new StringBuilder();
            for (byte resultByte : resultBytes) {
                String binary = String.format("%8s", Integer.toBinaryString(resultByte & 0xFF)).replace(' ', '0');
                binaryString.append(binary);
            }

            return binaryString.toString();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
            return null;
        }
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




