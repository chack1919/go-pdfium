package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerAnnotation(annotation C.FPDF_ANNOTATION) *AnnotationHandle {
	ref := uuid.New()
	handle := &AnnotationHandle{
		handle:    annotation,
		nativeRef: references.FPDF_ANNOTATION(ref.String()),
	}

	p.annotationRefs[handle.nativeRef] = handle

	return handle
}
