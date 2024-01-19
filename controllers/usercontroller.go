package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"project/connection"
	"project/models"
	"fmt"
	// "context"
	// "html/template"
	"golang.org/x/crypto/bcrypt"
	// "github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

)

type UserLoginSession struct {
	IsLogin 	bool
	Username    string
}

// Session := UserLoginSession{}

func Login(c echo.Context) error {
	username := c.FormValue("user")
	password := c.FormValue("pass")

	// Cari user dengan username yang sama
	var user models.User	
	if err := connection.DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("Username not found")
		return RedirectWithMessage(c, "Username not found", false, "/login")
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Wrong password")
		return RedirectWithMessage(c, "Wrong password", false, "/login")
	}

	// Buat session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("Error getting session")
		return RedirectWithMessage(c, "Error getting session", false, "/login")
	}

	sess.Values["isLogin"] = true
	sess.Values["username"] = username
	sess.Values["message"] = "Login success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	
	sess.Save(c.Request(), c.Response())

	return RedirectWithMessage(c, "Login success", true, "/")
}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	username := c.FormValue("user")
	password := c.FormValue("pass")
	email := c.FormValue("email")

	// Cari apakah username dan email sudah ada
	var user models.User
	// connection.DB.Where("username = ?", username).Or("email = ?", email).First(&user)
	connection.DB.Where("username = ?", username).Or("email = ?", email).First(&user)

	if user.Username == username {
		fmt.Println("Username atau email sudah ada")
		return RedirectWithMessage(c, "Username or email already exists", false, "/register")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password")
		return RedirectWithMessage(c, "Error hashing password", false, "/register")
	}

	// Buat user baru
	newUser := models.User{
		Name:     name,
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	// Masukkan user baru ke database
	if err := connection.DB.Create(&newUser).Error; err != nil {
		fmt.Println("Error creating user")
		return RedirectWithMessage(c, "Error creating user", false, "/register")
	}

	fmt.Println("User added successfully")
	fmt.Println("--------------------------")
	fmt.Println("Name : " + newUser.Name)
	fmt.Println("Username : " + newUser.Username)
	fmt.Println("Password : " + newUser.Password)
	fmt.Println("Email : " + newUser.Email)
	fmt.Println("--------------------------")

	return RedirectWithMessage(c, "User added successfully", true, "/login")
}

func Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("Error getting session")
		return RedirectWithMessage(c, "Error getting session", false, "/")
	}

	sess.Values["isLogin"] = false
	delete(sess.Values, "username")
	delete(sess.Values, "name")
	delete(sess.Values, "email")

	sess.Save(c.Request(), c.Response())

	fmt.Println("Breakpoint 1")
	// return c.Redirect(http.StatusMovedPermanently, "/")

	return RedirectWithMessage(c, "Logout success", true, "/")
}

// Fungsi untuk melihat isi dari semua users
func GetAllUsers(c echo.Context) error {
	var users []models.User
	connection.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func RedirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	fmt.Println("Redirecting to " + redirectPath)
	fmt.Println("Message : " + message)
	fmt.Println("Status : " + fmt.Sprintf("%t", status))
	fmt.Println("--------------------------")
	
	sess, errSess := session.Get("session", c)

	
	if errSess != nil {
		fmt.Println("Error getting session")
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}
	
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	fmt.Println("Sukses get session")
	// Usahakan returnya dengan method GET
	
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}