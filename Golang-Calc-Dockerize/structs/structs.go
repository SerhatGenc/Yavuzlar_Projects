package structs

type Request struct {
	FuncName string `json:"FuncName"`
	Code     string `json:"Code"`
}
type SingleTest struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}
type Test struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	FuncName    string       `json:"funcName"`
	Tests       []SingleTest `json:"tests"`
}
type Response struct {
	Result      string `json:"result"`
	FailedTests string `json:"failedTests"`
}
type AllTests struct {
	Test []Test `json:"test"`
}
