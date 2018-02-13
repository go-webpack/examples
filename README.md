#### Example projects for go-webpack

https://github.com/go-webpack/webpack

This example shows usage with manifest plugin, for usage with statsplugin see [statsplugin example](https://github.com/go-webpack/examples/tree/statsplugin)

#### Running examples

```
yarn install # or npm install
# for development mode
./node_modules/.bin/webpack-dev-server --config webpack.config.js --hot --inline
# Or for production mode
./node_modules/.bin/webpack --config webpack.config.js --bail
go get
go run iris/main.go -dev
go run qor/main.go -dev
```
