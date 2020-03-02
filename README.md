# ass2inXSS
## intro
   ass2inXSS is a xss scanner which use regexp to analyze the context. 
  After finishing scanning, it will list xss vuln details according to the context. 
  With this, you can get which type the vuln belongs and make decision on which payload to choose.
  

## usage
  
```
  Usage of ass2inXSS:
  -C value
        -C a=1;b=2
  -D string
        -D "a=1&b=2"
  -H value
        -H "x-forward-for:127.0.0.1"
  -X string
        -X GET (default "GET")
  -timeout int
        -timeout 2 (default 5)
  -url string
        -u example.com

```

## result
```
Context: html
VulPara: a
VerboseInfo: use normal html tag
```
