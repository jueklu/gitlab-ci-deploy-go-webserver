// Declare the main package / required for any standalone executable Go program
package main

import (
        "github.com/gin-gonic/contrib/static"
        "github.com/gin-gonic/gin"
)

func main() {
        r := gin.Default()

        // Serve static files from the "./views" directory
        r.Use(static.Serve("/", static.LocalFile("./views", true)))

        // Start the server on port 8080
        r.Run("0.0.0.0:8080")
}