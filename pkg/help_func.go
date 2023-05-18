package pkg

import (
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
