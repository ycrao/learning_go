package main

import "github.com/kataras/iris/v12"

func info(ctx iris.Context) {
	method := ctx.Method()
	subDomain := ctx.Subdomain()
	path := ctx.Path()
	url := ctx.FullRequestURI()
	ctx.Writef("\n-----\n[Request Info]\n")
	ctx.Writef("Method: %s\nSubDomain: %s\nPath: %s\nFull URL: %s", method, subDomain, path, url)
	ctx.Writef("\n-----\n")
}

func main() {
	app := iris.New()
	
	www := app.Subdomain("www")
	app.SubdomainRedirect(app, www)
	
	www.Get("/", func(ctx iris.Context) {
		ctx.Writef("INDEX from www.example.test")
		info(ctx)
	})
	
	www.Get("/hey", func(ctx iris.Context) {
		ctx.Writef("HEY from www.example.test")
		info(ctx)
	})
	
	zh := app.Subdomain("zh")
	
	zh.Get("/", func(ctx iris.Context) {
		ctx.Writef("来自 zh.example.test 的主页")
		info(ctx)
	})
	
	zh.Get("/hey", func(ctx iris.Context) {
		ctx.Writef("来自 zh.example.test 的问候")
		info(ctx)
	})
	
	/*
	For local development you'll have to edit your hosts, 
	for example in windows operating system open the C:\Windows\System32\Drivers\etc\hosts file and append:
	
	```plaintext
127.0.0.1 example.test
127.0.0.1 www.example.test
127.0.0.1 zh.example.test
	```
	 */
	app.Run(iris.Addr("example.test:8080"))
}