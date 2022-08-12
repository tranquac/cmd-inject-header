package core

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{
	Transport: transport,
}

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: time.Second,
		DualStack: true,
	}).DialContext,
}

func ReadFromFile(filename string) []string {
	url_file, err := os.Open(filename)
	if err != nil {
		log.Fatal("File could not be read\n")
	}
	defer url_file.Close()
	uScanner := bufio.NewScanner(url_file)
	var urls []string
	for uScanner.Scan() {
		urls = append(urls, uScanner.Text())
	}
	if err := uScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return urls
}

func MakeRequestHeader(url string, headers, payloads []string, ch chan<- string) {
	chh := make(chan string)
	for _, header := range headers {
		go MakeRequestPayload(url, header, payloads, chh) // Using all payload with each header
	}
	for range headers {
		fmt.Println(<-chh)
	}
	ch <- fmt.Sprintf("Done Url %s", url)
}

func MakeRequestPayload(url, header string, payloads []string, ch chan<- string) { // Make request with each payload in payload
	chh := make(chan string)
	for _, payload := range payloads {
		go MakeRequestFinal(url, header, payload, chh) // Using all payload with each header
	}
	for range payloads {
		fmt.Println(<-chh)
	}
	ch <- fmt.Sprintf("Done Header %s", header)
}

func MakeRequestFinal(url, header, payload string, ch chan<- string) {
	start := time.Now()
	// resp, _ := http.Get(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	if header != "User-Agent" {
		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36")
	}
	req.Header.Add(header, payload)
	resp, err := httpClient.Do(req)
	if err == nil {
		secs := time.Since(start).Seconds()
		status := resp.StatusCode
		ch <- fmt.Sprintf("%.2f elapsed for URL: %s | StatusCode: %d | Header: %s | Payload: %s", secs, url, status, header, payload)
	} else {
		ch <- fmt.Sprintf("%v", err)
	}
}
