package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type fileProp struct {
	Dir  string
	Name string
	Date string
	Size string
	Hash string
}

func main() {
	var pathCatalogsCsv string = "catalogs.csv"
	var pathImportCsv string = "md5files.csv"

	catalogsPath := getFileFromCatalog(pathCatalogsCsv)
	fmt.Println(catalogsPath)

	writeToCsv(pathImportCsv, getFileHash(catalogsPath))
}

//функция получения списка каталогов из csv файла
func getFileFromCatalog(pathCsv string) map[int]string {
	csvFile, err := os.Open(pathCsv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	catalogsPath := make(map[int]string)
	for i := 0; i < len(csvLines); i++ {
		catalogsPath[i] = csvLines[i][0]
		fmt.Println(catalogsPath[i])
	}

	return catalogsPath
}

//функция получения всех файлов в каталоге
func getFileHash(catalogsPath map[int]string) []fileProp {

	filesProp := make([]fileProp, 0)

	for _, dir := range catalogsPath {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name(), file.IsDir())
			if !file.IsDir() {
				var file = fileProp{
					dir,
					file.Name(),
					file.ModTime().String(),
					strconv.FormatInt(file.Size(), 10),
					getHashMD5(dir + file.Name()),
				}
				filesProp = append(filesProp, file)
			}
		}
	}

	return filesProp
}

//функция получения хэша md5 файла
func getHashMD5(fileName string) string {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(hash.Sum(nil)))

	return hex.EncodeToString(hash.Sum(nil))
}

//функция записи в csv
func writeToCsv(pathImportCsv string, arrFile []fileProp) {

	file, err := os.Create(pathImportCsv)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range arrFile {
		arrValue := []string{value.Dir, value.Name, value.Date, value.Size, value.Hash}
		err := writer.Write(arrValue)
		checkError("Cannot write to file", err)
	}
}

//функция проверки ошибки
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
