package syntheticsv

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeValues(t *testing.T){
	//Reset the settings to default
	SetAllBasicSettings()
	mergedInfo := MergeValues([]string{"iata_code","country"},"-")
	//Test the return value
	assert.Equal(t,mergedInfo, "AAAFrench Polynesia", "First and Last value should merge.")
}



//Function to evoke all basic settings for Unit Test purpose
func SetAllBasicSettings() {
	SetGlobalVars(",","./sample_csv/airports.csv")
	csvFile := ReadFile(directory)
	csvFileStore = csvFile
	MoveToNextLine()
	//Sets header information
	StoreFieldsHeader(csvFileStore.Text())
}
