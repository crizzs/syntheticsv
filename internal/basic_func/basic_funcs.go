package syntheticsv

import(
	"errors"
)
/*******
Merge values on a single row with delimiter
********/
func MergeValues(fields []string, delim string) string{
	MoveToNextLine()
	//Split the Line into Array
	anotherLineArr := SplitLine(csvFileStore.Text())

	i := 0

	for i < len(anotherLineArr) {
		
	}

	return ""
}

/*******
Gets Value from String Array
*******/
func RetrieveValue(allValues []string, fieldName string) (string,error){
	return "" ,_
}