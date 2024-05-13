package com.example.server1.service.eval;

public class Getbit {

    public static int getBit(String binaryString, int pos) {
        // 确保位置参数在有效范围内
        if (pos < 0 || pos >= binaryString.length()) {
            throw new IllegalArgumentException("Position out of range");
        }

        // 获取指定位置的比特值并返回
        char bit = binaryString.charAt(pos);
        return Character.getNumericValue(bit);
    }

    public static int  getbit(String binaryString,int pos) {
        int bitValue = getBit(binaryString, pos-1);
        return bitValue;
    }
}
