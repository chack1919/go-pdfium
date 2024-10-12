package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_edit.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerGlyphPath(glyphPath C.FPDF_GLYPHPATH) *GlyphPathHandle {
	ref := uuid.New()
	handle := &GlyphPathHandle{
		handle:    glyphPath,
		nativeRef: references.FPDF_GLYPHPATH(ref.String()),
	}

	p.glyphPathRefs[handle.nativeRef] = handle

	return handle
}
