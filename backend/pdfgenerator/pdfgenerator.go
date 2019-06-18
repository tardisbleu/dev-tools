package pdfgenerator

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

// PdfGenerator is our bound PdfGenerator Struct
type PdfGenerator struct {
}

func generatePDF(pNbPDF int) {
	for index := 0; index < pNbPDF; index++ {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "PDF numéro : "+strconv.Itoa(index+1))
		pdf.OutputFileAndClose("pdf-" + strconv.Itoa(index+1) + ".pdf")
	}
}
