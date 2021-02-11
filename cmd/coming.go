package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// comingCmd represents the coming command
var comingCmd = &cobra.Command{
	Use:   "coming",
	Short: "Display coming birthdays",
	Long:  `Display all the birthday that will occur in less than a month (31 days).`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := getHomeDir()
		if err != nil {
			logrus.Fatalf("cannot find user home directory: %s", err)
		}

		file := home + "/.bd/dates.json"
		bd.readBirthdays(file)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Date of Birth", "Days remaining"})
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, people := range bd.Birthdays {
			days, err := getRemainingDays(people.Dob)
			if err != nil {
				logrus.Fatalf("cannot get remaining days of %s: %s", people.Name, err)
			}

			daysInt, err := strconv.Atoi(days)
			if err != nil {
				logrus.Fatalf("cannot convert %s to int: %s", days, err)
			}

			if daysInt > 31 {
				continue
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
	rootCmd.AddCommand(comingCmd)
}
