package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// modify URL if queries exist
func URLmod(req *http.Request, q_flags []string) (new_url string) {
	query := req.URL.Query()
	for _, flag := range q_flags {
		//split flags to ["key", "val"}, using = as a delimiter
		split_str := strings.Split(flag, "=")
		query.Add(split_str[0], split_str[1])
	}
	req.URL.RawQuery = query.Encode() // RawQuery (from URL struct) encoded query values

	return req.URL.String()
}

// check if string is in json form or not
func Validjson(s string) (json_fmt []byte) {
	input := []byte(s) //TODO: have to check if json format we send is []byte???? is this right??

	decoder := json.NewDecoder(bytes.NewBuffer(input))
	for {
		_, err := decoder.Token()
		//fmt.Println(i) //FIXME: test 123 is correct???
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("invalid input: input is not in JSON format")
			panic(err)
		}
	}
	return input
}
