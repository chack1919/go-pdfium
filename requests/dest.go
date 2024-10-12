package requests

import "github.com/chack1919/go-pdfium/references"

type GetDestInfo struct {
	Document references.FPDF_DOCUMENT
	Dest     references.FPDF_DEST
}
