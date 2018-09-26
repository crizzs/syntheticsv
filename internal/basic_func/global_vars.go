package syntheticsv

import(
	"bufio"
)

var delimiter string
var fields map[string]int
var directory string
var csvFileStore *bufio.Scanner
