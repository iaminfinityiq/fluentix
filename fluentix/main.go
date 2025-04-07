package main

import (
	"bufio"
	// "fluentix/backend/initializer"
	// "fluentix/backend/interpreter"
	// "fluentix/backend/scopes"
	// "fluentix/backend/value_types"
	"fluentix/frontend/lexer"
	"fluentix/frontend/parser"
	"fluentix/runtime"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	file, err := os.Open("test.flu")
	var file_extension string = "fl"
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	var scanner bufio.Scanner = *bufio.NewScanner(file)
	var data string = ""
	for scanner.Scan() {
		data += scanner.Text()
		data += ";"
	}

	var rt runtime.RuntimeResult = lexer.Tokenize(string(data), file_extension)
	if rt.Error != nil {
		runtime.DisplayError(*rt.Error)
		return
	}

	var tokens []lexer.Token = rt.Result.([]lexer.Token)
	rt = parser.ProduceBlock(&tokens, file_extension, false)
	if rt.Error != nil {
		runtime.DisplayError(*rt.Error)
		return
	}

	spew.Dump(rt.Result)

	// var s map[string]value_types.Object = make(map[string]value_types.Object)
	// s["type"] = initializer.MakeType("type")
	// s["int"] = initializer.MakeType("int")
	// s["double"] = initializer.MakeType("double")
	// s["boolean"] = initializer.MakeType("boolean")
	// s["builtin_function"] = initializer.MakeType("builtin_function")
	// s["defined_function"] = initializer.MakeType("defined_function")

	// var c map[string]bool = make(map[string]bool)
	// c["type"] = true
	// c["int"] = true
	// c["double"] = true
	// c["boolean"] = true
	// c["builtin_function"] = true
	// c["defined_function"] = true

	// var scope scopes.Scope = scopes.Scope{
	// 	Scope: s,
	// 	Parent: nil,
	// 	Constants: c,
	// }

	// __, _ := rt.Result.(parser.Statement)
	// rt = interpreter.Evaluate(&__, &scope)
	// if rt.Error != nil {
	// 	runtime.DisplayError(*rt.Error)
	// 	return
	// }

	// spew.Dump(rt.Result)
}