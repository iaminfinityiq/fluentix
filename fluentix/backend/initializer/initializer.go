package initializer

import (
	"fluentix/backend/value_types"
	"fluentix/runtime"
	"fmt"
)

func MakeType(value string) value_types.ValuedObject {
	var properties map[string]value_types.Object = make(map[string]value_types.Object)
	var constants map[string]bool = make(map[string]bool)
	return value_types.ValuedObject{
		Name:       "type",
		Properties: properties,
		Constants:  constants,
		Value:      value,
	}
}

func MakeBuiltinFunction(value func([]*value_types.Object) runtime.RuntimeResult) value_types.ValuedObject {
	var properties map[string]value_types.Object = make(map[string]value_types.Object)
	var constants map[string]bool = make(map[string]bool)

	return value_types.ValuedObject{
		Name:       "builtin_function",
		Properties: properties,
		Constants:  constants,
		Value:      value,
	}
}

func MakeInt(value int64) value_types.ValuedObject {
	var properties map[string]value_types.Object = make(map[string]value_types.Object)
	var constants map[string]bool = make(map[string]bool)
	constants["Plus"] = true
	constants["Minus"] = true
	constants["Multiply"] = true
	constants["Divide"] = true
	constants["Modulus"] = true
	constants["Absolute"] = true
	constants["Factorial"] = true
	constants["Positive"] = true
	constants["Negative"] = true
	constants["Equals"] = true
	constants["NotEquals"] = true
	constants["Greater"] = true
	constants["Smaller"] = true
	constants["GreaterThanOrEquals"] = true
	constants["SmallerThanOrEquals"] = true

	properties["Plus"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) + float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) + right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '+' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Minus"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) - float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) - right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '-' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Multiply"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) * float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) * right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '*' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Divide"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				if right.Value.(int64) == 0 {
					return runtime.Failure(runtime.MathError{Reason: fmt.Sprintf("Cannot divide %d by 0", value)})
				}

				var result float64 = float64(value) / float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				if right.Value.(float64) == 0 {
					return runtime.Failure(runtime.MathError{Reason: fmt.Sprintf("Cannot divide %d by 0", value)})
				}

				var result float64 = float64(value) / right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '/' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Modulus"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				if right.Value.(int64) == 0 {
					return runtime.Failure(runtime.MathError{Reason: fmt.Sprintf("Cannot modulo %d by 0", value)})
				}

				var result int64 = value % right.Value.(int64)
				var mod int64 = int64(result / 1)
				if result == mod {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(float64(result)))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				if right.Value.(float64) == 0 {
					return runtime.Failure(runtime.MathError{Reason: fmt.Sprintf("Cannot modulo %d by 0", value)})
				}

				var result float64 = float64(value) / right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation 'modulo' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Positive"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			return runtime.Success(MakeInt(value))
		},
	)

	properties["Negative"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			return runtime.Success(MakeInt(-value))
		},
	)

	properties["Absolute"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if value < 0 {
				return runtime.Success(MakeInt(-value))
			}

			return runtime.Success(MakeInt(value))
		},
	)

	properties["Factorial"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if value < 0 {
				return runtime.Failure(runtime.MathError{Reason: "Cannot perform operation '!' on a negative integer"})
			}

			var res int64 = 1
			for i := value; i > 1; i -= (*arguments[0]).(value_types.ValuedObject).Value.(int64) {
				res *= i
			}

			return runtime.Success(MakeInt(res))
		},
	)

	properties["Equals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Success(MakeBoolean(false))
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			return runtime.Success(MakeBoolean(value == new.Value))
		},
	)

	properties["NotEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Success(MakeBoolean(true))
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			return runtime.Success(MakeBoolean(value != new.Value))
		},
	)

	properties["Greater"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '>' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(float64(value) > result))
		},
	)

	properties["Smaller"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '<' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(float64(value) < result))
		},
	)

	properties["GreaterThanOrEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '>=' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(float64(value) >= result))
		},
	)

	properties["SmallerThanOrEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '<=' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(float64(value) <= result))
		},
	)

	return value_types.ValuedObject{
		Name:       "int",
		Properties: properties,
		Constants:  constants,
		Value:      value,
	}
}

func MakeDouble(value float64) value_types.ValuedObject {
	var properties map[string]value_types.Object = make(map[string]value_types.Object)
	var constants map[string]bool = make(map[string]bool)
	constants["Plus"] = true
	constants["Minus"] = true
	constants["Multiply"] = true
	constants["Divide"] = true
	constants["Modulus"] = true
	constants["Absolute"] = true
	constants["Factorial"] = true
	constants["Positive"] = true
	constants["Negative"] = true
	constants["Equals"] = true
	constants["NotEquals"] = true
	constants["Greater"] = true
	constants["Smaller"] = true
	constants["GreaterThanOrEquals"] = true
	constants["SmallerThanOrEquals"] = true

	properties["Plus"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) + float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) + right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '+' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Minus"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) - float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) - right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '-' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Multiply"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) * float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				var result float64 = float64(value) * right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '*' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Divide"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() == "int" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				if right.Value.(int64) == 0 {
					return runtime.Failure(runtime.MathError{Reason: fmt.Sprintf("Cannot divide %v by 0", value)})
				}

				var result float64 = float64(value) / float64(right.Value.(int64))
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			} else if (*arguments[0]).Name_() == "double" {
				var right value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)

				if right.Value.(float64) == 0 {
					return runtime.Failure(runtime.MathError{Reason: fmt.Sprintf("Cannot divide %v by 0", value)})
				}

				var result float64 = float64(value) / right.Value.(float64)
				var mod int64 = int64(result / 1)
				if result == float64(mod) {
					return runtime.Success(MakeInt(int64(result)))
				} else {
					return runtime.Success(MakeDouble(result))
				}
			}

			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Unexpected operation '/' on %s and %s", (*arguments[0]).Name_(), (*arguments[1]).Name_())})
		},
	)

	properties["Positive"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			return runtime.Success(MakeDouble(value))
		},
	)

	properties["Negative"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			return runtime.Success(MakeDouble(-value))
		},
	)

	properties["Absolute"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if value < 0 {
				return runtime.Success(MakeDouble(-value))
			}

			return runtime.Success(MakeDouble(value))
		},
	)

	properties["Equals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Success(MakeBoolean(false))
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			return runtime.Success(MakeBoolean(value == new.Value))
		},
	)

	properties["NotEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Success(MakeBoolean(true))
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			return runtime.Success(MakeBoolean(value != new.Value))
		},
	)

	properties["Greater"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '>' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(value > result))
		},
	)

	properties["Smaller"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '<' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(value < result))
		},
	)

	properties["GreaterThanOrEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '>=' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(value >= result))
		},
	)

	properties["SmallerThanOrEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "int" && (*arguments[0]).Name_() != "double" {
				return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '<=' for int and %s", (*arguments[0]).Name_())})
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			var result float64
			if new.Name_() == "int" {
				result = float64(new.Value.(int64))
			} else {
				result = new.Value.(float64)
			}

			return runtime.Success(MakeBoolean(value <= result))
		},
	)

	return value_types.ValuedObject{
		Name:       "double",
		Properties: properties,
		Constants:  constants,
		Value:      value,
	}
}

func MakeBoolean(value bool) value_types.ValuedObject {
	var properties map[string]value_types.Object = make(map[string]value_types.Object)
	var constants map[string]bool = make(map[string]bool)
	constants["Equals"] = true
	constants["NotEquals"] = true

	properties["Equals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "boolean" {
				return runtime.Success(MakeBoolean(false))
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			return runtime.Success(MakeBoolean(value == new.Value))
		},
	)

	properties["NotEquals"] = MakeBuiltinFunction(
		func(arguments []*value_types.Object) runtime.RuntimeResult {
			if (*arguments[0]).Name_() != "boolean" {
				return runtime.Success(MakeBoolean(true))
			}

			var new value_types.ValuedObject = (*arguments[0]).(value_types.ValuedObject)
			return runtime.Success(MakeBoolean(value != new.Value))
		},
	)

	return value_types.ValuedObject{
		Name:       "boolean",
		Properties: properties,
		Constants:  constants,
		Value:      value,
	}
}
