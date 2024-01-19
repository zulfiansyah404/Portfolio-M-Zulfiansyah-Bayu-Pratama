package controllers

import (
	"net/http"
	"project/connection"
	"project/models"
	"strconv"
	"text/template"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var ProjectList []models.Project

type SessionLogin struct {
	IsLogin bool
	Username string
}



// func updateProjectList() {
// 	// connection.DB.Find(&ProjectList)
// }

// Fungsi mengeluarkan respond Hello World
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Worldl!")
}

// Fungsi menampilkan index.html beserta ListProject
func Home(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/index.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	// Tampilkan hanya project yang authornya adalah username yang sedang login
	if err := connection.DB.Where("author = ?", sess.Values["username"]).Find(&ProjectList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	
	if sess.Values["isLogin"] == nil {
		sess.Values["isLogin"] = false
	}
	

	fmt.Println("Home")
	fmt.Println("------------------")
	fmt.Println("IsLogin: ", sess.Values["isLogin"])
	fmt.Println("Username: ", sess.Values["username"])
	fmt.Println("Message: ", sess.Values["message"])
	fmt.Println("Status: ", sess.Values["status"])
	fmt.Println("Name: ", sess.Values["name"])
	fmt.Println("Email: ", sess.Values["email"])
	fmt.Println("------------------")
	fmt.Println("")
	fmt.Println("")	

	dataResponds := map[string]interface{} {
		"Projects": ProjectList,
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
		"Name": sess.Values["name"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	// // Print isi projects
	// fmt.Println(projects)

	return tmp.Execute(c.Response(), dataResponds)
}

func LoginView(c echo.Context) error {
	// Pastikan bahwa user tidak bisa mengakses halaman login jika sudah login
	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	if sess.Values["isLogin"] == true {
		return c.Redirect(http.StatusFound, "/")
	}

	var tmp, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmp.Execute(c.Response(), dataResponds)
}

func RegisterView(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/register.html")

	// Pastikan bahwa user tidak bisa mengakses halaman login jika sudah login
	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	if sess.Values["isLogin"] == true {
		return c.Redirect(http.StatusFound, "/")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi menampilkan Menu untuk menambah project
func AddProjectView(c echo.Context) error {
	// Pastikan dahulu bahwa user sudah login
	fmt.Println("Add Project View")
	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	if sess.Values["isLogin"] != true {
		sess.Save(c.Request(), c.Response())
		return RedirectWithMessage(c, "You must login first", false ,"/login")
	}

	fmt.Println("Add Project View")
	var tmp, err = template.ParseFiles("views/project.html")

	if err != nil {
		fmt.Println("Break 1")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
		"TechError" : sess.Values["techError"],
		"DateError" : sess.Values["dateError"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	delete(sess.Values, "techError")
	delete(sess.Values, "dateError")

	sess.Save(c.Request(), c.Response())
	
	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi untuk menampilkan contact.html
func Contact(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi untuk menampilkan testimonial.html
func TestimonialView(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	sess.Save(c.Request(), c.Response())

	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi untuk menampilkan project-detail.html
func ProjectDetailView(c echo.Context) error {
	fmt.Println("Project Detail View")
	id, _ := strconv.Atoi(c.Param("id"))

	// Pastikan dahulu bahwa user sudah login
	// Pastikan bahwa user yang membuka project detail adalah user yang membuat project tersebut
	var ProjectDetail = models.Project{}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		fmt.Println("Break 1")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	if sess.Values["isLogin"] != true {
		fmt.Println("Break 2")
		sess.Save(c.Request(), c.Response())
		return RedirectWithMessage(c, "You must login first", false ,"/login")
	}

	// Pastikan bahwa user yang membuka project detail adalah user yang membuat project tersebut
	connection.DB.Where("id = ?", id).First(&ProjectDetail)
	if ProjectDetail.Author != sess.Values["username"] {
		fmt.Println("Break 3")
		sess.Save(c.Request(), c.Response())
		return RedirectWithMessage(c, "You are not allowed to access this page", false ,"/")
	}



	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id":		 id,
		"StartDate": ProjectDetail.StartDate.Format("2 January 2006"),
		"EndDate":   ProjectDetail.EndDate.Format("2 January 2006"),
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), data)
}

// Fungsi untuk menampilkan menu untuk mengedit suatu project berdasarkan ID
func EditProjectView(c echo.Context) error {
	fmt.Println("Edit Project View")
	id, _ := strconv.Atoi(c.Param("id"))

	// Pastikan dahulu bahwa user sudah login
	sess, errSess := session.Get("session", c)
	if errSess != nil {
		fmt.Println("Break 1")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	if sess.Values["isLogin"] != true {
		sess.Save(c.Request(), c.Response())
		return RedirectWithMessage(c, "You must login first", false ,"/login")
	}

	// Lalu pastikan bahwa user yang login adalah author dari project tersebut
	var ProjectDetail = models.Project{}
	connection.DB.Where("id = ?", id).First(&ProjectDetail)

	if (sess.Values["username"] != ProjectDetail.Author) {
		sess.Save(c.Request(), c.Response())
		return RedirectWithMessage(c, "You can't access this page", false ,"/")
	}

	// for _, data := range ProjectList {
	// 	if id == data.ID {
	// 		if (sess.Values["username"] != data.Author) {
	// 			fmt.Println("Tidak sesuai!")
	// 			sess.Save(c.Request(), c.Response())
	// 			return RedirectWithMessage(c, "You can't access this page", false ,"/")
	// 		} else {
	// 			ProjectDetail = models.Project{
	// 				NameProject:  	data.NameProject,
	// 				StartDate:  	data.StartDate,
	// 				EndDate:    	data.EndDate,
	// 				Duration:   	data.Duration,
	// 				Description: 	data.Description,
	// 				NodeJs:     	data.NodeJs,
	// 				ReactJs:    	data.ReactJs,
	// 				Golang:     	data.Golang,
	// 				Java: 			data.Java,
	// 				Image: 			data.Image,
	// 			}
	// 		}	
	// 	}
	// }

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id": id,
		"StartDate": ProjectDetail.StartDate.Format("2006-01-02"),
		"EndDate":   ProjectDetail.EndDate.Format("2006-01-02"),
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
		"TechError" : sess.Values["techError"],
		"DateError" : sess.Values["dateError"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	delete(sess.Values, "techError")
	delete(sess.Values, "dateError")

	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/edit-project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}