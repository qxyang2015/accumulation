package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func main() {
	paths := []string{"name.first", "age", "friends"}
	values := gjson.GetMany(json, paths...)
	for _, v := range values {
		fmt.Println(v.Type, v.Raw)
		if v.Type == gjson.JSON {
			cv := v.Get("#.first")
			fmt.Println("cv", cv.Type, cv.Raw)
		}
	}
}
