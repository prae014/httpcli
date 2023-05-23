package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Get(url string, queries []string) (proto string, status_code int, header http.Header, body []byte) {
	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	//modify url if queries exist
	new_url := URLmod(req, queries)

	resp, err := c.Get(new_url)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body_byte, err := io.ReadAll(resp.Body)
	//fmt.Printf("Body : %s\n has type %T", body, body)

	return resp.Proto, resp.StatusCode, resp.Header, body_byte
}

func Post(url string, queries []string, input string) (proto string, status_code int, header http.Header, body []byte) {
	c := http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	//modify url if queries exist
	new_url := URLmod(req, queries)

	//check whether the input is in correct format(json) or not
	json_input := Validjson(input)

	resp, err := c.Post(new_url, "application/json", bytes.NewBuffer(json_input))
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body_byte, err := io.ReadAll(resp.Body)
	//fmt.Printf("Body : %s\n has type %T", body, body)

	return resp.Proto, resp.StatusCode, resp.Header, body_byte
}

func Delete(url string, queries []string) (proto string, status_code int, header http.Header, body []byte) {
	c := http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	//modify url if queries exist
	new_url := URLmod(req, queries)

	mod_req, err := http.NewRequest("DELETE", new_url, nil)

	resp, err := c.Do(mod_req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body_byte, err := io.ReadAll(resp.Body)
	//fmt.Printf("Body : %s\n has type %T", body, body)

	return resp.Proto, resp.StatusCode, resp.Header, body_byte
}

func Put(url string, queries []string, input string) (proto string, status_code int, header http.Header, body []byte) {
	c := http.Client{}

	req, err := http.NewRequest("PUT", url, nil)
	//modify url if queries exist
	new_url := URLmod(req, queries)

	//check whether the input is in correct format(json) or not
	json_input := Validjson(input)
	//fmt.Printf("json input is %v\n", string(json_input))

	mod_req, err := http.NewRequest("PUT", new_url, bytes.NewBuffer(json_input))

	resp, err := c.Do(mod_req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body_byte, err := io.ReadAll(resp.Body)

	return resp.Proto, resp.StatusCode, resp.Header, body_byte
}
