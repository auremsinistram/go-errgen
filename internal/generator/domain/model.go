package generator

type InputError struct {
	Status int `json:"status"`

	Name string `json:"name"`

	Messages map[string]string `json:"messages"`
}

type Response struct {
	Success bool `json:"success"`

	Error ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code int `json:"code"`

	Message string `json:"message"`
}
