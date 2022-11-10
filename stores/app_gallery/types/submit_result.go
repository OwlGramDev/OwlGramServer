package types

type SubmitResult struct {
	Ret struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"ret"`
}
