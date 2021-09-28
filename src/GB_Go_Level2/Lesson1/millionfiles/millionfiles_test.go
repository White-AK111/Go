package millionfiles

import (
	"fmt"
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

func ExampleCreateMillionFiles() {
	dir := "/home/white/MillionFiles"
	err := CreateMillionFiles(dir)
	if err != nil {
		fmt.Println(err)
	}
	// Output: one million empty "*.txt" files in directory "dir" or error if it exists.
}
