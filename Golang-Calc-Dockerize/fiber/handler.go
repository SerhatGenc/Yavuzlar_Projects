package fiber

import (
	dck "docCalc/docker"
	"docCalc/structs"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {

	var req structs.Request
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile, err := os.Open("./JSON/tests.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	decode := json.NewDecoder(jsonFile)
	var tests structs.AllTests

	err = decode.Decode(&tests)
	if err != nil {
		fmt.Println(err)
	}

	selectedTest := structs.Test{}

	for _, v := range tests.Test {
		fmt.Println(v.FuncName, "   ", req.FuncName)
		if v.FuncName == req.FuncName {
			selectedTest = v
		} else {
			fmt.Println("test adi yanlis")

		}
	}

	controller := ""
	fmt.Println(selectedTest)

	for i, v := range selectedTest.Tests {
		resultB := dck.DockerUp(req, v)

		if !resultB {
			controller += strconv.Itoa(i) + "/"
		}
	}
	result := ""
	if controller == "" {

		result = "success"
	} else {
		result = "fail"

	}

	resp := structs.Response{
		Result:      result,
		FailedTests: controller,
	}

	c.Set("Content-Type", "application/json")
	return c.JSON(resp)
}
