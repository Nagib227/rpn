package rpn

import (
	"fmt"
	// "strconv"
	// "strings"
)

type World struct {
	Height int 
    Width int 
    Cells [][]bool 
}

 func NewWorld(height, width int) (*World, error){
 	if height <= 0 || width <= 0 {
 		return nil, fmt.Errorf("err")
 	}
 	return &World{Height: height, Width: width}, nil
 }
