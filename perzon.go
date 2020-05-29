// ================================================================================================
// Sidcloud by DKT/Samar
// ================================================================================================

package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ErrCheck - obsługa błedów
// ================================================================================================
func ErrCheck(errNr error) bool {
	if errNr != nil {
		log.Println(errNr)
		return false
	}
	return true
}

// ErrCheck2 - obsługa błedów bez komunikatu
// ================================================================================================
func ErrCheck2(errNr error) bool {
	if errNr != nil {
		return false
	}
	return true
}

// ReadFile - Odczyt bazy
// ================================================================================================
func ReadFile(filename string) []byte {
	file, _ := ioutil.ReadFile(filename)
	return file
}

// APIGetModel - odczytanie modelu
// ================================================================================================
func APIGetModel(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	data := ReadFile("./dist/model.json")

	c.Data(http.StatusOK, "application/text/html", data)
}

// APIGetMetadata - odczytanie metadata
// ================================================================================================
func APIGetMetadata(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	data := ReadFile("./dist/metadata.json")

	c.Data(http.StatusOK, "application/text/html", data)
}

// APIGetWeights - odczytanie weights
// ================================================================================================
func APIGetWeights(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	data := ReadFile("./dist/weights.bin")

	c.Data(http.StatusOK, "application/octet-stream", data)
}

// ================================================================================================
// MAIN()
// ================================================================================================
func main() {
	//
	// Logowanie do pliku
	//
	logFileApp, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	ErrCheck(err)
	log.SetOutput(io.MultiWriter(os.Stdout, logFileApp))

	gin.DisableConsoleColor()
	logFileGin, err := os.OpenFile("gin.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	ErrCheck(err)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFileGin)

	// Info powitalne
	//
	log.Println("==========================================")
	log.Println("=======          APP START        ========")
	log.Println("==========================================")

	//
	// Konfiguracja serwera
	//
	r := gin.Default()

	r.LoadHTMLGlob("./dist/*.html")

	// r.StaticFS("/fonts", http.Dir("./dist/fonts"))
	r.StaticFS("/css", http.Dir("./dist/css"))
	r.StaticFS("/js", http.Dir("./dist/js"))

	r.StaticFile("/", "./dist/index.html")
	r.StaticFile("favicon.ico", "./dist/favicon.ico")
	r.StaticFile("dkt.png", "./dist/dkt.png")
	// r.StaticFile("model.json", "./dist/model.json")
	// r.StaticFile("metadata.json", "./dist/metadata.json")
	// r.StaticFile("weights.bin", "./dist/weights.bin")

	r.GET("/api/v1/model", APIGetModel)
	r.GET("/api/v1/metadata", APIGetMetadata)
	r.GET("/api/v1/weights.bin", APIGetWeights)

	//
	// Start serwera
	//
	r.Run(":80")
}
