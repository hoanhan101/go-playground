package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {
	content, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	data := strings.FieldsFunc(string(content[:]), f)

	type KeyValue struct {
		Key   string
		Value string
	}

	var kvs []KeyValue

	for _, value := range data {
		kvs = append(kvs, KeyValue{value, "1"})
	}

	for _, value := range kvs {
		fmt.Println(value)
	}
}
