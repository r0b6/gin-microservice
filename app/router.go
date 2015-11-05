package app

import (
    "github.com/gin-gonic/gin"
    "github.com/stretchr/stew/objects"
    "fmt"
)

func Router(r *gin.Engine) {

    r.Use(ValidCityMiddleware())

    r.GET("/temperature/:city", func(c *gin.Context) {
    	city := c.Param("city")

    	url := fmt.Sprintf("%s/weather?q=%s&units=imperial&appid=%s", 
                                apiBaseUrl, city, appId)

        GetOpenWeatherMapData(url, func(data gin.H, err error) {
            if err != nil {
                c.JSON(404, gin.H{
                        "message": "Data Not Available",
                    })
            } else {
                dataMap := objects.Map(data)
                data = gin.H{
                    "temperature": gin.H{
                        "units": "Fahrenheit",
                        "low": dataMap.Get("main.temp_min"),
                        "high": dataMap.Get("main.temp_max"),
                        "current": dataMap.Get("main.temp"),
                    },
                }
                c.JSON(200, data)
            }
        })
    })

    r.GET("/wind/:city", func(c *gin.Context) {
    	city := c.Param("city")

    	url := fmt.Sprintf("%s/weather?q=%s&units=imperial&appid=%s", 
                                apiBaseUrl, city, appId)

        GetOpenWeatherMapData(url, func(data gin.H, err error) {
            if err != nil {
                c.JSON(404, gin.H{
                        "message": "Data Not Available",
                    })
            } else {
                dataMap := objects.Map(data)
                data = gin.H{
                    "wind": gin.H{
                        "speed" : gin.H{
                            "value": dataMap.Get("wind.speed"),
                            "units": "mph",
                        },
                        "direction": gin.H{
                            "value": dataMap.Get("wind.deg"),
                            "units": "degrees",
                        },
                    },
                }
                c.JSON(200, data)
            }
        })
    })

}