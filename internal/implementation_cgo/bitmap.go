package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerBitmap(bitmap C.FPDF_BITMAP) *BitmapHandle {
	ref := uuid.New()
	handle := &BitmapHandle{
		handle:    bitmap,
		nativeRef: references.FPDF_BITMAP(ref.String()),
	}

	p.bitmapRefs[handle.nativeRef] = handle

	return handle
}
