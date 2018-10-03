package syntheticsv

import(
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

/************
Set Global Variables
************/
func SetGlobalVars(delim string,dir string,opsDelim string){
	//Default settings
	if delim == "" {
		delim = ","
	}
	if opsDelim == "" {
		operationDelim = "|"
	}
	//If identical, operation delimiter will be empty string
	if delim == opsDelim {
		operationDelim = ""
	}
	//Throws error when no directory is found
	if dir == "" {
		ErrorMsgGenerator(1)
	}

	delimiter = delim
	directory = dir
	operationDelim = opsDelim
}

/************
Read CSV file line by line
*************/
func ReadFile(dir string) *bufio.Scanner {
	file, err := os.Open(dir)
    check(err)

    scanner := bufio.NewScanner(file)
    //Reset to default setting
    currentLine = 0
    totalLine = 0
    return scanner
}
/*************
Open File
*************/
func openFile(dir string) *os.File {
    var file, err = os.OpenFile(dir, os.O_APPEND|os.O_WRONLY, 0644)
    check(err)
    return file
}
/************
Write to File
*************/
func writeToFile(sqlStr string,file *os.File){
    if _, err := file.WriteString(sqlStr); err != nil {
        check(err)
    }
    //Synchronise file
    err := file.Sync()
    check(err)
}
/*************
Store Fields to a map
**************/
func StoreFieldsHeader(firstLine string) map[string]int {
	allFields := strings.Split(firstLine,delimiter)
	fields = make(map[string]int)
	firstFieldName = allFields[0]

	i := 0

	for i < len(allFields){
		fields[allFields[i]] = i
		i++
	}

	return fields
}

/*********
Split Line into string array (Non-Header)
*********/
func SplitLine(eachLine string) []string{
	allValues := strings.Split(eachLine,delimiter)
	return allValues
}
/*********
Split Line With Specific delimiter (Non-Header)
*********/
func SplitLineWithDelim(eachLine string,delim string) []string{
	allValues := strings.Split(eachLine,delim)
	return allValues
}
/**************
Move to next line
**************/
func MoveToNextLine(){
	csvFileStore.Scan()
	currentLine++
}
/*****
Only Show Error Message
******/
func ErrorMsgGeneratorWithoutExit(msgNum int) {
	fmt.Println(listOfReplies[msgNum] + "(Encountered on Line "+ strconv.Itoa(currentLine) +")")
}
/*****
Exit with Error Message
******/
func ErrorMsgGenerator(msgNum int) {
	fmt.Println(listOfReplies[msgNum] + "(Encountered on Line "+ strconv.Itoa(currentLine) +")")
	os.Exit(0)
}
/*************
Safe exit from error(s)
*************/
func check(e error) {
    if e != nil {
        os.Exit(0)
    }
}