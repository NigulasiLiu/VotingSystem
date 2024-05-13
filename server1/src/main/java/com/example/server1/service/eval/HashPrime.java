package com.example.server1.service.eval;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
public class HashPrime {

    // 哈希函数，将输入的二进制字符串通过SHA-256算法计算哈希值，并截取长度为原字符串一半的二进制字符串作为结果
    public static String hash(String input) {
        try {
            // 创建SHA-256哈希算法实例
            MessageDigest digest = MessageDigest.getInstance("SHA-256");

            // 计算哈希值
            byte[] hashBytes = digest.digest(input.getBytes());

            // 计算截取的字节数
            int numBytes = input.length() / 16; // 输入字符串的一半长度（字节数）

            // 截取长度为原字符串一半的字节数作为结果
            byte[] resultBytes = new byte[numBytes];
            System.arraycopy(hashBytes, 0, resultBytes, 0, numBytes);

            // 将截取的字节数组转换为二进制字符串
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

    public static String hashPrime(String input) {

        String hashedOutput = hash(input); // 计算哈希值
        return hashedOutput;

    }
}


