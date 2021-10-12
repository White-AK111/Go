// Package generator
// Task#3
// *не обязательное*. Написать кодогенератор под какую-нибудь задачу.
//go:generate go run generator.go MyInt

// This program generates {{MyType}}_queue.go. It can be invoked by running
// go generate
package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

var tpl = `package {{.Package}}
	
	type {{.MyType}}Queue struct {
		q []{{.MyType}}
	}

	func New{{.MyType}}Queue() *{{.MyType}}Queue {
		return &{{.MyType}}Queue{
			q: []{{.MyType}}{},
		}
	}

	func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
		o.q = append(o.q, v)
	}
`

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)

		if err != nil {
			fmt.Printf("Could not create %s: %s (skip) \n", dest, err)
			continue
		}
		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}

		tt.Execute(file, vals)
		file.Close()
	}
}
