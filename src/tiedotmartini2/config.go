// config
package tiedotmartini2

import (
//	"github.com/QLeelulu/goku"
//	"gotiedotweb/models"
//	"path"
//	"runtime"
//	"time"
)

/*var Config *goku.ServerConfig = &goku.ServerConfig{
	Addr:           "localhost:8000",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
	StaticPath:     "static",
	ViewPath:       "views",
	LogLevel:       goku.LOG_LEVEL_LOG,
	Debug:          true,
}
*/
const (
	DATABASE_DIR = "C:\\tmp\\TiedotMartini2"
	BAND_COL     = "bands"
	LOCATION_COL = "locations"
	GENRE_COL    = "genres"
	ALBUM_COL    = "albums"
)

/*
func init() {
	// project root directory
	_, filename, _, _ := runtime.Caller(1)
	Config.RootDir = path.Dir(filename)
}
*/
