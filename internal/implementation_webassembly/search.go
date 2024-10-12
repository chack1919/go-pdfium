package implementation_webassembly

import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerSearch(search *uint64, documentHandle *DocumentHandle) *SearchHandle {
	ref := uuid.New()
	handle := &SearchHandle{
		handle:      search,
		nativeRef:   references.FPDF_SCHHANDLE(ref.String()),
		documentRef: documentHandle.nativeRef,
	}

	documentHandle.searchRefs[handle.nativeRef] = handle
	p.searchRefs[handle.nativeRef] = handle

	return handle
}
