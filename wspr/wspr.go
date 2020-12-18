package wspr

import (
	"fmt"
	"strconv"
)

//Convert a function
func Convert(i int64) {
	fmt.Printf("Decmial %d, Hex %2s, Binary %8s \n", i, strconv.FormatInt(i, 16), strconv.FormatInt(i, 2))
}
