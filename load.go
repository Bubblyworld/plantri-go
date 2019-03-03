package plantri

import (
	"bytes"
	"errors"
	"io/ioutil"
)

const (
	// Headers for each file format output by plantri.
	//   see: http://users.cecs.anu.edu.au/~bdm/plantri/plantri-guide.txt
	headerPlanarCode = ">>planar_code<<"
)

var ErrNoHeader = errors.New("plantri: no header found in given file")
var ErrUnsupportedFileType = errors.New("plantri: given file is not of a supported type")
var ErrFileCorrupted = errors.New("plantri: given file is corrupted")

func Load(path string) ([]Graph, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	header, err := getHeader(data)
	if err != nil {
		return nil, err
	}

	var getGraphFn func([]byte, int) (Graph, int, error)
	switch header {
	case headerPlanarCode:
		getGraphFn = getGraphPlanarCode

	default:
		return nil, ErrUnsupportedFileType
	}

	var res []Graph
	offset := len(header)
	for offset != len(data) {
		graph, n, err := getGraphFn(data, offset)
		if err != nil {
			return nil, err
		}

		res = append(res, graph)
		offset += n
	}

	return res, nil
}

func getHeader(data []byte) (string, error) {
	if len(data) < 4 || !(data[0] == '>' && data[1] == '>') {
		return "", ErrNoHeader
	}

	cc := 0
	res := ">>"
	for i := 2; i < len(data); i++ {
		res += string(data[i])

		if data[i] == '<' {
			cc++
		} else {
			cc = 0
		}

		if cc == 2 {
			return res, nil
		}
	}

	return "", ErrNoHeader
}

// getGraphPlanarCode parses `data` from the given `offset`, assuming the
// data is in the PlanarCode format, and returns the next graph found in the
// data along with the number of bytes parsed to find it.
func getGraphPlanarCode(data []byte, offset int) (Graph, int, error) {
	cnt := 0
	buf := bytes.NewBuffer(data)
	buf.Next(offset)

	n, err := buf.ReadByte()
	if err != nil {
		return nil, 0, ErrFileCorrupted
	}
	cnt++

	res := newAdjMatrix(int(n))
	for i := 0; i < int(n); i++ {
		for {
			v, err := buf.ReadByte()
			if err != nil {
				return nil, 0, ErrFileCorrupted
			}
			cnt++

			if int(v) == 0 {
				break
			}

			res.addEdge(i, int(v)-1) // planarcode is 1-indexed
		}
	}

	return res, cnt, nil
}
