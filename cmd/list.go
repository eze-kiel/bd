package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

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
}

func (bd *birthdays) readBirthdays(f string) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &bd)
}

func getRemainingDays(date string) (string, error) {
	dob, err := time.Parse(dateLayout, date)
	if err != nil {
		return "", err
	}
	curr := time.Now().Format(dateLayout)
	t2, err := time.Parse(dateLayout, curr)
	if err != nil {
		return "", err
	}

	var t1 time.Time
	t1 = dob.AddDate(t2.Year()-dob.Year(), 0, 0)
	days := t1.Sub(t2).Hours() / 24
	if days < 0 {
		days = 365 + days
	}

	return fmt.Sprintf("%0.f", days), nil
}
