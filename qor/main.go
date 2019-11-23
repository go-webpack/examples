package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webpack/webpack"
	"github.com/qor/render"
)

var renderer *render.Render

func init() {
	// this is because public folder is shared between examples
	webpack.FsPath = "../public/webpack"

	renderer = render.New(&render.Config{
		//ViewPaths:     []string{"app/views"},
		//DefaultLayout: "application", // default value is application
		FuncMapMaker: func(*render.Render, *http.Request, http.ResponseWriter) template.FuncMap {
			return viewHelpers()
		},
	})
}

func homeIndex(ctx *gin.Context) {
	// Alternative (without FuncMapMaker):
	//renderer.Funcs(viewHelpers()).Execute(

	renderer.Execute(
		"home_index",
		gin.H{},
		ctx.Request,
		ctx.Writer,
	)
}

func viewHelpers() map[string]interface{} {
	return map[string]interface{}{"asset": webpack.AssetHelper}
}

func main() {
	isDev := flag.Bool("dev", false, "development mode")
	flag.Parse()

	webpack.Init(*isDev)

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.GET("/", homeIndex)

	if !*isDev {
		router.Static("/webpack", "../public/webpack")
	}

	log.Println("Listening on: 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
