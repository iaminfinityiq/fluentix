package interpreter

import (
	"fluentix/backend/initializer"
	"fluentix/frontend/parser"
	"fluentix/backend/value_types"
	"fluentix/backend/scopes"
	"fluentix/runtime"
	"fmt"
)

func MakeDefinedFunction(arguments []string, body parser.Block, scope *scopes.Scope) value_types.ValuedObject {
	var array struct {
		a []string
		b parser.Block
	} = struct {
		a []string
		b parser.Block
	}{a: arguments, b: body}
	var properties map[string]value_types.Object = make(map[string]value_types.Object)
	var constants map[string]bool = make(map[string]bool)
	constants["Call"] = true

	properties["Call"] = initializer.MakeBuiltinFunction(func(o []*value_types.Object) runtime.RuntimeResult {
		if len(o) != len(arguments) {
			return runtime.Failure(runtime.ArgumentError{Reason: fmt.Sprintf("Expected %d arguments, got %d/%d", len(arguments), len(o), len(arguments))})
		}

		var new_scope scopes.Scope = scopes.Scope{
			Scope: make(map[string]value_types.Object),
			Parent: scope,
			Constants: make(map[string]bool),
		}

		for i, argument := range arguments {
			var rt runtime.RuntimeResult = new_scope.Assign(argument, *o[i])
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}
		}

		__, _ := interface{}(body).(parser.Statement)
		var rt runtime.RuntimeResult = Evaluate(&__, &new_scope)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	})

	return value_types.ValuedObject{
		Name:       "defined_function",
		Properties: properties,
		Constants:  constants,
		Value:      array,
	}
}