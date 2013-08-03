package instructions

import (
	"errors"
	"net/url"
	"strconv"
)

type RequestInstructions struct {
	Origin string
	Action string
	Format string
	Width  int
	Height int
}

func ParseInstructions(p url.Values) (*RequestInstructions, error) {
	var err error

	instr := &RequestInstructions{}
	instr.Origin, _ = url.QueryUnescape(p.Get("origin"))
	instr.Action = p.Get("action")
	instr.Format = p.Get("format")

	instr.Width, err = strconv.Atoi(p.Get("width"))
	if err != nil {
		return nil, errors.New("width is not an integer")
	}

	instr.Height, err = strconv.Atoi(p.Get("height"))
	if err != nil {
		instr.Height = 0
	}

	err = instr.validate()
	if err != nil {
		return nil, err
	}

	return instr, nil
}

func (i *RequestInstructions) validate() error {
	if i.Action == "crop" && (i.Width <= 0 || i.Height <= 0) {
		return errors.New("both width and height need to be bigger than 0 with crop action")
	}

	if i.Format != "png" && i.Format != "jpg" {
		return errors.New("format needs to be either jpg or png")
	}

	originUrl, err := url.ParseRequestURI(i.Origin)
	if err != nil || originUrl.Host == "" || originUrl.Scheme == ""{
		return errors.New("origin is not a valid url")
	}

	return nil
}
