package syntheticsv

import(
	"bufio"
	"regexp"
)
/*List of Global Variables*/
var delimiter string
var directory string
var operationDelim string
var firstFieldName string
var csvFileStore *bufio.Scanner
var fields map[string]int
var currentLine int = 0
var totalLine int = 0
var regExpLatitude,_ = regexp.Compile(`^(\+|-)?(?:90(?:(?:\.0{1,12})?)|(?:[0-9]|[1-8][0-9])(?:(?:\.[0-9]{1,12})?))$`)
var regExpLongitude,_ = regexp.Compile(`^(\+|-)?(?:180(?:(?:\.0{1,12})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\.[0-9]{1,12})?))$`)
var listOfReplies = map[int]string{	1: "Error: Please provide the location of your csv file.",
									2: "Error: Please ensure your stand-alone field name is not containing in another.",
									3: "Error: Your pair of Latitude and Longitude is incorrect!"}