package main

import (
	"cmd-inject-header/core"
	"flag"
	"fmt"
	"os"
)

func main() {
	flagCheck := flag.NewFlagSet("cmd-inject-header", flag.ExitOnError)
	hfName := flagCheck.String("hd", "", "Path to list header file")
	urName := flagCheck.String("ur", "", "Path to list url file (Sample URL: http://example.com/)")
	plName := flagCheck.String("pl", "", "Path to list payload file")
	hnName := flagCheck.String("it", "", "Your interact server to check the interaction (dnslog.cn/burp collabarator/interact.sh...)")
	flagCheck.Parse(os.Args[1:])
	if *hfName == "" || *plName == "" || *urName == "" || *hnName == "" {
		fmt.Printf("Usage : go run . -hd headers.txt -pl payloads.txt -ur urls.txt -it xxx.burpcollaborator.net")
	} else {
		core.Url_Execute(*hfName, *plName, *urName, *hnName)
	}
}
