package syntheticsv

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

//Setting Global Variables (Test)
func TestSetGlobalVars(t *testing.T){
	SetGlobalVars(",","./sample_csv/airports.csv")
	assert.Equal(t,delimiter, ",", "This sets the delimiter of csv file.")
	assert.Equal(t,directory, "./sample_csv/airports.csv", "This sets the location of csv file.")
}

//Test the read file function
func TestReadFile(t *testing.T){
	csvFile := ReadFile(directory)
	csvFileStore = csvFile
	MoveToNextLine()

	//Check if the file is being retrieved properly
	assert.Equal(t,csvFileStore.Text(), "iata_code,city,airport,country", "This is the header.")
}

//Test storage of fields in header
func TestFieldHeaderStorage(t *testing.T){
	allFields := StoreFieldsHeader(csvFileStore.Text())
	//Check if header fields have been digested
	assert.Equal(t,allFields["iata_code"], 0, "The field must correspond with its position")
}

//Test Split String
func TestSplitString(t *testing.T){
	MoveToNextLine()
	secondLine := SplitLine(csvFileStore.Text())
	//Check the string array is being created
	assert.Equal(t,secondLine, []string([]string{"AAA", "Anaa", "Anaa", "French Polynesia"}), "String Array should match")
}

