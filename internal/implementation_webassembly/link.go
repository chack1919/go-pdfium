package implementation_webassembly

import (
	"github.com/google/uuid"

	"github.com/chack1919/go-pdfium/references"
)

func (p *PdfiumImplementation) registerLink(dest *uint64) *LinkHandle {
	ref := uuid.New()
	handle := &LinkHandle{
		handle:    dest,
		nativeRef: references.FPDF_LINK(ref.String()),
	}

	p.linkRefs[handle.nativeRef] = handle

	return handle
}
