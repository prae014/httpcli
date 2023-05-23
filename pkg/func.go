package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string, queries []string) (proto string, status_code int, header http.Header, body []byte) {
	c := http.Client{
		Timeout: time.Second * 90, // we always add a timeout value for a http client to prevent request hanging
	}

	req, err := http.NewRequest("GET", url, nil)
	// always handle the error by returning the error out of the function if err is not nil

	//modify url if queries exist
	new_url := URLmod(req, queries)

	resp, err := c.Get(new_url)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close() // good practice, always close the body to prevent a leak (i forgot what leak but p copter told me so)
	body_byte, err := io.ReadAll(resp.Body)
	//fmt.Printf("Body : %s\n has type %T", body, body)

	return resp.Proto, resp.StatusCode, resp.Header, body_byte
}

// Please not that in GO, we handle the error by returning it.
// Usually the error will be handle by the caller, so we don't need to worry about doing any action when the error happens inside a function
// For example, the my Get function would look thing like this:

func Get2(url string, queries []string) (proto string, status_code int, header http.Header, body []byte, err error) { // add err as a part of returning values {
	c := http.Client{
		Timeout: time.Second * 90, // we always add a timeout value for a http client to prevent request hanging
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// return the nil value for every variable along with the error produced
		return "", -1, http.Header{}, []byte{}, err
	}

	//modify url if queries exist
	new_url := URLmod(req, queries)

	resp, err := c.Get(new_url)
	if err != nil {
		// therefore instead of handling the error by printing it here, we can return the error and let the caller decided what to do
		return "", -1, http.Header{}, []byte{}, err
	}
	defer resp.Body.Close()
	body_byte, err := io.ReadAll(resp.Body)
	//fmt.Printf("Body : %s\n has type %T", body, body)

	return resp.Proto, resp.StatusCode, resp.Header, body_byte, nil // if we catches no error, we return it as a nil value
}

// Then in the caller code,
// proto, status_code, header, body, err := pkg.Put(args[0], query_flags, json_flags)
// if err != nil {
// 	fmt.Print(err.Error())
// }

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
