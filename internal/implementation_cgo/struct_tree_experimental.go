//go:build pdfium_experimental
// +build pdfium_experimental

package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_structtree.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"

	"github.com/google/uuid"
)

type StructElementAttributeHandle struct {
	handle    C.FPDF_STRUCTELEMENT_ATTR
	nativeRef references.FPDF_STRUCTELEMENT_ATTR // A string that is our reference inside the process. We need this to close the references in DestroyLibrary.
}

func (p *PdfiumImplementation) registerStructElementAttribute(structElementAttribute C.FPDF_STRUCTELEMENT_ATTR) *StructElementAttributeHandle {
	ref := uuid.New()
	handle := &StructElementAttributeHandle{
		handle:    structElementAttribute,
		nativeRef: references.FPDF_STRUCTELEMENT_ATTR(ref.String()),
	}

	p.structElementAttributeRefs[handle.nativeRef] = handle

	return handle
}
