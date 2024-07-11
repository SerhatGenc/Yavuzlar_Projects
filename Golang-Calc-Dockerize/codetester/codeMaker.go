package codetester

import (
	"docCalc/structs"
	"log"
	"os"
	"strconv"
	"strings"
)

func CodeMaker(request structs.Request, input float64) {
	mainCode := `package main 
	import "fmt"
	func main(){
		fmt.Println(FunctionTester(inputTest))
		}
		`

	mainCode = strings.Replace(mainCode, "FunctionTester", request.FuncName, -1)
	mainCode = strings.Replace(mainCode, "inputTest", strconv.FormatFloat(input, 'f', -1, 64), -1)

	mainCode += request.Code

	file, err := os.Create("./TestCode/TestCode.go")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err2 := file.WriteString(mainCode)

	if err2 != nil {
		log.Fatal(err2)
	}

}
