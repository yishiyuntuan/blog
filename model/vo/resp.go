package vo

// Result
// @model Result
// @description 返回结果
type Result struct {
	// 返回状态
	Status string `json:"status"`
	// 状态码
	Code int `json:"code"`
	// 返回消息
	Message string `json:"message"`
	// 返回数据
	Data any `json:"data"`
}
type Option func(res *Result)

func WithData(data any) Option {
	return func(res *Result) {
		res.Data = data
	}
}
func WithCode(code int) Option {
	return func(res *Result) {
		res.Code = code
	}
}
func WithMessage(msg string) Option {
	return func(res *Result) {
		res.Message = msg
	}
}

var defaultResp = Result{
	Status:  "success",
	Code:    200,
	Message: "",
	Data:    nil,
}

func Success(opts ...Option) *Result {
	result := defaultResp
	for _, o := range opts {
		o(&result)
	}
	return &result
}
func Fail(opts ...Option) *Result {
	result := defaultResp
	result.Status = "fail"
	for _, o := range opts {
		o(&result)
	}
	return &result
}
