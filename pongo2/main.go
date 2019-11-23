package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/go-webpack/webpack"
	"gitlab.com/glebtv/pongo2gin"

	// This is the important part:
	_ "github.com/go-webpack/pongo2"
)

func init() {
	// this is because public folder is shared between examples
	webpack.FsPath = "../public/webpack"
}

func main() {
	isDev := flag.Bool("dev", false, "development mode")
	flag.Parse()

	webpack.Init(*isDev)
	if *isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.HTMLRender = pongo2gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", pongo2.Context{})
	})

	if !*isDev {
		router.Static("/webpack", "../public/webpack")
	}

	log.Println("Listening on: 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
