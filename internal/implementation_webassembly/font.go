package implementation_webassembly

import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerFont(pageObject *uint64) *FontHandle {
	ref := uuid.New()
	handle := &FontHandle{
		handle:    pageObject,
		nativeRef: references.FPDF_FONT(ref.String()),
	}

	p.fontRefs[handle.nativeRef] = handle

	return handle
}
