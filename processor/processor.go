package processor

import (
	"github.com/chdorner/imagine/instructions"
	"github.com/gosexy/canvas"
	"io"
	"io/ioutil"
	"math"
)

type Processor struct {
	instr *instructions.RequestInstructions
	image *canvas.Canvas
}

func NewProcessor(instr *instructions.RequestInstructions) *Processor {
	return &Processor{instr, nil}
}

func (p *Processor) Process(r io.Reader, w io.Writer) error {
	p.image = canvas.New()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = p.image.OpenBlob(data, uint(len(data)))
	if err != nil {
		return err
	}
	defer p.image.Destroy()

	switch p.instr.Action {
	case "crop":
		p.crop()
	case "shrink-w":
		p.shrinkWidth()
	}

	data, err = p.image.GetImageBlob()
	if err != nil {
		return err
	}

	w.Write(data)

	return nil
}

func (p *Processor) crop() {
	nwf := float64(p.instr.Width)
	nhf := float64(p.instr.Height)

	wf := float64(p.image.Width())
	hf := float64(p.image.Height())

	if nwf != wf || nhf != hf {
		scale := math.Max(nwf/wf, nhf/hf)
		p.image.Resize(uint(scale*wf+0.5), uint(scale*hf+0.5))
	}

	// Center gravity
	nx := int(p.image.Width()/2) - int(nwf/2)
	ny := int(p.image.Height()/2) - int(nhf/2)

	p.image.Crop(nx, ny, uint(p.instr.Width), uint(p.instr.Height))
}

func (p *Processor) shrinkWidth() {
	scale := float64(p.instr.Width) / float64(p.image.Width())

	if scale > 1 {
		return
	}

	nw := uint(p.instr.Width)
	nh := uint(scale * float64(p.image.Height()))

	p.image.Resize(nw, nh)
}
