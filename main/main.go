package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"program/com.ypc/helloGin/controller"
)

func main()  {

	// Engin
	router := gin.Default()

	//router := gin.New()
	router.LoadHTMLGlob("template/*")
	router.GET("/hello", hello)

	// 请求参数在request header
	//router.POST("/user/add", func(context *gin.Context) {
	//	name := context.Request.FormValue("username")
	//	age := context.Request.PostForm.Get("age")
	//	moblie := context.Request.FormValue("mobile")
	//	sex := context.Request.FormValue("sex")
	//	param := ">>>> name=" + name + ",age=" + age + ",mobile" + moblie + ",sex=" + sex + " <<<<"
	//	log.Println(param)
	//
	//	var res *model.CommonResponse = new(model.CommonResponse)
	//
	//	res.Status = "200"
	//	res.Code = 200
	//	res.Message = "success"
	//	bytes,err := json.Marshal(res)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	context.Writer.Write(bytes)
	//})
	//
	//// 请求参数在body
	// router.POST("/user/insert", func(context *gin.Context) {
	//	 var res *model.CommonResponse = new(model.CommonResponse)
	//	 var user model.User
	//	 // 从请求体获取二进制流，转结构体
	//	 bytes,err := ioutil.ReadAll(context.Request.Body)
	//	 if err != nil {
	//		 log.Fatal(err)
	//	 }
	//
	//	 err = json.Unmarshal(bytes,&user)
	//	 if err != nil {
	//		 log.Fatal(err)
	//	 }
	//
	//	 // 绑定结构体
	// 	 err = context.ShouldBindJSON(&user)
	//	 if err != nil {
	//		 res.Code = 404
	//		 res.Status = "false"
	//		 res.Message = "error"
	//	 } else {
	//	 	 res.Message = "success"
	//	 	 res.Code = 200
	//	 	 res.Status = "true"
	//	 }
	//
	//
	//
	//	 bytes,er := json.Marshal(res)
	//	 if er != nil {
	//		log.Fatal(err)
	//	 }
	// 	 context.Writer.Write(bytes)
	// })

	// 路由组
	user := router.Group("/user")
	{	// 请求参数在请求路径上
		user.GET("/get/:id/:username",controller.QueryById)
		user.GET("/query",controller.QueryParam)
		user.POST("/insert",controller.InsertNewUser)
		user.GET("/form",controller.RenderForm)
		user.POST("/form/post",controller.PostForm)
		//可以自己添加其他，一个请求的路径对应一个函数

		// ...
	}

	file := router.Group("/file")
	{
		// 跳转上传文件页面
		file.GET("/view",controller.RenderView)
		// 根据表单上传
		file.POST("/insert",controller.FormUpload)
		file.POST("/multiUpload",controller.MultiUpload)
		// base64上传
		file.POST("/upload",controller.Base64Upload)
	}

	// 指定地址和端口号
	router.Run(":9090")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"success":true,
	})
}

