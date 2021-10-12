// Package assignreflect
// Task#2
// Написать функцию, которая принимает на вход имя файла и название функции.
// Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
// Результат работы должен возвращать количество вызовов int и ошибку error.
// Разрешается использовать только go/parser, go/ast и go/token.
package astparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
)

// CountGoroutines count go func in target file and function
func CountGoroutines(fileName string, funcName string) (res int, err error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return 0, err
	}

	for _, f := range node.Decls {
		// check only func
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// inspect target function on "go" declarations
		if fn.Name.Name == funcName {
			ast.Inspect(fn, func(n ast.Node) bool {
				_, ok := n.(*ast.GoStmt)
				if ok {
					res++
					return true
				}
				return true
			})
		}
	}

	return res, nil
}

// ExecGoroutines 7 goroutines and send in result channel
func ExecGoroutines() error {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	resCh := make(chan int, 7)

	wg.Add(7)
	go start(1, resCh, &wg)
	go start(2, resCh, &wg)
	go start(3, resCh, &wg)
	go start(4, resCh, &wg)
	go start(5, resCh, &wg)
	go start(6, resCh, &wg)
	go start(7, resCh, &wg)

	for i := 1; i <= 7; i++ {
		fmt.Println(<-resCh)
	}
	return nil
}

// start function start goroutine
func start(i int, ch chan int, wg *sync.WaitGroup) {
	ch <- i
	wg.Done()
}
