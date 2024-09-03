package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main1() {
	password := "123456" // 这是你需要加密的密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("加密错误: %v", err)
	}
	fmt.Println("加密后的密码:", string(hashedPassword))
}
