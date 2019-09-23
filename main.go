package main

import (
	"fmt"
	"gintest/pkg/setting"
	"gintest/router"
	//"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//router.InitRouter()
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "start test"})
	//})

	s := &http.Server{
		Addr:fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:router.InitRouter(),
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}
	fmt.Println(setting.HTTPPort)

	s.ListenAndServe()
}
