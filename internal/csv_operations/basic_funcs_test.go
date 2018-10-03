package syntheticsv

import(
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)
/****
Test fields merging capability 
****/
func TestMergeValues(t *testing.T){
	//Reset the settings to default
	SetAllBasicSettings(",","./sample_csv/airports.csv","")
	//(Scenario 1: Different Ops Delim)
	SetAllBasicSettings(",","./sample_csv/airports.csv","")
	mergedInfo := MergeValues([]string{"iata_code","country"})
	assert.Equal(t,mergedInfo, "AAAFrench Polynesia", "First and Last value should merge.")
	//(Scenario 2: Different Ops Delim)
	SetAllBasicSettings(",","./sample_csv/airports.csv","|")
	mergedInfo = MergeValues([]string{"iata_code","country"})
	assert.Equal(t,mergedInfo, "AAA|French Polynesia", "First and Last value should merge.")
	//(Scenario 3: Identical Ops Delim and Delimiter)
	SetAllBasicSettings(",","./sample_csv/airports.csv",",")
	mergedInfo = MergeValues([]string{"iata_code","country"})
	assert.Equal(t,mergedInfo, "AAAFrench Polynesia", "Comma is identical, and be replaced with empty string.")
}

func TestRetrieveVals(t *testing.T){
	
	anotherLineArr := SplitLine(csvFileStore.Text())
	//Retrieve Value based on field
	val := RetrieveStrValue(anotherLineArr,"country")

	assert.Equal(t,val, "French Polynesia", "Must match the position of header field.")

	//Test when field name doesn't exist
	noVal := RetrieveStrValue(anotherLineArr,"fake_country")
	assert.Equal(t,noVal, "", "The return needs to be empty string.")
}
/****
Test basic arithmetic capability between fields
****/
func TestBasicArithmetic(t *testing.T){
	//Loads supermarket data set for arithmetic functions
	SetAllBasicSettings(",","./sample_csv/supermarket.csv","")
	secondLineArr := SplitLine(csvFileStore.Text())

	val := BasicCalculation("selling_price*(1-(goods_and_services_tax+alcohol_tax))-cost_price",secondLineArr)
	//Test for bodmas rule
	assert.Equal(t, val, 1.2355, "Profit must match with result")
	//Test data which contains empty string
	MoveToNextLine()
	thirdLineArr := SplitLine(csvFileStore.Text())

	invalidVal := BasicCalculation("selling_price*(1-(goods_and_services_tax+alcohol_tax))-cost_price",thirdLineArr)
	assert.Equal(t, invalidVal, 0.0, "An invalid equation will yield default 0 as result.")
}
/*****
Test whether row containing particular attribute
*****/
func TestContainsAttribute(t *testing.T){
	thirdLineArr := SplitLine(csvFileStore.Text())

	containInFieldStatus := ContainInField("item_name","beer",thirdLineArr)
	containInRowStatus := ContainInRow("milk",thirdLineArr)
	assert.Equal(t, containInFieldStatus, true, "Check whether beer is containing in this field.")
	assert.Equal(t, containInRowStatus, false, "Check whether milk is containing in this row.")
}
/******
Test the generating of geojson for a single Point based on csv data
*******/
func TestCreatePointGeojsonFunc(t *testing.T){
	thirdLineArr := SplitLine(csvFileStore.Text())
	//Retrieve Location (Latitude/Longitude) and create geojson
	location := RetrieveStrValue(thirdLineArr,"location")
	splitLocation := SplitLineWithDelim(location,"|")
	//Create geojson for a single point(Based on Latitude and Longitude)
	pointGeoJson := CreatePointGeoJson(splitLocation[0],splitLocation[1])
	assert.Equal(t, pointGeoJson, "", "Point Geojson should be generated.")

	//Test different scenarios
	incorrectLatLngOne := CreatePointGeoJson("-91","180")
	incorrectLatLngTwo := CreatePointGeoJson("91","180")
	incorrectLatLngThree := CreatePointGeoJson("90","181")
	incorrectLatLngFour := CreatePointGeoJson("90","-181")
	incorrectLatLngFive := CreatePointGeoJson("90","")
	incorrectLatLngSix := CreatePointGeoJson("abc","-180")

	assert.Equal(t, incorrectLatLngOne, "", "Should catch abnormalities.")
	assert.Equal(t, incorrectLatLngTwo, "", "Should catch abnormalities.")
	assert.Equal(t, incorrectLatLngThree, "", "Should catch abnormalities.")
	assert.Equal(t, incorrectLatLngFour, "", "Should catch abnormalities.")
	assert.Equal(t, incorrectLatLngFive, "", "Should catch abnormalities.")
	assert.Equal(t, incorrectLatLngSix, "", "Should catch abnormalities.")
}
//Function to evoke all basic settings for Unit Test purpose
func SetAllBasicSettings(delim string,dir string,opsDelim string) {
	SetGlobalVars(delim,dir,opsDelim)
	csvFile := ReadFile(directory)
	csvFileStore = csvFile
	MoveToNextLine()
	//Sets header information
	StoreFieldsHeader(csvFileStore.Text())
	MoveToNextLine()
}
