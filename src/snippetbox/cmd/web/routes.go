package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	mux := http.NewServeMux()
	// Используем методы из структуры в качестве обработчиков маршрутов.
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux.HandleFunc("/", app.home)
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в маршрутизаторе servemux
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	//fileServer := http.FileServer(http.Dir("../../ui/static/"))
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("../../ui/static/")})
	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		//Если делать как в примере, то формируется некорректный путь + некорректно обрабатывается ошибка.
		//И в тоге web-сервер падает в ошибку.
		//index := filepath.Join(path, "index.html")
		index := "index.html"
		if _, err := nfs.fs.Open(index); err == nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
		} else {
			return nil, err
		}
	}

	return f, nil
}
