package util

import (
	"encoding/json"
	"fmt"
	"io"
)

func ConvertJSONtoFormData(r io.Reader, d interface{}) {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(d)
	if err != nil {
		fmt.Println(err.Error())
	}
}

