package requests

import "github.com/chack1919/go-pdfium/references"

type GetActionInfo struct {
	Document references.FPDF_DOCUMENT
	Action   references.FPDF_ACTION
}
