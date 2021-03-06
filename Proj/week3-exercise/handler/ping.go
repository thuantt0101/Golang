package handler

import
(
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
) 

var counter int = 0

func pingHandler(c *gin.Context) {

	counter += 1
	//Race condittion =>Hai CPU cung access vo 1 cai bien
	//Atomic=> co 100 request vao nhung expected :counter =100,actual :80
c.Writer.Header().Set("X-Counter",strconv.Itoa(counter))
port:=os.Getenv("HTTP_PORT")
message:= "Pong from"

if port =="8082"{
	message = message + "2nd server!"
}
else {
	message = message + "1st server"
}

c.String(200,message)

}