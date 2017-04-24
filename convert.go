package goutils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

func ToFloat(v interface{}) float64 {
	f, err := strconv.ParseFloat(ToString(v), 64)
	if err != nil {
		return 0
	} else {
		return f
	}
}

func ToString(v interface{}) (s string) {
	switch v.(type) {
	case nil:
		return ""
	case string:
		s = v.(string)
	case []byte:
		s = string(v.([]byte))
	case io.Reader:
		b, _ := ioutil.ReadAll(v.(io.Reader))
		s = string(b)
	case error:
		s = v.(error).Error()
	default:
		b, err := json.Marshal(v)
		if err == nil {
			s = string(b)
		} else {
			s = fmt.Sprintf("%s", b)
		}
	}
	return
}

func ToInt(v interface{}) int64 {
	i, err := strconv.Atoi(ToString(v))
	if err != nil {
		return 0
	} else {
		return int64(i)
	}
}
