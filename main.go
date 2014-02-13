// main
package main

import (
	//	"fmt"
	//	"github.com/HouzuoGuo/tiedot/db"
	"github.com/codegangsta/martini"
	"math/rand"
	//	"tiedotmartini2"
	"tiedotmartini2/controllers"
	//	"tiedotmartini2/models"
	"time"
)

func main() {
	//	fmt.Println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano())
	/*	database := models.GetDB()
		database.Drop(tiedotmartini2.BAND_COL)
		database.Drop(tiedotmartini2.LOCATION_COL)
		database.Drop(tiedotmartini2.GENRE_COL)
		database.Create(tiedotmartini2.BAND_COL, 1)
		database.Create(tiedotmartini2.LOCATION_COL, 1)
		database.Create(tiedotmartini2.GENRE_COL, 1)
		col := database.Use(tiedotmartini2.BAND_COL)
		col.Index([]string{"albums", "genre_id"})
		database.Close()
	*/
	m := martini.Classic()
	m.Get("/", controllers.HomeIndex)
	m.Get("/home/index", controllers.HomeIndex)
	m.Get("/band/add", controllers.BandAdd)
	m.Post("/band/verify", controllers.BandVerify)
	m.Get("/album/index/:id", controllers.AlbumIndex)
	m.Get("/album/add/:id", controllers.AlbumAdd)
	m.Post("/album/verify/:id", controllers.AlbumVerify)
	m.Get("/home/genrelist", controllers.HomeGenreList)
	m.Get("/home/bygenre/:id", controllers.HomeByGenre)
	m.Use(martini.Static("assets"))
	m.Run()
}
