package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StrResponse ...
type StrResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/translate", Handler)

	e.Logger.Fatal(e.Start(":" + port))

}

// Handler e.GET("/translate?word=hi", show)
func Handler(c echo.Context) error {

	// Get team and member from the query string
	word := c.QueryParam("word")

	translatedWord := Translate(word)

	strResponse := StrResponse{}
	strResponse.Status = http.StatusOK
	strResponse.Data = translatedWord
	strResponse.Error = nil

	return c.JSON(http.StatusOK, &strResponse)
}

// Translate ...
func Translate(word string) string {

	var response string

	switch strings.ToUpper(word) {
	case "HI":
		response = "Hola"
	case "HELLO":
		response = "Hola"
	case "BYE":
		response = "Adios"
	default:
		response = "No implementado"
	}

	return response

}
