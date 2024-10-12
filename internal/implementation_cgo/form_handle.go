package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_formfill.h"
import "C"
import (
	"unsafe"

	"github.com/chack1919/go-pdfium/references"
	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerFormHandle(formHandle C.FPDF_FORMHANDLE, formInfo unsafe.Pointer) *FormHandleHandle {
	ref := uuid.New()
	handle := &FormHandleHandle{
		handle:           formHandle,
		nativeRef:        references.FPDF_FORMHANDLE(ref.String()),
		formInfo:         formInfo,
		pagePointers:     map[unsafe.Pointer]references.FPDF_PAGE{},
		documentPointers: map[unsafe.Pointer]references.FPDF_DOCUMENT{},
	}

	p.formHandleRefs[handle.nativeRef] = handle

	return handle
}
