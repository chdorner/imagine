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

func TestProcessShrinkWidth(t *testing.T) {
	f, err := os.Open("../test/rectangle.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	instr := &instructions.RequestInstructions{}
	instr.Action = "shrink-w"
	instr.Width = 200

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

	if actual.Width() != 200 || actual.Height() != 300 {
		t.Fatalf("width shrinked image is not 200x300, got: %dx%d", actual.Width(), actual.Height())
	}
}

func TestProcessShrinkWidthBigger(t *testing.T) {
	f, err := os.Open("../test/rectangle.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	instr := &instructions.RequestInstructions{}
	instr.Action = "shrink-w"
	instr.Width = 500

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

	if actual.Width() != 400 || actual.Height() != 600 {
		t.Fatalf("width shrinked image should not have changed, got: %dx%d", actual.Width(), actual.Height())
	}
}
