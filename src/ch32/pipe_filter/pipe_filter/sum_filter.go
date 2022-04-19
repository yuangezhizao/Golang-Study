package pipefilter

import "errors"

var SumFilterWrongFormatError = errors.New("SumFilterWrongFormatError")

type Sumfilter struct {
}

func NewSumFilter() *Sumfilter {
	return &Sumfilter{}
}

func (sf *Sumfilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
