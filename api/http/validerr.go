package http

import (
	"net/http"
	"sort"
)

type validateError map[string]string

func (v validateError) Error() string {
	keys := make([]string, 0, len(v))
	for key := range v {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var str string
	for i, key := range keys {
		if i > 0 {
			str += "\n"
		}

		str += key + " " + v[key]
	}

	return str
}

func (validateError) Code() int {
	return http.StatusUnprocessableEntity
}

func (v validateError) Message() string {
	return v.Error()
}
