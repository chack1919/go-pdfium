package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerPageRange(pageRange C.FPDF_PAGERANGE, documentHandle *DocumentHandle) *PageRangeHandle {
	ref := uuid.New()
	handle := &PageRangeHandle{
		handle:      pageRange,
		nativeRef:   references.FPDF_PAGERANGE(ref.String()),
		documentRef: documentHandle.nativeRef,
	}

	p.pageRangeRefs[handle.nativeRef] = handle
	documentHandle.pageRangeRefs[handle.nativeRef] = handle

	return handle
}
