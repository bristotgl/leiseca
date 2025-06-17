package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
	
	req, err := http.NewRequest("GET", "https://www.planalto.gov.br/ccivil_03/leis/l8112cons.htm", nil)
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	fmt.Println(resp)

	body, err := io.ReadAll(resp.Body)
		if err != nil {
		fmt.Println("Body read failed:", err)
		return
	}

	html := string(body)
	fmt.Println(html)
}