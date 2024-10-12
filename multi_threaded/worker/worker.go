package worker

import (
	"github.com/chack1919/go-pdfium"
	"github.com/chack1919/go-pdfium/internal/implementation_cgo"
)

func StartWorker(config *pdfium.LibraryConfig) {
	implementation_cgo.StartPlugin(config)
}
