# ass2inXSS
## intro
   ass2inXSS is a xss scanner which use regexp to analyze the context. 
  After finishing scanning, it will list xss vuln details according to the context. 
  With this, you can get which type the vuln belongs and make decision on which payload to choose.
  
## install
```
go get -u github.com/YuLinSecTeam/ass2inXSS
```
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

## how to test url list?
### url.txt
```
http://brutelogic.com.br/xss.php?a=1
http://brutelogic.com.br/xss.php?b1=1

```
### one liner command
```
cat url.txt | xargs -t -l ass2inXSS -url

you will get:

ass2inXSS -url http://brutelogic.com.br/xss.php?a=1 

Context: html
VulPara: a
VerboseInfo: use normal html tag

ass2inXSS -url http://brutelogic.com.br/xss.php?b1=1 

Context: htmli
VulPara: b1
VerboseInfo: ' or " can use in html tag


Context: htmlo
VulPara: b1
VerboseInfo: '> or "> can break out html tag


```
