//Стандартная библиотека
package main

import (
	"bytes"
	"container/list"
	"crypto/sha1"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func server() {
	// слушать порт
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// принятие соединения
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// обработка соединения
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// получение сообщения
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	// соединиться с сервером
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// послать сообщение
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func serverRpc() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func clientRpc() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
	)
}

func main() {

	//Строки
	fmt.Println(
		// true
		strings.Contains("test", "es"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "e"),

		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),

		// == "aaaaa"
		strings.Repeat("a", 5),

		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),

		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),

		// "test"
		strings.ToLower("TEST"),

		// "TEST"
		strings.ToUpper("test"),
	)

	//Иногда нам понадобится работать с бинарными данными. Чтобы преобразовать строку в набор байт (и наоборот), выполните следующие действия:
	//arr := []byte("test")
	//str := string([]byte{'t', 'e', 's', 't'})

	//Ввод / Вывод
	//func Copy(dst Writer, src Reader) (written int64, err error)
	var buf bytes.Buffer
	buf.Write([]byte("test"))

	//Файлы и папки
	file, err := os.Open("test.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return
	}
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	//через ReadFile
	bs2, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str2 := string(bs2)
	fmt.Println(str2)

	//создание файла
	file3, err := os.Create("test3.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return
	}
	defer file3.Close()

	file3.WriteString("test3")

	//содержимое каталога
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	//рекурсивно обойти каталоги
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	//Ошибки
	errNew := errors.New("error message")
	fmt.Println(errNew)

	//Контейнеры и сортировки
	//Список
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	//Сортировка
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	//Хэши и криптография
	//Некриптографические функции
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test3.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)

	//Криптографические хэш-функции
	hSha := sha1.New()
	hSha.Write([]byte("test"))
	bsSha := hSha.Sum([]byte{})
	fmt.Println(bsSha)

	//Серверы
	go server()
	go client()

	//var input string
	//fmt.Scanln(&input)

	//HTTP
	http.HandleFunc("/hello", hello)

	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)

	http.ListenAndServe(":9000", nil)

	//RPC
	go serverRpc()
	go clientRpc()

	//var input string
	//fmt.Scanln(&input)

	//Получение аргументов из командной строки
	// Определение флагов
	maxp := flag.Int("max", 6, "the max value")
	// Парсинг
	flag.Parse()
	// Генерация числа от 0 до max
	fmt.Println(rand.Intn(*maxp))
	//args := flag.Args()

	//Синхронизация примитивов
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
