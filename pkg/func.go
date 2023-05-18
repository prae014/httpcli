package pkg

import (
	"fmt"
	"io"
	"net/http"
)

//func Get(url string) (proto string, status_code int, header http.Header, body []byte, req_url *http.Request) {
//	c := http.Client{}
//
//	resp, err := c.Get(url)
//	if err != nil {
//		fmt.Printf("Error %s", err)
//		return
//	}
//	defer resp.Body.Close()
//	body_byte, err := io.ReadAll(resp.Body)
//	//fmt.Printf("Body : %s\n has type %T", body, body)
//
//	return resp.Proto, resp.StatusCode, resp.Header, body_byte, resp.Request
//}

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

// TODO: add json as an argument since this is post command
// TODO: add json body flag this will validate JSON input whether it's correct or not
func Post(url string, queries []string) (proto string, status_code int, header http.Header, body []byte) {
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
