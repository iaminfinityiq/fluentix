package runtime

type SyntaxError struct {
	Reason string
}

func (s SyntaxError) ErrorType() string {
	return "SyntaxError"
}

func (s SyntaxError) Reason_() string {
	return s.Reason
}

type DataTypeError struct {
	Reason string
}

func (d DataTypeError) ErrorType() string {
	return "DataTypeError"
}

func (d DataTypeError) Reason_() string {
	return d.Reason
}

type MathError struct {
	Reason string
}

func (m MathError) ErrorType() string {
	return "MathError"
}

func (m MathError) Reason_() string {
	return m.Reason
}

type VariableError struct {
	Reason string
}

func (v VariableError) ErrorType() string {
	return "VariableError"
}

func (v VariableError) Reason_() string {
	return v.Reason
}

type ArgumentError struct {
	Reason string
}

func (a ArgumentError) ErrorType() string {
	return "ArgumentError"
}

func (a ArgumentError) Reason_() string {
	return a.Reason
}