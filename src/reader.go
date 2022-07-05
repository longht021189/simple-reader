package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/longht021189/simplereader/src/parser"
)

func Read(file, key, filetype string) {
	if filetype != "json" {
		log.Panicf("Not Support %s", filetype)
	}

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		log.Panicf("Read File Error %v, File: %s", err, file)
	}

	output, err := parser.GetJSON(key, bytes)

	if err != nil {
		log.Panicf("Read File Error %v, File: %s", err, file)
	}

	for _, item := range output {
		value := reflect.ValueOf(item)
		valueKind := value.Kind()

		switch {
		case (valueKind >= reflect.Bool && valueKind <= reflect.Uint64) ||
			(valueKind == reflect.Float32) || (valueKind == reflect.Float64) ||
			(valueKind == reflect.String):
			fmt.Printf("%v\n", value)
		default:
			log.Panicf("Not Support Output Type %s", valueKind)
		}
	}
}
