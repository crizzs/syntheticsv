package syntheticsv

import(
	"strconv"
	"strings"
	"github.com/soniah/evaler"
)
/*******
Merge values on a single row with delimiter
********/
func MergeValues(rowValues []string) string{
	//Check if the delimiter for MergedValues is identical to csv delimiter
	if operationDelim == delimiter {
		operationDelim = ""
	}

	//Split the Line into Array
	anotherLineArr := SplitLine(csvFileStore.Text())

	i := 0
	concatStr := ""

	for i < len(rowValues) {
		concatStr += RetrieveStrValue(anotherLineArr,rowValues[i])
		if i != (len(rowValues)-1){
			concatStr += operationDelim
		}

		i++
	}
	
	return concatStr
}
/***********
Carry out basic arithmetic function
for single row based on BODMAS rule
************/
func BasicCalculation(formula string,rowValues []string) float64{
	//This check ensures field name replacer works
	CheckFieldNameValidity()

	for eachField, _ := range fields {
		fieldFloatVal , fieldValidity := RetrieveFloatValue(rowValues,eachField)

		floatStr := strconv.FormatFloat(fieldFloatVal,'g' ,-1, 64)
		//Subsitute an empty string to destabilise the formula
		if fieldValidity == false {
			floatStr = ""
		}
		
		formula = strings.Replace(formula,eachField,floatStr,-1)
	}
	//Eval formula based on bodmas rule
	result, err := evaler.Eval(formula)
	if err != nil {
		return 0.0
	}

	evalFloatResult := evaler.BigratToFloat(result)
	
	return evalFloatResult
}
/********
Check whether string is containing in this field
**********/
func ContainInField(fieldName string,containVal string,rowValues []string) bool{
	val := RetrieveStrValue(rowValues,fieldName)

	if strings.Contains(val,containVal) {
		return true
	}
	return false
}
/********
Check whether string is containing in this row
**********/
func ContainInRow(containVal string,rowValues []string) bool{
	for eachField,_ := range fields{
		status := ContainInField(eachField,containVal,rowValues)

		if status == true {
			return true
		}
	}
	return false
}
/*********
Creates a Point GeoJson
**********/
func CreatePointGeoJson(lat string,lon string) string{
	//Check if Latitude and Longitude are both valid
	latValid := regExpLatitude.MatchString(lat)
	lonValid := regExpLongitude.MatchString(lon)

	if latValid == false || lonValid == false {
		ErrorMsgGeneratorWithoutExit(3)
	}

 	return ""
}
/**********
Check for Entire Field Name
existing in another Field Name
***********/
func CheckFieldNameValidity(){
	//Field List is generally small, an O(n^2) comparator function is being implemented
	for eachField, _ := range fields {
		for checkAgainstField, _ := range fields {
			//Using a map ensures the second condition will not happen
			if strings.Contains(eachField,checkAgainstField) && eachField != checkAgainstField {
				//Once found, it will give a message and os exit 
				ErrorMsgGenerator(2)
			}
		}
	}
}
/*******
Gets string Value from String Array
*******/
func RetrieveStrValue(allValues []string, fieldName string) (string){
	//Gets field position and return value
	fieldPos := fields[fieldName]

	if fieldPos == 0 && firstFieldName != fieldName{
		return ""
	}

	fieldVal := allValues[fieldPos]

	return fieldVal
}
/*******
Gets "float" Value from String Array
*******/
func RetrieveFloatValue(allValues []string, fieldName string) (float64,bool){
	//Gets field position and return value
	fieldPos := fields[fieldName]

	if fieldPos == 0 && firstFieldName != fieldName{
		return 0.0,false
	}

	floatVal, err := strconv.ParseFloat(allValues[fieldPos],64)
	if err != nil {
		return 0.0,false
	}

	return floatVal,true
}