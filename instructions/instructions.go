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
		return nil, errors.New("height is not an integer")
	}

	return instr, nil
}
