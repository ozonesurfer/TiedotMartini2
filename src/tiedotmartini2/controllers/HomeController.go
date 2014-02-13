// HomeController
package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"html/template"
	"net/http"
	//	"strconv"
	"tiedotmartini2"
	"tiedotmartini2/models"
)

/* func main() {
	fmt.Println("Hello World!")
} */

func HomeIndex(r http.ResponseWriter, rw *http.Request) {
	bands := models.GetAll(tiedotmartini2.BAND_COL)
	t, err := template.ParseFiles("src/tiedotmartini2/views/home/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(r, struct {
		Title string
		Bands []models.DocWithID
	}{Title: "My CD Catalog", Bands: bands})
}

func HomeGenreList(r http.ResponseWriter, rw *http.Request) {
	genres := models.GetAll(tiedotmartini2.GENRE_COL)
	t, err := template.ParseFiles("src/tiedotmartini2/views/home/genrelist.html")
	if err != nil {
		panic(err)
	}
	t.Execute(r, struct {
		Title  string
		Genres []models.DocWithID
	}{Title: "List of Genres", Genres: genres})
}

func HomeByGenre(params martini.Params, r http.ResponseWriter, rw *http.Request) {
	id := params["id"]
	//	id, _ := strconv.ParseUint(rawId, 10, 64)
	//	id := models.ToObjectId(rawId)
	bands := models.GetBandsByGenre(id)
	genreName := models.GetGenreName(id)
	//	title := genreName + " Albums"
	title := fmt.Sprintf("%s Albums", genreName)
	t, err := template.ParseFiles("src/tiedotmartini2/views/home/bygenre.html")
	if err != nil {
		panic(err)
	}
	t.Execute(r, struct {
		Title string
		Bands []models.DocWithID
	}{Title: title, Bands: bands})
}
