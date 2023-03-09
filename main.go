package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

type Blog struct {
	Title string
	Content string
	PostDate string
	Author string
}

var dataBlog = []Blog{
	{
		Title: "Dikasih title masbro",
		Content: "Dikasih content masbro",
		Author: "masbro",
		PostDate: "Tanggal jadian",
	},
	{
		Title: "Siapa masbro",
		Content: "Dikasih uang masbro",
		Author: "masbro",
		PostDate: "Tanggal putus",
	},
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// untuk mengakses folder public
	e.Static("/public", "public")

	// renderer
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	// Routing
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", addProject)
	e.GET("/blog-detail/:id", blogDetail) //localhost:5000/blog-detail/0 | :id = url params
	e.GET("/form-blog", formAddBlog)      //localhost:5000/form-blog
	e.POST("/post-project", postProject)
	e.GET("/delete-blog/:id", deleteBlog)          //localhost:5000/add-blog

	fmt.Println("Server berjalan di port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))

	}



	func helloWorld(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	}

	func home(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	}

	func contact(c echo.Context) error {
		return c.Render(http.StatusOK, "contact.html", nil)
	}

	func addProject(c echo.Context) error {
		return c.Render(http.StatusOK, "addProject.html", nil)
	}

	func blog(c echo.Context) error {
		blogs := map[string]interface{} {
			"Blogs": dataBlog,
		}
		return c.Render(http.StatusOK, "index.html", blogs)
	}

	func blogDetail(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id")) // url params | dikonversikan dari string menjadi int/integer

		var BlogDetail = Blog{}

		for i, data := range dataBlog {
			if id == i {
				BlogDetail = Blog {
					Title: data.Title,
					Content: data.Content,
					PostDate: data.PostDate,
					Author: data.Author,
				}
			}
		}

		detailBlog := map[string]interface{}{
			"Blog": BlogDetail,
		}

		// data := map[string]interface{}{ // data yang akan digunakan/dikirimkan ke html menggunakan map interface
		// 	"Id":      id,
		// 	"Title":   "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		// 	"Content": "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
		// }

		return c.Render(http.StatusOK, "blog-detail.html", detailBlog)
	}

	func formAddBlog(c echo.Context) error {
		return c.Render(http.StatusOK, "add-blog.html", nil)
	}

	func postProject(c echo.Context) error {
		inputProjectName := c.FormValue("inputProjectName")
		description := c.FormValue("inputDescription")
		checkReact := c.FormValue("checkReact")
		checkNode := c.FormValue("checkNode")
		checkNext := c.FormValue("checkNext")
		checkTypescript := c.FormValue("checkTypescript")
		// request := c.Request().Form["name"]

		println("Project Name: " + inputProjectName)
		println("Description: " + description)
		println("Node Js: " + checkReact)
		println("React Js: " + checkNode)
		println("React Js: " + checkNext)
		println("React Js: " + checkTypescript)
		// println("React Js: " + request[name])

		var newBlog = Blog {
			Title: inputProjectName,
			Content: description,
			Author: "Masbro",
			PostDate: time.Now().String(),
		}

		dataBlog = append(dataBlog, newBlog)

		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	func deleteBlog(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
	}


