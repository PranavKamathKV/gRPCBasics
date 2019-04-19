package main


import (
	"fmt"
	"gRPCBasics/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func main(){
	conn, err:= grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}

	client:= proto.NewAddServiceClient(conn)

	g:= gin.Default()

	g.GET("add/:a/:b", func(ctx *gin.Context){
		a, err:= strconv.ParseUint(ctx.Param("a"), 10,64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Parameter for A"})
			return
		}

		b, err:= strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Parameter for B"})
			return
		}

		request:= &proto.Request{A:int64(a), B: int64(b)}
		response, err:= client.Add(ctx, request)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}else {
			ctx.JSON(http.StatusOK, gin.H{"Result":response.C})
		}

	})

	g.GET("mul/:a/:b", func(ctx *gin.Context){
		a, err:=strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter for A"})
			return
		}

		b, err:= strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter for B"})
			return
		}
		request:=&proto.Request{A:int64(a), B:int64(b)}
		response, err:= client.Multiply(ctx, request)

		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"Result":response.C})
		}
	})

	err= g.Run(":8080")
	if err!=nil{
		fmt.Println(err)
	}


}