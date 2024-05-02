package functions

import (
	stc "Calc/Structs"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HandleGo(w http.ResponseWriter, r *http.Request) {
	var runReq stc.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&runReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	file, err := os.Open("testcases.json")
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer file.Close()

	var data stc.Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("JSON decode error:", err)
		return
	}
	var response string

	for _, problem := range data.Problems {
		if problem.Description == runReq.Fonk {
			fmt.Println("Problem:", problem.Description)
			allCorrect := true
			for _, testCase := range problem.TestCases {
				input := testCase.Input
				expectedOutput := testCase.ExpectedOutput

				actualOutput := runDocker(runReq, w, input)
				clearstringactualoutput := strings.Replace(actualOutput, "\n", "", -1)

				i, err := strconv.ParseFloat(clearstringactualoutput, 64)

				if err != nil {
					fmt.Println(err)
					if i != expectedOutput {
						allCorrect = false
					}
				}
			}
			if allCorrect {
				response = "ALL TESTS SUCSESS"
			} else {
				response = "SOME TESTS ARE FAILED"
			}

		} else {
			response = "SUCCSESS"
		}

	}

	resp := stc.RunResponse{
		Stdout: response,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
