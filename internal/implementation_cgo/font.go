package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_edit.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerFont(pageObject C.FPDF_FONT) *FontHandle {
	ref := uuid.New()
	handle := &FontHandle{
		handle:    pageObject,
		nativeRef: references.FPDF_FONT(ref.String()),
	}

	p.fontRefs[handle.nativeRef] = handle

	return handle
}
