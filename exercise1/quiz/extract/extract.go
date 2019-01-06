package extract

import (
	"io"
)

// QAExtractor interface that models the operation of retrieving QAs from a reader
type QAExtractor interface {
	Extract(r io.Reader) ([]byte, error)
}
