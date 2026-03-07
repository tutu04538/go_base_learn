package main

import "github.com/tidwall/gjson"

const json = `{"name":"tutu"}`

func main() {
	value := gjson.Get(json, "name")
	println(value.String())
}
