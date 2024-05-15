package com.example.server0.service.eval;

import java.util.Random;

public class PG {

    // 将比特串扩展并返回
    public static String pg(String bitString) {
        long seed = Long.parseLong(bitString, 2); // 将比特串转换为种子值
        StringBuilder extendedBitString = new StringBuilder(); // 扩展后的比特串

        Random random = new Random(seed); // 使用种子值创建 Random 实例

        // 循环生成比特串
        for (int i = 0; i < bitString.length(); i++) {
            // 复制当前比特位
            extendedBitString.append(bitString.charAt(i));

            // 生成一个随机比特位
            double rnd = random.nextDouble();
            char randomBit = (rnd >= 0.5) ? '1' : '0';

            // 将随机生成的比特位追加到扩展比特串中
            extendedBitString.append(randomBit);
            if (i < 2) {
                randomBit = (rnd >= 0.5) ? '0' : '1';
                extendedBitString.append(randomBit);
            }
        }

        return extendedBitString.toString();
    }

}
