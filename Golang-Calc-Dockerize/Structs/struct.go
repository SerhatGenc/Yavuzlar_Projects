package structs

type TestCase struct {
	Input          int     `json:"input"`
	ExpectedOutput float64 `json:"expected_output"`
}

type Problem struct {
	Description string     `json:"description"`
	TestCases   []TestCase `json:"test_cases"`
}

type Data struct {
	Problems []Problem `json:"problems"`
}

type RunRequest struct {
	Code string `json:"Code"`
	Fonk string `json:"FuncName"`
}

type RunResponse struct {
	Stdout string `json:"Result"`
}
