package vo

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

func OK(opts ...Option) *Result {
	result := defaultResp
	for _, o := range opts {
		o(&result)
	}
	return &result
}

// func Ok[T any]() Result[T] {
// 	return Result[T]{Flat: true, Code: 200, Message: "ok"}
// }
// func Ok_data[T any](data T) Result[T] {
// 	return Result[T]{Flat: true, Code: 200, Message: "ok", Data: data}
// }
// func Ok_data_message[T any](data T, message string) Result[T] {
// 	return Result[T]{Flat: true, Code: 200, Message: message, Data: data}
// }
