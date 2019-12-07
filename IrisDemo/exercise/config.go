package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:   true,
		EnableOptimizations: true,
		Charset:             "UTF-8",
	})

	// config := iris.WithConfiguration(iris.YAML("./iris.yml"))
	// config := iris.WithConfiguration(iris.TOML("./iris.tml"))
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutAutoFireStatusCode,
		iris.WithOptimizations,
		iris.WithTimeFormat("Mon, 01 Jan 2006 15:04:05 GMT"),
		iris.WithOtherValue("ServerName", "my amazing iris server"),
		iris.WithOtherValue("ServerOwner", "admin@example.com"),
	)

	app.Run(iris.Addr(":8080"), config)
}
