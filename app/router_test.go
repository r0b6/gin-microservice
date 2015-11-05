package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
    "github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Fixture struct {
    Sources []OpenWeatherMapSource
    ExpectedJson string
}

type OpenWeatherMapSource struct {
    Url string
    Data string
    Status int
}

func TestTemperature_Success(t *testing.T) {

    city := "chicago"

    fixture := Fixture{
        Sources: []OpenWeatherMapSource{
            {
                Url: fmt.Sprintf("%s/weather?q=%s&units=imperial&appid=%s", apiBaseUrl, city, appId),
                Data: "../test_data/open_weather_map_api_data/search_by_city/chicago.json",
                Status: 200,
            },
        },
        ExpectedJson: "../test_data/expected_json/temperature/chicago.json",
    }

    testSuccess(t, "/temperature/chicago", fixture)
}

func TestWind_Success(t *testing.T) {

    city := "chicago"

    fixture := Fixture{
        Sources: []OpenWeatherMapSource{
            {
                Url: fmt.Sprintf("%s/weather?q=%s&units=imperial&appid=%s", apiBaseUrl, city, appId),
                Data: "../test_data/open_weather_map_api_data/search_by_city/chicago.json",
                Status: 200,
            },
        },
        ExpectedJson: "../test_data/expected_json/wind/chicago.json",
    }

    testSuccess(t, "/wind/chicago", fixture)
}

func testSuccess(t *testing.T, apiEndpoint string, fixture Fixture) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    for _, source := range fixture.Sources {
        weatherApiJson, wErr := ioutil.ReadFile(source.Data)

        if wErr != nil {
            t.Log("File error")
            t.Fail()
        }

        httpmock.RegisterResponder("GET", source.Url,
            func(req *http.Request) (*http.Response, error) {
                resp := httpmock.NewStringResponse(source.Status, string(weatherApiJson))
                return resp, nil
            },
        )

    }

    r := gin.New()

    Router(r)
    r.Use(ValidCityMiddleware())

    req, _ := http.NewRequest("GET", apiEndpoint, nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    var data gin.H

    decoder := json.NewDecoder(resp.Body)
    
    jsonErr := decoder.Decode(&data)
    
    if jsonErr != nil {
        t.Log("JSON error")
        t.Fail()
    }

    var expectedJsonData gin.H

    expectedJson, expectedJsonFileErr := ioutil.ReadFile(fixture.ExpectedJson)

    if expectedJsonFileErr != nil {
        t.Log("File error")
        t.Fail()
    }

    expectedJsonParseErr := json.Unmarshal(expectedJson, &expectedJsonData)
        
    if expectedJsonParseErr != nil {
        t.Log("JSON error")
        t.Fail()
    }

    diffResults := pretty.Diff(data, expectedJsonData)

    for _, delta := range diffResults {
        t.Logf("Difference found: %+v", delta)
    }

    assert.Equal(t, len(diffResults), 0, "Response Body and Expected JSON should have no differences")

    assert.Equal(t, resp.Code, 200)
}