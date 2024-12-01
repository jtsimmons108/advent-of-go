package days

import (
	"fmt"
	"sort"
	"strings"

	"simmons.com/advent-of-go/utils"
)

const (
	width  = 25
	height = 6
)

type day8 struct {
	layers [][][]byte
}

func Day8() utils.Day {
	d := day8{
		layers: [][][]byte{},
	}
	input := utils.DayInput(2019, 8)
	numLayers := len(input) / width / height

	for i := range numLayers {
		layer := [][]byte{}
		for j := range height {
			row := []byte{}
			for k := range width {
				row = append(row, input[i*width*height+j*width+k])
			}
			layer = append(layer, row)
		}
		d.layers = append(d.layers, layer)
	}

	return d
}

func (d day8) SolvePart1() string {
	counts := []map[byte]int{}
	for _, layer := range d.layers {
		counts = append(counts, CountBytesInLayer(layer))
	}

	sort.Slice(counts, func(i int, j int) bool {
		return counts[i]['0'] < counts[j]['0']
	})
	return fmt.Sprintf("%d", counts[0]['1']*counts[0]['2'])
}

func (d day8) SolvePart2() string {
	img := [][]byte{}
	for r := 0; r < height; r++ {
		row := []byte{}
		for c := 0; c < width; c++ {
			row = append(row, d.calcPixelVal(r, c))
		}
		img = append(img, row)
	}
	res := strings.Builder{}
	res.WriteString("\n")

	pixels := []rune{' ', '#'}
	for r := range height {
		for c := range width {
			res.WriteRune(pixels[img[r][c]-'0'])
		}
		res.WriteString("\n")
	}
	return res.String()
}

func CountBytesInLayer(layer [][]byte) map[byte]int {
	res := map[byte]int{}

	for _, row := range layer {
		for _, b := range row {
			res[b]++
		}
	}
	return res
}

func (d day8) calcPixelVal(r int, c int) byte {
	for _, layer := range d.layers {
		if layer[r][c] != '2' {
			return layer[r][c]
		}
	}
	panic(`Unable to find pixel value`)
}
