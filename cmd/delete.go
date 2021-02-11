package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: `Delete an entry from the birthday database`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return bd.getNames(toComplete), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		for idx, p := range bd.Birthdays {
			if p.Name != args[0] {
				continue
			}
			bd.remove(idx)
			name = p.Name
		}

		home, err := getHomeDir()
		if err != nil {
			logrus.Fatalf("cannot find user home directory: %s", err)
		}

		file := home + "/.bd/dates.json"
		bd.updateDatabase(file)

		logrus.Infof("successfully removed %s from database", name)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	home, err := getHomeDir()
	if err != nil {
		logrus.Fatalf("cannot find user home directory: %s", err)
	}

	file := home + "/.bd/dates.json"
	bd.readBirthdays(file)
}

func (bd *birthdays) remove(i int) {
	bd.Birthdays[i] = bd.Birthdays[len(bd.Birthdays)-1]
	bd.Birthdays = bd.Birthdays[:len(bd.Birthdays)-1]
}
