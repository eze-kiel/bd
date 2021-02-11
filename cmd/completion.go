package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(bd completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ bd completion bash > /etc/bash_completion.d/bd
  # macOS:
  $ bd completion bash > /usr/local/etc/bash_completion.d/bd

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ bd completion zsh > "${fpath[1]}/_bd"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ bd completion fish | source

  # To load completions for each session, execute once:
  $ bd completion fish > ~/.config/fish/completions/bd.fish

PowerShell:

  PS> bd completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> bd completion powershell > bd.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
