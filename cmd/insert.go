package cmd

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type entry struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a birthday date into the base",
	Long:  `Insert a person into the base. A name and a date of birth will be asked.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := getHomeDir()
		if err != nil {
			logrus.Fatalf("cannot find user home directory: %s", err)
		}

		file := home + "/.bd/dates.json"
		bd.readBirthdays(file)

		name, err := askInput("Full Name")
		if err != nil {
			logrus.Fatalf("cannot read name: %s", err)
		}
		var yes string
		if len(bd.getNames(name)) != 0 {
			logrus.Warnf("%s is already in the database", name)
			yes = askChoice("Overwrite")
			if yes != "y" {
				logrus.Infof("insertion cancelled")
				return
			}

			bd.removeFromDatabase(bd.getNameIndex(name))
		}

		dob, err := askInput("Date of Birth (DD-MM-YYYY)")
		if err != nil {
			logrus.Fatalf("cannot read dob: %s", err)
		}

		bd.Birthdays = append(bd.Birthdays, entry{Name: name, Dob: dob})
		bd.updateDatabase(file)
		logrus.Infof("successfully created new entry for %s", name)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}

func askInput(label string) (string, error) {
	var validate func(input string) error
	switch label {
	case "Full Name":
		validate = func(input string) error {
			if len(input) <= 0 {
				return errors.New("invalid length")
			}
			return nil
		}
	case "Date of Birth (DD-MM-YYYY)":
		validate = func(input string) error {
			if !strings.Contains(input, "-") || len(input) != 10 {
				return errors.New("invalid length/format")
			}
			return nil
		}
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func askChoice(label string) string {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return "n"
	}

	return "y"
}
