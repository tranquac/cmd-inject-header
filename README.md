## cmd-inject-header
A simple tool to check command injection in headers of http request, faster with goroutines

Designed to make easy check command injection in headers for bug hunter, pentester, red team-er

### Inspired by
https://twitter.com/p3n73st3r/status/1556645395928866818

![https://raw.githubusercontent.com/tranquac/cmd-inject-header/master/image/insprired.png](https://raw.githubusercontent.com/tranquac/cmd-inject-header/master/image/insprired.png)
### Usage
```
Usage of cmd-inject-header:
  -hd string
        Path to list header file
  -it string
        Your interact server to check the interaction (dnslog.cn/burp collabarator/interact.sh...)
  -pl string
        Path to list payload file
  -se
        If you want send request for status code 4xx/5xx
  -ur string
        Path to list url file (URL have / in the end: http://example.com/)
```
Usage : `go run . -hd headers_common.txt -pl payloads.txt -ur urls.txt -it xxx.burpcollaborator.net`

Usage : `go run . -hd headers_common.txt -pl payloads.txt -ur urls.txt -it xxx.burpcollaborator.net -se=true`

(To send request for status code 4xx/5xx. Sometime app only vuln with this status code!)
