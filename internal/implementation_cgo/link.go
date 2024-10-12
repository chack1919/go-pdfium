package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerLink(dest C.FPDF_LINK) *LinkHandle {
	ref := uuid.New()
	handle := &LinkHandle{
		handle:    dest,
		nativeRef: references.FPDF_LINK(ref.String()),
	}

	p.linkRefs[handle.nativeRef] = handle

	return handle
}
