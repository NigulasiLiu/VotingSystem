package com.example.server1.service.eval;
import java.math.BigInteger;

public class ConvertG {

    // 将二进制字符串转换为整数
    private static int binaryStringToInt(String binaryString) {
        return Integer.parseInt(binaryString, 2);
    }

    // ConvertG函数，将λ位二进制数映射为群中的一个元素
    private static BigInteger convertG(String binaryString) {
        // 将二进制字符串转换为整数
        int intValue = binaryStringToInt(binaryString);

        // 群的二元运算，这里简单地使用加法
        // 可根据具体需求和群的定义来修改

        // 这里简单示范一下，假设群运算是对输入进行乘方
        BigInteger element = BigInteger.valueOf(intValue);
        BigInteger res = element.pow(2); // 进行群运算，这里是乘方
        return res;
    }

    public static String convertToBinary(BigInteger element) {
        // 将群中的元素转换为二进制字符串表示
        return element.toString(2);
    }

    public static String Convert(String input) {

        BigInteger result = convertG(input);
        return convertToBinary(result);
    }
}
