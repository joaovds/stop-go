package response

type HttpRes struct {
	StatusCode int    `json:"status_code"`
	HasError   bool   `json:"has_error"`
	Error      string `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func NewHttpRes(statusCode int, hasError bool, error string, data any) *HttpRes {
	return &HttpRes{
		StatusCode: statusCode,
		HasError:   hasError,
		Error:      error,
		Data:       data,
	}
}

func NewHttpResError(statusCode int, error string) *HttpRes {
	return &HttpRes{
		StatusCode: statusCode,
		HasError:   true,
		Error:      error,
	}
}

func NewHttpResData(statusCode int, data any) *HttpRes {
	return &HttpRes{
		StatusCode: statusCode,
		Data:       data,
	}
}
