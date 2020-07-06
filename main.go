package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"rogchap.com/v8go"
)

func main() {
	ctx, _ := v8go.NewContext(nil)
	content, err := ioutil.ReadFile("value.js")
	if err != nil {
		log.Fatal(err)
	}
	newCont := "let testVar='yas';" + string(content)
	//val, err := ctx.RunScript(string(content), "value.js")
	val, err := ctx.RunScript(string(newCont), "value.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
