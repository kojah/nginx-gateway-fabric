package framework

import (
	"errors"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

type failingWriteCloser struct {
	closed bool
}

func (*failingWriteCloser) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func (w *failingWriteCloser) Close() error {
	w.closed = true
	return nil
}

func TestNewCSVWriterClosesDestinationWhenHeaderWriteFails(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	destination := &failingWriteCloser{}

	writer, err := newCSVWriter(destination, strings.Repeat("x", 8192))

	g.Expect(err).To(MatchError("write failed"))
	g.Expect(writer).To(BeNil())
	g.Expect(destination.closed).To(BeTrue())
}
