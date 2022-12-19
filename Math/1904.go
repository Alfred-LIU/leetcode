package Math

import (
	"strconv"
)

/*
You are participating in an online chess tournament. There is a chess round that starts every 15 minutes. The first round of the day starts at 00:00, and after every 15 minutes, a new round starts.

For example, the second round starts at 00:15, the fourth round starts at 00:45, and the seventh round starts at 01:30.
You are given two strings loginTime and logoutTime where:

loginTime is the time you will login to the game, and
logoutTime is the time you will logout from the game.
If logoutTime is earlier than loginTime, this means you have played from loginTime to midnight and from midnight to logoutTime.

Return the number of full chess rounds you have played in the tournament.

Note: All the given times follow the 24-hour clock. That means the first round of the day starts at 00:00 and the last round of the day starts at 23:45.
*/
func numberOfRounds(startTime string, finishTime string) int {
	startH, _ := strconv.Atoi(startTime[:2])
	startM, _ := strconv.Atoi(startTime[3:])
	finishH, _ := strconv.Atoi(finishTime[:2])
	finishM, _ := strconv.Atoi(finishTime[3:])

	start := startH*60 + startM
	finish := finishH*60 + finishM
	if finish < start {
		finish += 1440
	}

	if remainder := finish % 15; remainder != 0 {
		finish -= remainder
	}

	return (finish - start) / 15
}
