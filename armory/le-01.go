package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_01() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The version control system MUST require all code contributors to assert that they are legally authorized to commit the associated contributions on every commit.",
		ControlID:   "OSPS-LE-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_01_T01)

	return "LE_01", result
}

// TODO
func LE_01_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_01
	return testResult
}
