package main

import (
	"cmd-inject-header/core"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	flagCheck := flag.NewFlagSet("cmd-inject-header", flag.ExitOnError)
	hdName := flagCheck.String("hd", "", "Path to list header file")
	urName := flagCheck.String("ur", "", "Path to list url file (URL have / in the end: http://example.com/)")
	plName := flagCheck.String("pl", "", "Path to list payload file")
	itServer := flagCheck.String("it", "", "Your interact server to check the interaction (dnslog.cn/burp collabarator/interact.sh...)")
	stterror := flagCheck.Bool("se", false, "If you want send request for status code 4xx/5xx")
	flagCheck.Parse(os.Args[1:])
	if *hdName == "" || *plName == "" || *urName == "" || *itServer == "" {
		fmt.Println("Usage : go run . -hd headers_common.txt -pl payloads.txt -ur urls.txt -it xxx.burpcollaborator.net")
		fmt.Println("Usage : go run . -hd headers_common.txt -pl payloads.txt -ur urls.txt -it xxx.burpcollaborator.net -se=true")
		fmt.Print("(To send request for status code 4xx/5xx. Sometime app only vuln with this status code!)")
	} else {
		start := time.Now()
		urls := core.ReadFromFile(*urName)
		headers := core.ReadFromFile(*hdName)
		_payloads := core.ReadFromFile(*plName)
		var payloads []string
		for _, payload := range _payloads {
			payload := strings.Replace(payload, "INTERACT_SERVER", *itServer, -1)
			payloads = append(payloads, payload)
		}
		total := len(payloads)*len(headers) * len(urls)
		fmt.Println("Total request will send: ", total)
		ch := make(chan string)
		for _, url := range urls {
			if *stterror {
				go core.MakeRequestHeader(url + "toMakeError", headers, payloads, ch)
			} else {
				go core.MakeRequestHeader(url, headers, payloads, ch)
			}
		}
		for range urls {
			fmt.Println(<-ch)
		}
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}
