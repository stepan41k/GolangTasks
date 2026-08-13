// Условие:
// В хранилище Диска поступает список абсолютных путей к файлам. Напиши функцию GroupFilesByDirectory(paths []string) map[string][]string, которая сгруппирует имена файлов по их директориям.

// Ключ мапы — путь к директории (без слэша на конце, кроме корневой директории /).
// Значение — слайс имен файлов, содержащихся в этой директории.
// Порядок файлов в слайсе должен соответствовать порядку из исходного массива paths.
// Для работы с путями можно использовать стандартный пакет path/filepath или strings.

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GroupFilesByDirectory(files []string) map[string][]string {
	res := map[string][]string{}
	
	for _, file := range files {
		dir := filepath.Dir(file)
		dir = strings.ReplaceAll(dir, `\`, "/")
		res[dir] = append(res[dir], filepath.Base(file))
	}

	return res
}

func main() {
	paths := []string{
		"/var/log/nginx/access.log",
		"/var/log/nginx/error.log",
		"/var/log/syslog",
		"/home/user/document.pdf",
		"/file_in_root.txt",
	}

	res := GroupFilesByDirectory(paths)

	for k, v := range res {
		fmt.Println(k, v)
	}
}
