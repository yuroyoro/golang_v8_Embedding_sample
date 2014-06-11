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

	script := `obj = {"foo": [1, 2], "bar": {"baz": true, "hoge": "fuga"}}`
	var result Foo

	RunV8(string(script), &result)
	fmt.Printf("Result -> %T: %#+v\n", result, result)

}
