package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Redirect(c *gin.Context) {
	// 提取请求方法、URL 和请求体
	method := c.Request.Method
	url := "http://localhost:1010" + c.Request.RequestURI
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建新的 HTTP 请求，并设置方法、URL 和请求体
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 复制当前请求的所有头部到新的请求中
	for k, v := range c.Request.Header {
		req.Header.Set(k, v[0])
	}

	// 发送新的 HTTP 请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// 复制响应的头部到当前请求的响应中
	for k, v := range resp.Header {
		c.Writer.Header().Set(k, v[0])
	}

	// 将响应的状态码和主体写入当前请求的响应中
	c.Status(resp.StatusCode)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Write(respBody)
}

func Redirect1(c *gin.Context) {
	// 提取请求方法、URL 和请求体
	method := c.Request.Method
	url := "http://localhost:1011" + c.Request.RequestURI
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建新的 HTTP 请求，并设置方法、URL 和请求体
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 复制当前请求的所有头部到新的请求中
	for k, v := range c.Request.Header {
		req.Header.Set(k, v[0])
	}

	// 发送新的 HTTP 请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// 复制响应的头部到当前请求的响应中
	for k, v := range resp.Header {
		c.Writer.Header().Set(k, v[0])
	}

	// 将响应的状态码和主体写入当前请求的响应中
	c.Status(resp.StatusCode)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Write(respBody)
}
