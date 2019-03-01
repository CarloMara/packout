package model

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"sort"
	"strconv"

	//"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"path/filepath"
)

type Position struct {
	X string
	Y string
}

type Point struct {
	Pos Position
	Col string
}

type FileData struct {
	ID        xid.ID
	Name      string
	Path      string
	Frequency map[int]int
	C1        colorful.Color
	C2        colorful.Color
	Data      []Point
}

func InitFile(path string) *FileData {
	c1, _ := colorful.Hex("#fdffcc")
	c2, _ := colorful.Hex("#242a42")
	return &FileData{ID: xid.New(), Name: filepath.Base(path), Path: filepath.Clean(path), C1: c1, C2: c2}

}

func (file *FileData) SampleTuple() {

	frequency := make([]int, 65536)

	data, err := ioutil.ReadFile(file.Path)
	fmt.Println(err)

	file_lenght := len(data)

	// if file has a odd number of bytes remove one so it's even. Shity hack to make later code not trow index out of bounds
	if file_lenght%2 == 1 {
		file_lenght--
	}

	var index uint16

	for i := 0; i < file_lenght; i += 2 {
		index = uint16(data[i]) << 8
		index += uint16(data[i+1])
		if index != 0 {
			frequency[int(index)]++
		}
	}

	min, max := frequency[0], frequency[0]
	for _, v := range frequency {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	//var min = 100
	//var max = 10000

	//file.Data = make([]Point, )

	//min = 10
	for pos, freq := range frequency {
		if freq > 10 {
			x := pos >> 8
			y := pos & 255
			//file.C1.BlendHsv(file.C2, float64(freq-min*255)/float64(max-min)).Hex()
			file.Data = append(file.Data, Point{Pos: Position{X: strconv.Itoa(int(x)), Y: strconv.Itoa(y)}, Col: file.C1.BlendLab(file.C2, float64(freq-min)/float64(max-min)).Hex()})
		}
	}

}

func (file *FileData) SampleByte() {

	file.Frequency = make(map[int]int)
	data, _ := ioutil.ReadFile(file.Path)

	file_lenght := len(data)

	for i := 0; i < file_lenght; i++ {
		file.Frequency[int(data[i])]++
	}
}

func minIntSlice(v []int) int {
	sort.Ints(v)
	return v[0]
}

func maxIntSlice(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}
