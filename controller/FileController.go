package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

const BASE_NAME = "./static/file/"

func RenderView (context *gin.Context) {
	println(">>>> render to file upload view action start <<<<")
	context.Header("Content-Type", "text/html; charset=utf-8")

	context.HTML(200,"fileUpload.html",gin.H{})
}

func FormUpload (context *gin.Context) {
	println(">>>> upload file by form action start <<<<")

	fh,err := context.FormFile("file")
	checkError(err)
	fileName := fh.Filename
	//context.SaveUploadedFile(fh,BASE_NAME + fh.Filename)

	file,err := fh.Open()
	defer file.Close()
	bytes,e := ioutil.ReadAll(file)
	e = ioutil.WriteFile(BASE_NAME + fileName,bytes,0666)
	checkError(e)

	if e != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}
}

func MultiUpload(context *gin.Context) {
	println(">>>> upload file by form action start <<<<")
	form,err := context.MultipartForm()
	checkError(err)
	files := form.File["file"]

	var er error
	for _,f := range files {

		// 使用gin自带保存文件方法
		er = context.SaveUploadedFile(f,BASE_NAME + f.Filename)
		checkError(err)
	}
	if er != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}

}

func Base64Upload (context *gin.Context) {
	println(">>>> upload file by base64 string action start <<<<")

	bytes,err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(string(bytes),",")
	head := strs[0]
	body := strs[1]
	println(head + " | " + body)
	start := strings.LastIndex(head,"/")
	end := strings.LastIndex(head,";")
	tp := head[start + 1:end]

	err = ioutil.WriteFile(BASE_NAME + strconv.Itoa(time.Now().Nanosecond()) + "." + tp,[]byte(body),0666)
	checkError(err)
	//bys,err := base64.StdEncoding.DecodeString(string(bytes))
	//err = ioutil.WriteFile("./static/file/" + strconv.Itoa(time.Now().Nanosecond()),bys,0666)
	if err != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}
}
