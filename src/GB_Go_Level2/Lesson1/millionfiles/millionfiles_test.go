package millionfiles

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// Test create a million files
func TestCreateMillionFiles(t *testing.T) {
	var err error
	dir := "/home/white/MillionFiles"

	assert.Equal(t, err, CreateMillionFiles(dir), "error on create files")

	files, _ := ioutil.ReadDir(dir)
	if l := len(files); l != 1000000 {
		t.Errorf("Uncorrect number of files expected 1 000 000 got %d", l)
	}
}
