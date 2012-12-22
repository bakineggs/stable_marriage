package main

import "fmt"
import "./bonsai/interpreter"

func main() {
	rules := CompiledRules()
	techniques := []bonsai.Technique{}

	for _, technique := range techniques {
		technique.Learn(rules)
	}

	interpreter := bonsai.Interpreter{Techniques: techniques}
	root := bonsai.MakeRootNode()
	interpreter.Interpret("", root)
	fmt.Print(root.ToString())
}
