package web

type Meta struct {
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
}
type ResponseWithDataPagination struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Meta
}

func ApiResponseWithDataPagination(code int, status string, message string, data interface{}, meta Meta) ResponseWithDataPagination {
	jsonResponse := ResponseWithDataPagination{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	return jsonResponse
}
