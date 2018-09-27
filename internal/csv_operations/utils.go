package syntheticsv

import(
	"os"
	"bufio"
	"strings"
	"fmt"
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
		errorMsgGenerator(1)
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

    return scanner
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

/**************
Move to next line
**************/
func MoveToNextLine(){
	csvFileStore.Scan()
}

func errorMsgGenerator(msgNum int) {
	fmt.Println(listOfReplies[msgNum])
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