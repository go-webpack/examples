#### Example projects for go-webpack

https://github.com/go-webpack/webpack

This example shows usage with manifest plugin, for usage with statsplugin see [statsplugin example](https://github.com/go-webpack/examples/tree/statsplugin)

#### Running examples

```
yarn install # or npm install
# for development mode (start and leave running in another console)
./node_modules/.bin/webpack-dev-server --config webpack.config.js --hot --inline
# Or for production mode
./node_modules/.bin/webpack --config webpack.config.js --bail

# Gin example
cd gin && go run main.go -dev

# Gin + Render example
cd gin-render && go run main.go -dev

# Iris example
go run iris/main.go -dev

# QOR-render example
go run qor/main.go -dev

# Gin + Pango2 example
go run pongo2/main.go -dev
```
