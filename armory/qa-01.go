package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_01() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project’s source code MUST be publicly readable and have a static URL.",
		ControlID:   "OSPS-QA-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_01_T01)

	return "QA_01", result
}

// TODO
func QA_01_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_01
	return moveResult
}
