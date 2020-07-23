# SPooF
A simple and fast application written in Go to check for domains vulnerable to email spoofing (missing SPF)
Accepts line-delimited domains on stdin, checks for missing SPF TXT record and output them on stdout.

Usage example:

```
▶ cat domains.txt | SPooF
```

Domains missing SPF TXT record are printed in red.  
