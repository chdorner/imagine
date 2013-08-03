package processor

import (
	"bytes"
	"github.com/chdorner/imagine/instructions"
	"io/ioutil"
	"os"
	"testing"
)

func TestProcessCrop(t *testing.T) {
	f, err := os.Open("../test/square.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	instr := &instructions.RequestInstructions{}
	instr.Width = 300
	instr.Height = 300

	p := NewProcessor(instr)
	b := bytes.NewBuffer(nil)
	p.Process(f, b)

	actual, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile("../test/expected_crop.jpg")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		ioutil.WriteFile("../actual_crop.jpg", actual, 0664)
		t.Fatal("cropped image is not as expected, got actual_crop.jpg expected test/expected_crop.jpg")
	}
}
