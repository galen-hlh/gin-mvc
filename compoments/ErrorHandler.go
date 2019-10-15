package compoments

type ErrorHandler struct {
	EventBeforeError string
}

func (e *ErrorHandler) renderException() {
	print(1)
}

func errorHandler() ErrorHandler {
	return ErrorHandler{
		EventBeforeError: "beforeError",
	}
}
