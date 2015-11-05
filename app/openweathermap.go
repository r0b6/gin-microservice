package app

import (
    "github.com/gin-gonic/gin"
    "encoding/json"
    "net/http"
)

var (
    config = GetConfig()
    apiBaseUrl = config.GetString("open_weather_map.api_url")
    appId = config.GetString("open_weather_map.app_id")
)

func GetOpenWeatherMapData(url string, handleResponse func(gin.H, error)) {

    var (
        data gin.H
        err error
    )

    res, httpErr := http.Get(url)

    defer res.Body.Close()

    if httpErr != nil {
        data, err = nil, httpErr
    } else {
        decoder := json.NewDecoder(res.Body)
        jsonErr := decoder.Decode(&data)
        
        if jsonErr != nil {
            data, err = nil, jsonErr
        } else {
            data = gin.H(data)
        }
    }

    handleResponse(data, err)
}

