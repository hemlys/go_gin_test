package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.RedirectFixedPath = true
	router.POST("/post", post)
	router.GET("/json", returnJson)
	router.GET("/json2", returnJson2)
	router.Any("/any", any)
	router.POST("/employee", demoHandler)
	router.Run(":8080")
}

func returnJson(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func returnJson2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"狀態": "ok",
	})
}

func post(c *gin.Context) {
	//msg := c.PostForm("input")
	// msg := c.DefaultPostForm("input", "表單沒有input。") // 沒有輸入參數時 可設定預設值

	// c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)

	json := User{}
	c.ShouldBind((&json))
	fmt.Printf("%v", &json)
	log.Println("長度", len(json.Name))
	if len(json.Name) == 0 {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": json.Name,
		"pwd":  json.Pwd,
	})

	// name := c.PostForm("name")
	// pwd := c.PostForm("pwd")
	// c.JSON(http.StatusOK, gin.H{
	// 	"name": name,
	// 	"pwd":  pwd,
	// })

	// name := c.Request.FormValue("name")
	// pwd := c.Request.FormValue("pwd")

	// c.JSON(http.StatusOK, gin.H{
	// 	"name": name,
	// 	"pwd":  pwd,
	// })

}

func any(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func demoHandler(c *gin.Context) {

	var m map[string]interface{}
	err := c.ShouldBind(&m)
	if err != nil {
		return
	}

	fmt.Printf("%v\n", m)

	c.JSON(200, gin.H{
		"name": m["name"],
		"pwd":  m["pwd"],
	})
}

type User struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}
