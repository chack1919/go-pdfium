package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_signature.h"
import "C"
import (
	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerSignature(signature C.FPDF_SIGNATURE, documentHandle *DocumentHandle) *SignatureHandle {
	ref := uuid.New()
	handle := &SignatureHandle{
		handle:      signature,
		nativeRef:   references.FPDF_SIGNATURE(ref.String()),
		documentRef: documentHandle.nativeRef,
	}

	documentHandle.signatureRefs[handle.nativeRef] = handle
	p.signatureRefs[handle.nativeRef] = handle

	return handle
}
