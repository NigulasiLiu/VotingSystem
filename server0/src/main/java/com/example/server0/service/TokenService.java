package com.example.server0.service;



import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import java.util.Date;

@Component
public class TokenService {
    @Value("${jwt.secret}")
    private static String secret;

    // 从配置文件中读取 token 过期时间
    @Value("${jwt.expiration}")
    private static long expiration;

    // 生成 token 方法，参数为用户名
    public static String generateToken(String voterName) {
        // 设置 token 过期时间
        Date expirationDate = new Date(System.currentTimeMillis() + expiration);

        // 使用 JWT 工具生成 token
        return Jwts.builder()
                .setSubject(voterName) // 设置 token 的主题为用户名
                .setExpiration(expirationDate) // 设置 token 过期时间
                .signWith(SignatureAlgorithm.HS512, secret) // 使用指定的算法和密钥签名 token
                .compact(); // 构建并返回 token 字符串
    }

    // 验证 token 方法，参数为 token 字符串和用户名
    public boolean validateToken(String token, String voterName) {
        // 从 token 中获取主题（用户名）
        String tokenVoterName = extractVoterName(token);
        // 判断 token 中的用户名是否与参数中的用户名一致，并且 token 未过期
        return tokenVoterName.equals(voterName) && !isTokenExpired(token);
    }

    // 从 token 中提取用户名
    public String extractVoterName(String token) {
        return Jwts.parser()
                .setSigningKey(secret) // 使用密钥解析 token
                .parseClaimsJws(token) // 解析 token
                .getBody() // 获取 token 的负载部分
                .getSubject(); // 获取主题（用户名）
    }

    // 判断 token 是否过期
    private boolean isTokenExpired(String token) {
        // 从 token 中获取过期时间
        Date expirationDate = Jwts.parser()
                .setSigningKey(secret) // 使用密钥解析 token
                .parseClaimsJws(token) // 解析 token
                .getBody() // 获取 token 的负载部分
                .getExpiration(); // 获取过期时间
        // 判断当前时间是否在过期时间之后
        return expirationDate.before(new Date());
    }
}