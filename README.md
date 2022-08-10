## cmd-inject-header
A simple tool to check command injection in headers of http request

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
  -ur string
        Path to list url file
```
Examples: `go run . -hd h.txt -pl payloads.txt -ur urls.txt -it xxx.burpcollaborator.net`
