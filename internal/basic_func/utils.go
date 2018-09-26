package syntheticsv

import(
	"os"
	"bufio"
	"strings"
)
/************
Set Global Variables
************/
func SetGlobalVars(delim string,dir string){
	delimiter = delim
	directory = dir
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

/*************
Safe exit from error(s)
*************/
func check(e error) {
    if e != nil {
        os.Exit(0)
    }
}