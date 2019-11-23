package main

import (
	"flag"
	"log"

	"github.com/go-webpack/webpack"
	"github.com/kataras/iris/v12"
)

func init() {
	// this is because public folder is shared between examples
	webpack.FsPath = "../public/webpack"
}

func homeIndex(ctx iris.Context) {
	ctx.View("home.html")
}

func main() {
	isDev := flag.Bool("dev", false, "development mode")
	flag.Parse()
	webpack.Init(*isDev)

	app := iris.New()
	app.Logger().SetLevel("debug")

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Reload(*isDev)
	// Important part:
	tmpl.AddFunc("asset", webpack.AssetHelper)
	tmpl.Layout("layout.html")

	app.RegisterView(tmpl)

	app.Get("/", homeIndex)

	log.Println("Iris demo app listening on http://localhost:9000")
	app.Run(iris.Addr(":9000"), iris.WithCharset("UTF-8"))
}
