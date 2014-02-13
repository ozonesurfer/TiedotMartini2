// datainit
package main

import (
	"fmt"
	"tiedotmartini2"
	"tiedotmartini2/models"
)

func main() {
	fmt.Println("Initializing the database...")
	database := models.GetDB()
	database.Drop(tiedotmartini2.BAND_COL)
	database.Drop(tiedotmartini2.LOCATION_COL)
	database.Drop(tiedotmartini2.GENRE_COL)
	database.Create(tiedotmartini2.BAND_COL, 1)
	database.Create(tiedotmartini2.LOCATION_COL, 1)
	database.Create(tiedotmartini2.GENRE_COL, 1)
	col := database.Use(tiedotmartini2.BAND_COL)
	col.Index([]string{"albums", "genre_id"})
	database.Close()
}
