package ywbtest
import "fmt"

func Struct1(){
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}
	for index, file := range files {
		fmt.Println(index)
		fmt.Println(file)
	}
}
