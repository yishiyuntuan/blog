package controller

type Controller interface {
	_controller()
}

type MessageController struct {
	name string
	Controller
}

func NewMessageController(name string) *MessageController {
	return &MessageController{
		name: name,
	}
}
