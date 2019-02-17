
package sample

import (
	"fmt"
	"io/ioutil"
)



func sample_2byte(file string) []int{
	ydata := make([]int, 65536)
	data, _ := ioutil.ReadFile(file)

	file_lenght := len(data)
	var is_even bool

	if file_lenght % 2 == 0{
		is_even = true
	} else{
		is_even = false
	}

	var index uint16

	for i := 0; i < file_lenght; i+=2 {
		// word := data[i:i+2]
		index = uint16(data[i]) << 8
		index += uint16(data[i+1])
		// fmt.Printf("%#x\n", index)
		// used to debug
		// if i == 10000{
		// 	break
		// }
		ydata[int(index)]++
	}
	fmt.Print(is_even)

	// var temp_byte int
	// var is_qword bool

	// is_qword = false
	// for _,byte := range data {

	// 	if is_qword{
	// 		ydata[int(temp_byte+byte)]++
	// 	}

	// 	if !is_qword {
	// 		temp_byte = byte >> 8
	// 		is_qword = true
	// 	}
	// }

	fmt.Print(ydata)
	return ydata
}