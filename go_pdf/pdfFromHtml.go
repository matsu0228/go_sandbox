package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	html := `
    <html>
    <head>
        <meta charset="UTF-8">
    </head>
    <body>
    <h1>こんにちは!</h1>
    <div>hello pdf</dic>
    </body>
    </html>`

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	// PDF作成
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// 出力
	err = pdfg.WriteFile("./test.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tada!")
}
