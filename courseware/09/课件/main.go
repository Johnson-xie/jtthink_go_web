package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"product.jtthink.com/App"
	"product.jtthink.com/App/Services"
	. "product.jtthink.com/AppInit"
)

func main() {
	logFile,err:= os.OpenFile("gin-log.log",os.O_CREATE|os.O_APPEND,0666)
	gin.DefaultWriter=io.MultiWriter(logFile)

	if err!=nil{
		log.Fatal("日志文件创建失败",err)
	}
	router:=gin.Default()
	v1:=router.Group("v1")
	{
		bookService_List_Endpoint:=Services.BookListEndPoint(&Services.BookService{})//图书endpoint
		bookService_Detail_Endpoint:=Services.BookDetailEndPoint(&Services.BookService{})//图书endpoint
		bookResponseFunc:=Services.CreateBookResponse()
		bookListHandler:=App.RegisterHandler(bookService_List_Endpoint,//业务最终函数
			Services.CreateBookListRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
			)
		bookDetailHandler:=App.RegisterHandler(bookService_Detail_Endpoint,//业务最终函数
			Services.CreateBookDetailRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
		)

		v1.Handle(HTTP_METHOD_GET,"/prods",bookListHandler)
		v1.Handle(HTTP_METHOD_GET,"/prods/:id", bookDetailHandler)
	}


	router.Run(SERVER_ADDRESS)
}
