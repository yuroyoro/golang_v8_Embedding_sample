//
// Embedding V8 Javascript Engine and Go | Brave New Method http://bravenewmethod.com/2011/03/30/embedding-v8-javascript-engine-and-go/
// Embedder's Guide - Chrome V8 — Google Developers https://developers.google.com/v8/embed?csw=1
// Hello,  C++ + JavaScript World! - Boost.勉強会 #11 東京 http://www.slideshare.net/hecomi/hello-cppandjsworld
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	script, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}

	var result interface{}
	RunV8(string(script), &result)
	fmt.Printf("Result -> %T: %#+v\n", result, result)

}
