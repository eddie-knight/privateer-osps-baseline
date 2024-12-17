package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_05() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "QA_05"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Any additional subproject code repositories produced by the project and compiled into a release MUST enforce security requirements as applicable to the status and intent of the respective codebase.",
		ControlID:   "OSPS-QA-05",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_05_T01)

	return
}

func QA_05_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_05
	return
}