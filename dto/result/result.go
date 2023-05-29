package resultdto

type SuccessResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success pisah
type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
