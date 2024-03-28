package main

import (
	"fmt"

	"github.com/Mireu-Lab/golang-practice-project/greeting"
	"github.com/Mireu-Lab/golang-practice-project/greeting/en"
	"github.com/Mireu-Lab/golang-practice-project/greeting/ko"
	"github.com/Mireu-Lab/golang-practice-project/src"
)

func main() {
	fmt.Println(greeting.HelloKo())
	fmt.Println(greeting.HelloEn())

	src.Asay()

	fmt.Print(en.Hello())
	fmt.Print(ko.Hello())
}
