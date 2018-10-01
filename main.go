package main

import (
	//"fmt"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"math"
	//"log"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	//"strconv"
	//"strings"

	"github.com/labstack/echo"
)

var (
	err error
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Restaurants struct {
	Collection []Restaturant
}

type Restaturant struct {
	Id int
	Name string
	Location int
	NumOnMap int `json:"num_on_map"`
	Photo string
	Logo string
	Info string
	MenuBreakfast string `json:"menu_breakfast"`
	MenuLunch string `json:"menu_lunch"`
	MenuDinner string `json:"menu_dinner"`
	MenuDrinks string `json:"menu_drinks"`
	WidgetLink string `json:"widget_link"`
	Favourite bool
	PDFName string
	NewRow bool
	EndRow bool

}

func ShowMenus(c echo.Context) error {
	Form := url.Values{}
	Form.Add("getRestaurants", "true")
	Form.Add("location", "3")
	resp, err := http.PostForm("https://foodatsky.com/sites-menus/", Form)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var rests []Restaturant
	err = json.Unmarshal(body, &rests)

	NumRests := len(rests)
	for i := 0; i < NumRests; i++ {
		if math.Mod(float64(i), 3) == 0 {
			rests[i].NewRow = true
			rests[i].EndRow = false
		} else if math.Mod(float64(i), 3) == 2 {
			rests[i].NewRow = false
			rests[i].EndRow = true
		} else {
			rests[i].NewRow = false
			rests[i].EndRow = false
		}
		rests[i].PDFName = path.Base(rests[i].MenuBreakfast)
		rests[i].Name = html.UnescapeString(rests[i].Name)
		switch rests[i].Name {
		case "The Dining Room":
			rests[i].Favourite = true
		case "The Loft":
			rests[i].Favourite = true
		case "The Market":
			rests[i].Favourite = true
		}
	}

	return c.Render(http.StatusOK, "index.html", rests[:len(rests) - 1])
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("Templates/*.html")),
	}

	e := echo.New()

	e.Renderer = t

	e.GET("/", ShowMenus)

	e.Logger.Fatal(e.Start(":1323"))
}
