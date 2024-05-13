package endtime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func currentTime() (int, int) {
    var hours, minutes, _ = time.Now().Clock()

    return hours, minutes
}

func isRealTime(userHours int, userMinutes int) bool {
    var hours, minutes, _ = time.Now().Clock()

    if userHours >= 24 {
        return false
    }

    if userMinutes >= 60 {
        return false
    }

    if userHours <= hours {
        if userMinutes < minutes {
            return false
        }
    }

    if isValidTime(userHours, userMinutes) == false {
        return false
    }

    return true
}

func isValidTime(userHours int, userMinutes int) bool {
    currentHours, currentMinutes := currentTime()

    if userHours < currentHours {
        return false
    }

    if userHours == currentHours && userMinutes < currentMinutes {
        return false
    }

    return true
}

func timeLeft(arg string) int {
    var mid int = len(arg)/2

    hours_str := arg[:mid]
    minutes_str := arg[mid:]

    var hours, _ = strconv.Atoi(hours_str)
    var minutes, _ = strconv.Atoi(minutes_str)

    if isRealTime(hours, minutes) == false {
        panic(1)
    }
    currentHours, currentMinutes := currentTime()
    newMidnight := (hours * 60) + minutes

    minsLeft := newMidnight - ((currentHours * 60) + currentMinutes)

    return minsLeft
}

func EndtimeCommand() *cobra.Command {
    return &cobra.Command{
        Use: "endt",
        Short: "Displays minutes until end-time",
        Long: "Displays minutes remaining until user-defined end-time (HHMM)",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Fprintf(cmd.OutOrStdout(), "%v", timeLeft(args[0]))
        },
    }
}
