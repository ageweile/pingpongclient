package server

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

// todo implement calculation. Specify body and provide the result
func handleCalculation(gc *gin.Context) {
	fmt.Println("Calculation called")

	var body struct {
	}

	var result struct {
		Result int `json:"result"`
	}

	if gc.Bind(&body) == nil {
		//result.Result = body.X * body.Y
		gc.JSON(200, result)
	}

}
