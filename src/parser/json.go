package parser

type jsonParser struct {
	keyParts []string
	key      string
	data     string
	output   []interface{}
}

func GetJSON(key, data string) ([]interface{}, error) {
	j := &jsonParser{
		key:  key,
		data: data,
	}

	err := j.Parse()

	if err != nil {
		return nil, err
	}

	return j.output, nil
}
