package main

import (
	"database/sql"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id int64 `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//对错误检查的封装
func check (err error){
	if err != nil{
		fmt.Println(err)
	}
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

var Db *sql.DB
func init(){
	var err error
	Db, err = sql.Open("mysql", "root:112233@tcp(127.0.0.1:3306)/go_test")
	check(err)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	err = Db.Ping()
	check(err)
	fmt.Println("数据库连接成功")
}

func Select() (users []User, err error)  {
	rows,err := Db.Query("select * from user")
	check(err)
	for rows.Next(){
		var user User
		//将从数据库中查询到的值对应到结构体中相应的变量中
		err = rows.Scan(&user.Id,&user.Username,&user.Password)
		check(err)
		users = append(users, user)
	}
	return users,err
}

func Add()  {
	ret,err := Db.Exec(`insert into user(username,password)values('cj','a123456')`)
	check(err)
	id,err := ret.LastInsertId()
	check(err)
	//修改第几行
	fmt.Println(id)
}

func Delete(id int)(rows int)  {
	ret,err := Db.Exec(`delete from user where id=?`,id)
	check(err)
	row,err := ret.RowsAffected()
	check(err)
	rows = int(row)
	return rows
}

func getUser(c *gin.Context) {
	res,err:=Select()
	check(err)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"msg":"success",
		"count":len(res),
		"data":res,
	})
}

func deleteUser(c *gin.Context)  {
	id := c.Query("id")
	Id, err := strconv.ParseInt(id, 10, 10)
	check(err)
	Delete(int(Id))
	if Delete(int(Id))==0 {
		c.JSON(http.StatusOK, gin.H{
			"code":-1,
			"msg": "删除失败或该用户已删除",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":200,
			"msg": "删除成功",
		})
	}
}

func getSend(c *gin.Context)  {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FybwHjRM7E2YDggWgPg", "jNwPYQm1Y7aBMuZGhVXuASiSoQdWyO")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "13015759718"
	request.SignName = "学习笔记App"
	request.TemplateCode = "SMS_191801632"
	request.TemplateParam = "{\"code\":\"1234\"}"

	response, err := client.SendSms(request)
	check(err)
	fmt.Printf("response is %#v\n", response)
	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"msg": "获取验证码成功",
	})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUser)
	router.POST("/users",deleteUser)
	router.POST("getSend",getSend)
	router.Run(":80") // listen and serve on 0.0.0.0:8080
}
