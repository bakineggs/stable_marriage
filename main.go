package main

import "fmt"
import "./bonsai/interpreter"
import "./bonsai/interpreter/techniques"

func main() {
	rules := CompiledRules()
	techniques := []bonsai.Technique{&techniques.Basic{}}

	for _, technique := range techniques {
		technique.Learn(rules)
	}

	interpreter := bonsai.Interpreter{Techniques: techniques}
	root := bonsai.MakeRootNode()
	interpreter.Interpret("", root)
	fmt.Print(root.ToString())
}
