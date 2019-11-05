package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Handle("GET", "/contact", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello from /contact</h1>")
	})
	
	// Method: "GET"
	app.Get("/", handler)

	// Method: "POST"
	app.Post("/", handler)

	// Method: "PUT"
	app.Put("/", handler)

	// Method: "DELETE"
	app.Delete("/", handler)

	// Method: "OPTIONS"
	app.Options("/", handler)

	// Method: "TRACE"
	app.Trace("/", handler)

	// Method: "CONNECT"
	app.Connect("/", handler)

	// Method: "HEAD"
	app.Head("/", handler)

	// Method: "PATCH"
	app.Patch("/", handler)

	// register the route for all HTTP Methods
	app.Any("/", handler)
	
	// offline route
	
	
	// grouping route
	
	users := app.Party("/users", handler)

	// http://localhost:8080/users/42/profile
	users.Get("/{id:uint64}/profile", handler)
	// http://localhost:8080/users/messages/1
	users.Get("/messages/{id:uint64}", handler)
	
	app.Run(iris.Addr(":8080"))
}

func handler(ctx iris.Context){
	ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}