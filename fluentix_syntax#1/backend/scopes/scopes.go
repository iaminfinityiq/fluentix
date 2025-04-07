package scopes

import (
	"fluentix/backend/value_types"
	"fluentix/runtime"
	"fmt"
)

type Scope struct {
	Scope     map[string]value_types.Object
	Parent    *Scope // nil if it's the global scope
	Constants map[string]bool
}

func (s *Scope) Assign(variable_name string, value value_types.Object) runtime.RuntimeResult {
	if s.Constants[variable_name] {
		return runtime.Failure(runtime.VariableError{Reason: fmt.Sprintf("Cannot assign to variable %s because it is a constant", variable_name)})
	}

	s.Scope[variable_name] = value
	return runtime.Success(nil)
}

func (s *Scope) Resolve(variable_name string) runtime.RuntimeResult {
	value, ok := s.Scope[variable_name]
	if !ok {
		if s.Parent == nil {
			return runtime.Failure(runtime.VariableError{Reason: fmt.Sprintf("Cannot find variable '%s'", variable_name)})
		}

		var rt runtime.RuntimeResult = s.Parent.Resolve(variable_name)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	}

	return runtime.Success(value)
}
