package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Обработчик главной страницы.
// Меняем сигнатуры обработчика home, чтобы он определялся как метод
// структуры *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		//http.NotFound(w, r)
		return
	}

	// Инициализируем срез содержащий пути к файлам шаблонов. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"../../ui/html/home.page.tmpl",
		"../../ui/html/base.layout.tmpl",
		"../../ui/html/footer.partial.tmpl",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Поскольку обработчик home теперь является методом структуры application
		// он может получить доступ к логгерам из структуры.
		// Используем их вместо стандартного логгера от Go.
		//log.Println(err.Error())
		app.errorLog.Println(err.Error())
		app.serverError(w, err) // Использование помощника serverError()
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		// Обновляем код для использования логгера-ошибок
		// из структуры application.
		//log.Println(err.Error())
		app.errorLog.Println(err.Error())
		app.serverError(w, err) // Использование помощника serverError()
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Обработчик для отображения содержимого заметки.
// Меняем сигнатуру обработчика showSnippet, чтобы он был определен как метод
// структуры *application
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Использование помощника notFound()
		//http.NotFound(w, r)
		return
	}

	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
	// и записываем его в http.ResponseWriter.
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
	//w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
// Меняем сигнатуру обработчика createSnippet, чтобы он определялся как метод
// структуры *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
	// что http.MethodPost является строкой и содержит текст "POST".
	if r.Method != http.MethodPost {
		// Если это не так, то
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		/*
			//через WriteHeader
			//вызывается метод w.WriteHeader() для возвращения статус-кода 405
			// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
			w.WriteHeader(405)
			w.Write([]byte("GET-Метод запрещён!"))
		*/

		//через Error
		app.clientError(w, http.StatusMethodNotAllowed) // Используем помощник clientError()
		//http.Error(w, "Метод запрещён!", http.StatusMethodNotAllowed)

		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}
