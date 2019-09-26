package excpetions

func Params(text string) error {
	return &ParamsError{text}
}

type ParamsError struct {
	s string
}

func (p *ParamsError) Error() string {
	return p.s
}
