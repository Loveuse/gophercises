package extract

import (
	"io"

	"github.com/Loveuse/gophercises/exercise1/quiz"
)

// Extractor interface that models the actvity of extracting/retrieving data from a reader.
type Extractor interface {
	Extract(r io.Reader) ([]quiz.QA, error)
}
