package syntheticsv

import(
	//"fmt"
	//"strconv"
)
/*******
Merge values on a single row with delimiter
********/
func MergeValues(fields []string) string{
	//Check if the delimiter for MergedValues is identical to csv delimiter
	if operationDelim == delimiter {
		operationDelim = ""
	}

	//Split the Line into Array
	anotherLineArr := SplitLine(csvFileStore.Text())

	i := 0
	concatStr := ""

	for i < len(fields) {
		concatStr += RetrieveValue(anotherLineArr,fields[i])
		if i != (len(fields)-1){
			concatStr += operationDelim
		}

		i++
	}
	
	return concatStr
}

/*******
Gets Value from String Array
*******/
func RetrieveValue(allValues []string, fieldName string) (string){
	//Gets field position and return value
	fieldPos := fields[fieldName]

	if fieldPos == 0 && firstFieldName != fieldName{
		return ""
	}

	fieldVal := allValues[fieldPos]

	return fieldVal
}