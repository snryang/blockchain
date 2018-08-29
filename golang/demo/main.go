package main
import (
	"bytes"
	"fmt"
	"unicode"
	"strings"
)

func main(){
	var ch byte
	fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
	fmt.Scanf("%c\n",&ch)
	var buffer bytes.Buffer
	err := buffer.WriteByte(ch)
	if err== nil{
		fmt.Println("写入一个字节成功，准备读取该字节")
		newCH,_ :=buffer.ReadByte()
		fmt.Printf("读取的字节: %c\n",newCH)
	}else{
		fmt.Println("写入错误");
	}
}