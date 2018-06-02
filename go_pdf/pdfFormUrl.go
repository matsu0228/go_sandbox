package main

import (
	"fmt"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	url := "https://google.com/"

	pdfg.AddPage(wkhtmltopdf.NewPage(url))

	// PDF作成
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// 出力
	err = pdfg.WriteFile("./google.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success to create pdf")
}
