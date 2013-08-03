package processor

import (
	"bytes"
	"github.com/chdorner/imagine/instructions"
	"github.com/gosexy/canvas"
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
	instr.Action = "crop"
	instr.Width = 300
	instr.Height = 300

	p := NewProcessor(instr)
	b := bytes.NewBuffer(nil)
	p.Process(f, b)

	data, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	actual := canvas.New()
	err = actual.OpenBlob(data, uint(len(data)))
	if err != nil {
		t.Fatal(err)
	}
	defer actual.Destroy()

	if actual.Width() != 300 || actual.Height() != 300 {
		t.Fatalf("cropped image is not as expected, got: %dx%d", actual.Width(), actual.Height())
	}
	}
}
