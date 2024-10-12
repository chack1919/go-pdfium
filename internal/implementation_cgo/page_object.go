package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerPageObject(pageObject C.FPDF_PAGEOBJECT) *PageObjectHandle {
	ref := uuid.New()
	handle := &PageObjectHandle{
		handle:    pageObject,
		nativeRef: references.FPDF_PAGEOBJECT(ref.String()),
	}

	p.pageObjectRefs[handle.nativeRef] = handle

	return handle
}
