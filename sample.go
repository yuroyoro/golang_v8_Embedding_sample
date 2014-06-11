package main

import (
	"fmt"
)

type Foo struct {
	Foo []int `json:"foo"`
	Bar Bar   `json:"bar"`
}

type Bar struct {
	Baz  bool   `json:"baz"`
	Hoge string `json:"hoge"`
}

func main() {

	scripts := []string{
		"null",
		"true",
		"123",
		"457.78",
		"[10, 20, 30]",
		"'Hello, World'",
		"new Date()",
		`obj = {"foo": [1, 2], "bar": {"baz": true, "hoge": "fuga"}}`,
	}

	for _, s := range scripts {
		var res interface{}
		RunV8(s, &res)
		fmt.Printf("Script -> %s\n", s)
		fmt.Printf("Result -> %T: %#+v\n\n", res, res)
	}

	script := `obj = {"foo": [1, 2], "bar": {"baz": true, "hoge": "fuga"}}`
	var result Foo

	RunV8(script, &result)
	fmt.Printf("Script -> %s\n", script)
	fmt.Printf("Result -> %T: %#+v\n", result, result)

}
