package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a specific birthday",
	Long:  `Search the birthday of a specific person in the database.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return bd.getNames(toComplete), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Date of Birth", "Days remaining"})
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, people := range bd.Birthdays {
			if people.Name != args[0] {
				continue
			}
			days, err := getRemainingDays(people.Dob)
			if err != nil {
				logrus.Fatalf("cannot get remaining days of %s: %s", people.Name, err)
			}

			data := []string{people.Name, people.Dob, days}
			table.Append(data)
		}

		if table.NumLines() > 0 {
			table.Render()
		} else {
			fmt.Println("No birthdays coming soon")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	home, err := getHomeDir()
	if err != nil {
		logrus.Fatalf("cannot find user home directory: %s", err)
	}

	file := home + "/.bd/dates.json"
	bd.readBirthdays(file)
}
