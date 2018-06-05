# pdf libraries of wkhtmltopdf

* download binary from following
  * https://github.com/wkhtmltopdf/wkhtmltopdf/blob/master/docs/downloads.md
```
wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

xz -dv wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
tar xfv wkhtmltox-0.12.4_linux-generic-amd64.tar 
```

* set link your path

* test following command
```
which wkhtmltopdf

wkhtmltopdf http://google.com google.pdf
```

* go library
```
go get -u github.com/SebastiaanKlippert/go-wkhtmltopdf
```
