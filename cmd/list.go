package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dateLayout = "02-01-2006"

type birthdays struct {
	Birthdays []struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	} `json:"birthdays"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the saved birthdays",
	Long:  `List all the people and birthdays that are registered in the base.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := getHomeDir()
		if err != nil {
			logrus.Fatalf("cannot find user home directory: %s", err)
		}

		file := home + "/.bd/dates.json"
		bd.readBirthdays(file)

		sortByDays, err := cmd.Flags().GetBool("days")
		if err != nil {
			logrus.Fatalf("cannot parse days flag: %s", err)
		}
		if sortByDays {
			bd.sortDatabaseByDays()
		} else {
			bd.sortDatabaseByDob()
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Date of Birth", "Days remaining"})
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, people := range bd.Birthdays {
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
			fmt.Println("No birthdays registered in list")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("days", "d", false, "Sort by remaining days")
}

func (bd *birthdays) readBirthdays(f string) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &bd)
}
