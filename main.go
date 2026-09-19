package myhttpfunction

import (
    "fmt"
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
    // Register an HTTP function with the Functions Framework
    functions.HTTP("MyHTTPFunction", HTTPFunction)
}

// Function myHTTPFunction is an HTTP handler
func HTTPFunction(w http.ResponseWriter, r *http.Request) {
    // Your code here

    // Send an HTTP response
    fmt.Fprintln(w, "OK")
}

