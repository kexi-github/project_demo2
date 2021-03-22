package main

import (
	"flag"
	"fmt"
	"project_demo2-gin/middleware"
	"project_demo2-gin/config"
	"github.com/gin-gonic/gin"
	"os"
)

func main(){

	//启动  go run main.go -c config/RedisConfig.yaml
	// 保存命令行输入的参数
	var configpath string

	// 获取命令行输入参数
	flag.StringVar(&configpath,"c","","配置文件路径")

	flag.Parse() // 解析命令行输入参数

	var RedisConfig config.Redis

	RedisConfig.InitRedisConfig(configpath)

	fmt.Println(RedisConfig.Addr)
	fmt.Println(RedisConfig.Port)
	fmt.Println(RedisConfig.Password)
	fmt.Println(RedisConfig.DbNumber)

	r := gin.Default()

	//跨域请求
	r.Use(middleware.Cors())

	value1 := os.Getenv("Xmx")
	value2 := os.Getenv("pub_port")
	value3 := os.Getenv("inter")
	value4:= os.Getenv("XYZ_OPTS")
	value5 := os.Getenv("ENV_OPTS")

	//查
	r.GET("/getdata", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"s":"student",
			"value1":value1,
			"value2":value2,
			"value3":value3,
			"XYZ_OPTS":value4,
			"ENV_OPTS":value5,
		})
	})

	r.Run(":9090")
}


