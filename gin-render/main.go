package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webpack/webpack"
	"github.com/unrolled/render"
)

func init() {
	// this is because public folder is shared between examples
	webpack.FsPath = "../public/webpack"
}

func viewHelpers() template.FuncMap {
	return template.FuncMap{
		"asset": webpack.AssetHelper,
	}
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

	r := render.New(render.Options{
		Layout:     "layout",
		IndentJSON: true,
		// Important part:
		Funcs:         []template.FuncMap{viewHelpers()},
		IsDevelopment: *isDev,
	})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		r.HTML(c.Writer, http.StatusOK, "main", gin.H{})
	})

	if !*isDev {
		router.Static("/webpack", "../public/webpack")
	}

	log.Println("Listening on: 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
