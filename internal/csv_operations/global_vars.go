package syntheticsv

import(
	"bufio"
)

var delimiter string
var fields map[string]int
var directory string
var operationDelim string
var firstFieldName string
var csvFileStore *bufio.Scanner
var listOfReplies = map[int]string{1: "Error: Please provide the location of your csv file."}