// base
package models

import (
	"encoding/json"
	"fmt"
	"tiedotmartini2"
	//	"loveoneanother.at/tiedot/db"
	"github.com/HouzuoGuo/tiedot/db"
	"os"
	"strconv"
)

type MyDoc map[string]interface{}

func ToObjectId(in string) uint64 {
	id, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return id
}

func GetDB() *db.DB {
	myDB, err := db.OpenDB(tiedotmartini2.DATABASE_DIR)
	if err != nil {
		panic(err)
	}
	return myDB
}

type DocWithID struct {
	DocKey uint64
	Value  map[string]interface{}
}

func GetAll(docType string) []DocWithID {
	database := GetDB()
	defer database.Close()
	collection := database.Use(docType)
	var query interface{}
	result := make(map[uint64]struct{})
	json.Unmarshal([]byte(`"all"`), &query)
	db.EvalQuery(query, collection, &result)
	var docs []DocWithID
	for id := range result {
		var doc MyDoc
		collection.Read(id, &doc)
		docObj := DocWithID{DocKey: id, Value: doc}
		docs = append(docs, docObj)
	}
	return docs
}

func AddDoc(doc MyDoc, docType string) (uint64, error) {
	database := GetDB()
	defer database.Close()
	collection := database.Use(docType)
	newId, err := collection.Insert(doc)

	return newId, err
}

func GetDoc(id uint64, docType string) (DocWithID, error) {
	database := GetDB()
	defer database.Close()
	collection := database.Use(docType)
	var value MyDoc
	a, err := collection.Read(id, &value)
	fmt.Println("Read returned", a)
	doc := DocWithID{DocKey: id, Value: value}
	return doc, err
}

type Album struct {
	Name    string `json:"album_name"`
	Year    int    `json:"year"`
	GenreId string `json:"genre_id"`
}

type Band struct {
	Name       string  `json:"name"`
	LocationId string  `json:"location_id"`
	Albums     []Album `json:"albums"`
}

type Genre struct {
	Name string `json:"name"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func (this *Album) GetGenreName() string {
	id, _ := strconv.ParseUint(this.GenreId, 10, 64)
	rawDoc, _ := GetDoc(id, tiedotmartini2.GENRE_COL)
	return rawDoc.Value["name"].(string)

}

func GetGenreName(id string) string {
	id2 := ToObjectId(id)
	rawDoc, _ := GetDoc(id2, tiedotmartini2.GENRE_COL)
	//	genre := rawDoc.Value.(map[string]interface{})
	return rawDoc.Value["name"].(string)
}

func GetBandsByGenre(id string) []DocWithID {
	database := GetDB()
	defer database.Close()
	collection := database.Use(tiedotmartini2.BAND_COL)
	var query interface{}
	//	id2, _ := strconv.ParseInt(id, 10, 64)
	q := `{"in": ["albums", "genre_id"], "eq": "` + id + `"}`
	//	json.Unmarshal([]byte(`"all"`), &query)
	err := json.Unmarshal([]byte(q), &query)
	if err != nil {
		fmt.Println("Unmarshall error on genre search:", err)
	}
	result := make(map[uint64]struct{})
	err = db.EvalQuery(query, collection, &result)
	if err != nil {
		fmt.Println("Eval error:", err)
		os.Exit(1)
	}
	fmt.Println("Ran query")
	var docs []DocWithID
	for id2 := range result {
		fmt.Println("Found", id2)
		var readback map[string]interface{}
		collection.Read(id2, &readback)
		doc := DocWithID{DocKey: id2, Value: readback}
		docs = append(docs, doc)
	}
	if docs == nil {
		fmt.Println("Returning empty value")
	}
	return docs
}

func (this *DocWithID) LocToString() string {
	location := this.Value
	city := location["city"].(string)
	state := location["state"].(string)
	country := location["country"].(string)
	var strCity, strState string
	if city == "" {
		strCity = "(city)"
	} else {
		strCity = city
	}
	if state == "" {
		strState = "(state/province)"
	} else {
		strState = state
	}
	result := strCity + ", " + strState + " " + country
	return result
}

func (this DocWithID) GetLocation() string {
	//	original := this.Value.(Band)
	original := this.Value
	idStr := original["location_id"].(string)
	fmt.Println("idStr =", idStr)
	id, _ := strconv.ParseUint(idStr, 10, 64)
	rawDoc, err := GetDoc(id, tiedotmartini2.LOCATION_COL)
	if err != nil {
		return err.Error()
	}
	result := rawDoc.LocToString()
	//	result := this
	return result
}

func (this *DocWithID) GetName() string {
	original := this.Value
	name := original["name"].(string)
	return name
}
func (this *DocWithID) AddAlbum(album Album) error {
	database := GetDB()
	defer database.Close()
	collection := database.Use(tiedotmartini2.BAND_COL)
	original := this.Value
	x := original["location_id"].(string)
	//	locationId := strconv.ParseUint(x, 10, 64)
	band := Band{Name: original["name"].(string),
		LocationId: x}
	band.Albums = []Album{}
	if original["albums"] != nil {
		for _, a := range original["albums"].([]interface{}) {
			x2 := a.(map[string]interface{})
			z := x2["genre_id"].(string)
			y := x2["year"].(float64)
			q := Album{Name: x2["album_name"].(string), Year: int(y),
				GenreId: z}
			band.Albums = append(band.Albums, q)
		}
	}
	band.Albums = append(band.Albums, album)
	bandMap := map[string]interface{}{"name": original["name"].(string),
		"location_id": x, "albums": band.Albums}

	bandMapJson, err := json.Marshal(bandMap)
	var bandMapMap map[string]interface{}
	err = json.Unmarshal(bandMapJson, &bandMapMap)
	fmt.Println(bandMapMap)
	collection.Update(this.DocKey, bandMapMap)

	//	err := collection.Update(this.DocKey, bandMap)
	//	err := collection.Update(this.DocKey, )
	//	this.DocKey = newKey
	return err
}

func (this DocWithID) GetAlbums() []Album {
	original := this.Value
	var cds []Album
	for _, a := range original["albums"].([]interface{}) {
		x := a.(map[string]interface{})
		q := Album{Name: x["album_name"].(string), Year: int(x["year"].(float64)),
			GenreId: x["genre_id"].(string)}
		cds = append(cds, q)
	}
	return cds
}
