package parser

import (
	"encoding/json"
	"reflect"
	"strings"
)

func (j *jsonParser) Parse() error {
	j.keyParts = strings.Split(j.key, ".")

	var raw interface{}
	err := json.Unmarshal(j.data, &raw)
	if err != nil {
		return err
	}

	return j.parse(0, raw)
}

func (j *jsonParser) parse(index int, raw interface{}) error {
	if index >= len(j.keyParts) {
		j.output = append(j.output, raw)
		return nil
	}

	value := reflect.ValueOf(raw)
	valueKind := value.Kind()

	switch valueKind {
	case reflect.Array, reflect.Slice:
		length := value.Len()
		for i := 0; i < length; i += 1 {
			err := j.parse(index, value.Index(i).Interface())
			if err != nil {
				return err
			}
		}

	case reflect.Map:
		currentKey := j.keyParts[index]
		item := value.MapIndex(reflect.ValueOf(currentKey))
		if item.IsValid() {
			err := j.parse(index+1, item.Interface())
			if err != nil {
				return err
			}
		}

	default:
		return ErrorTypeNotMatch
	}

	return nil
}
