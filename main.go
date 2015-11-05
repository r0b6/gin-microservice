package main

import "github.com/gin-gonic/gin"
import "github.com/rbbll/gin-microservice/app"

func main() {
    r := gin.Default()

    app.Router(r)

    // Listen and server on 0.0.0.0:8080
    r.Run(":8080")
}
