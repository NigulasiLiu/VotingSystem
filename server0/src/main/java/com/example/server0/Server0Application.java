package com.example.server0;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.example.server0.mapper")
public class Server0Application {

    public static void main(String[] args) {
        SpringApplication.run(Server0Application.class, args);
    }

}
