package com.example.server0.service.eval;

import java.util.Random;

public class PG {

    public static String pg(String bit) {
        // 生成扩展后的比特串
        String extendedBits = extendBits(bit);
        return extendedBits;
    }

    // 将长度为lambda的比特串扩展为长度为2*(lambda+1)的比特串
    public static String extendBits(String input) {
        StringBuilder extendedBits = new StringBuilder();

        // 使用伪随机生成器
        Random random = new Random();

        // 从输入比特串中的每个比特位开始
        for (int i = 0; i < input.length(); i++) {
            char bit = input.charAt(i);

            // 复制当前比特位
            extendedBits.append(bit);

            // 使用伪随机生成器生成一个随机比特位
            char randomBit = (char) ('0' + random.nextInt(2));

            // 将随机生成的比特位追加到扩展比特串中
            extendedBits.append(randomBit);
        }
        char bb = (char) ('0' + random.nextInt(2));
        char aa = (char) ('0' + random.nextInt(2));
        extendedBits.append(bb).append(aa);


        return extendedBits.toString();
    }
}