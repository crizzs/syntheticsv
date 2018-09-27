package syntheticsv

import(
	"testing"
	"github.com/stretchr/testify/assert"
)
/****
Test fields merging capability 
****/
func TestMergeValues(t *testing.T){
	//Reset the settings to default
	SetAllBasicSettings()
	//(Scenario 1: Different Ops Delim)
	SetGlobalVars(",","./sample_csv/airports.csv","")
	mergedInfo := MergeValues([]string{"iata_code","country"})
	assert.Equal(t,mergedInfo, "AAAFrench Polynesia", "First and Last value should merge.")
	//(Scenario 2: Different Ops Delim)
	SetGlobalVars(",","./sample_csv/airports.csv","|")
	mergedInfo = MergeValues([]string{"iata_code","country"})
	assert.Equal(t,mergedInfo, "AAA|French Polynesia", "First and Last value should merge.")
	//(Scenario 3: Identical Ops Delim and Delimiter)
	SetGlobalVars(",","./sample_csv/airports.csv",",")
	mergedInfo = MergeValues([]string{"iata_code","country"})
	assert.Equal(t,mergedInfo, "AAAFrench Polynesia", "Comma is identical, and be replaced with empty string.")
}
/****
Test basic arithmetic capability between fields
****/
func TestBasicArithmetic(t *testing.T){

}

func TestRetrieveVals(t *testing.T){
	
	anotherLineArr := SplitLine(csvFileStore.Text())
	//Retrieve Value based on field
	val := RetrieveValue(anotherLineArr,"country")

	assert.Equal(t,val, "French Polynesia", "Must match the position of header field.")

	//Test when field name doesn't exist
	noVal := RetrieveValue(anotherLineArr,"fake_country")
	assert.Equal(t,noVal, "", "The return needs to be empty string.")
}

//Function to evoke all basic settings for Unit Test purpose
func SetAllBasicSettings() {
	SetGlobalVars(",","./sample_csv/airports.csv","")
	csvFile := ReadFile(directory)
	csvFileStore = csvFile
	MoveToNextLine()
	//Sets header information
	StoreFieldsHeader(csvFileStore.Text())
	MoveToNextLine()
}
