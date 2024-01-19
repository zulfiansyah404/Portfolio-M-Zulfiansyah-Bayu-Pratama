package main

import (
	"project/connection"
	"project/controllers"
	"project/middleware"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	connection.DatabaseConnect()
	
	// Tambahkan middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// Mengatur penanganan file static
	e.Static("/assets", "assets")
	e.Static("/uploads", "uploads")

	// Daftar Routes GET
	// Diambil dari controllers/viewcontroller.go
	e.GET("/hello", controllers.HelloWorld)
	e.GET("/", controllers.Home)
	e.GET("/add-project", controllers.AddProjectView)
	e.GET("/contact", controllers.Contact)
	e.GET("/testimonials", controllers.TestimonialView)
	e.GET("/project-detail/:id", controllers.ProjectDetailView)
	e.GET("/edit-project/:id", controllers.EditProjectView)
	e.GET("/login", controllers.LoginView)
	e.GET("/register", controllers.RegisterView)
	e.GET("/users", controllers.GetAllUsers)

	//Daftar Routes POST
	// Diambil dari controllers/projectcontroller.go
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	e.POST("/logout", controllers.Logout)
	e.POST("/", middleware.UploadFile(controllers.AddProject))
	e.POST("/edit-project/:id", middleware.UploadFile(controllers.EditProject))
	e.POST("/delete-project/:id", controllers.DeleteProject)

	// Server
	e.Logger.Fatal(e.Start("localhost:5000"))
}

