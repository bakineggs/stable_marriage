package main

import "fmt"
import "./bonsai/interpreter"

func main() {
	rules := []bonsai.Rule{}
	techniques := []bonsai.Technique{}

	for _, technique := range techniques {
		technique.Learn(rules)
	}

	interpreter := bonsai.Interpreter{Techniques: techniques}
	var root bonsai.Node
	interpreter.Interpret("", &root)
	fmt.Print(root.ToString())
}