package root

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/tokisuno/goltime/cmd/goltime/endtime"
)


var hours, minutes, _ = time.Now().Clock()

func TimeLeft(hours int, minutes int) int {
    minsLeft := 1440 - ((hours * 60) + minutes)
    return minsLeft
}

func RootCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "goltime",
        Short: "Minutes until midnight",
        Long: "Displays minutes until midnight",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Fprintf(cmd.OutOrStdout(), "%v", TimeLeft(hours, minutes))
        },
    }

    // For when I add different commands
    cmd.AddCommand(endtime.EndtimeCommand())
    return cmd
}
