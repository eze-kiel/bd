package cmd

import (
	"os"
	"os/user"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	path string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize bd components",
	Long: `Create $HOME/.bd directory and $HOME/.bd/dates.json file containing all
the birthday dates.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := getHomeDir()
		if err != nil {
			logrus.Fatalf("cannot find user home directory: %s", err)
		}

		path := home + "/.bd"
		logrus.Infof("creating %s directory with perms 0755", path)
		err = os.Mkdir(path, 0755)
		if err != nil {
			logrus.Fatalf("cannot create %s: %s", path, err)
		}

		logrus.Infof("%s successfully created", path)
		logrus.Infof("creating %s/dates.json with perms 0666", path)

		_, err = os.Create(path + "/dates.json")
		if err != nil {
			logrus.Fatalf("error creating %s/dates.json: %s", path, err)
		}

		logrus.Infof("successfully created %s/dates.json", path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func getHomeDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}
