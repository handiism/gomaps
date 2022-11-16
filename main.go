package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	app.GET("/", SearchPlaces())
	app.Run(":6000")
}

func SearchPlaces() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := c.DefaultQuery("input", "Jakarta, Indonesia")
		inputType := c.DefaultQuery("inputtype", "textquery")
		fields := c.DefaultQuery("fields", "formatted_address,name,geometry")
		key := c.DefaultQuery("key", "")

		query := url.Values{}
		query.Add("input", input)
		query.Add("inputtype", inputType)
		query.Add("fields", fields)
		query.Add("key", key)

		uri := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?%s", query.Encode())

		res, err := http.Get(uri)
		if err != nil {
			c.JSON(res.StatusCode, map[string]any{
				"error": err.Error(),
			})
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
			return
		}

		var responseBody map[string]any
		if err := json.Unmarshal(body, &responseBody); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		c.JSON(res.StatusCode, responseBody)
	}
}
