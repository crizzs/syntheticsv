package syntheticsv

import(
	"bufio"
)
/*List of Global Variables*/
var delimiter string
var directory string
var operationDelim string
var firstFieldName string
var csvFileStore *bufio.Scanner
var fields map[string]int
var listOfReplies = map[int]string{	1: "Error: Please provide the location of your csv file.",
									2: "Error: Please ensure your stand-alone field name is not containing in another."}