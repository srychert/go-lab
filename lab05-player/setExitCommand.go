package main

import "regexp"

func setExitCommand(exitMsg string, exitCommand *string) {
	// regex for findig stop command
	rE, _ := regexp.Compile("'([^']*)'")
	matchE := rE.FindString(exitMsg)

	if matchE != "" {
		*exitCommand = matchE[1 : len(matchE)-1]
	}

}
