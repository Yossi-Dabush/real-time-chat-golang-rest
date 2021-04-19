package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"dataBaseConn.com/db"
	"github.com/gin-gonic/gin"
)


func chatRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chatRoom" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

}

func setupRoutes(router *gin.Engine) {
	router.LoadHTMLFiles("login.html")

	router.GET("/login.html", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	router.POST("/login", func(c *gin.Context) {

		uml := strings.Split(c.FullPath(), "/")
		if uml[len(uml)-1] != "login" {
			http.Error(c.Writer, "404 not found.", http.StatusNotFound)
			return
		}
		if c.Request.Method != "POST" {
			http.Error(c.Writer, "Method is not supported.", http.StatusNotFound)
			return
		}

		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			println(err)
		}
		myString := string(jsonData[:])
		res := strings.Split(myString, ",")

		if db.UserLogin(res[0], res[1]) {
			router.LoadHTMLFiles("index.html")
			c.HTML(200, "index.html", nil)
			if err != nil {
				fmt.Fprintf(c.Writer, "Unable to load template")
			}
		} else {
			println("wrong")
		}
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/test", func(c *gin.Context) {
		fmt.Println("test")
	})
}

func main() {

	go h.run()

	router := gin.New()
	db.ConnectDB()
	setupRoutes(router)
	router.Run("localhost:8080")
}
