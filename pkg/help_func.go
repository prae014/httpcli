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
		//split flags to ["key", "val"}, using "="
		split_str := strings.Split(flag, "=")
		query.Add(split_str[0], split_str[1])
	}
	req.URL.RawQuery = query.Encode() // RawQuery (from URL struct) encoded query values

	return req.URL.String()
}

// check if string is in json form or not
func Validjson(s string) (json_fmt []byte) {
	input := []byte(s)

	decoder := json.NewDecoder(bytes.NewBuffer(input))
	for {
		_, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("invalid input: input is not in JSON format")
			panic(err) // in case of validating a user input, we don't use the panic() function
			// instead we should return it to outside and let caller handle it
			// panic() is used only if it want to throw an error from developers
		}
	}
	return input
}
