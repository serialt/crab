package crab

import (
	"testing"

	"github.com/serialt/crab/internal"
)

func TestTgz(t *testing.T) {
	src := "testdata/date.txt"
	dest := "testdata/date-test.tgz"
	untgzFile := "testdata/date-test"
	assert := internal.NewAssert(t, "Tgz")

	err := TarGzip(dest, src)
	assert.IsNil(err)
	err = UnTarGzip(dest, untgzFile)
	assert.IsNil(err)

}
