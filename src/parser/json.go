package parser

type jsonParser struct {
	keyParts []string
	key      string
	data     []byte
	output   []interface{}
}

func GetJSON(key string, data []byte) ([]interface{}, error) {
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
