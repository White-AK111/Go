package millionfiles

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// Test create a million files
func TestCreateMillionFiles(t *testing.T) {
	dir := "/home/white/MillionFiles"

	assert.Nil(t, CreateMillionFiles(dir), "error on create files")

	files, _ := ioutil.ReadDir(dir)
	assert.Equal(t, 1000000, files, "Incorrect number of files expected 1 000 000 got %d", len(files))
}
