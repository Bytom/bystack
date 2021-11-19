package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/bytom/bystack/util"
)

// bystackd usage template
var usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletDisable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

  available with wallet enable:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletEnable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// commandError is an error used to signal different error situations in command handling.
type commandError struct {
	s         string
	userError bool
}

func (c commandError) Error() string {
	return c.s
}

func (c commandError) isUserError() bool {
	return c.userError
}

func newUserError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: true}
}

func newSystemError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: false}
}

func newSystemErrorF(format string, a ...interface{}) commandError {
	return commandError{s: fmt.Sprintf(format, a...), userError: false}
}

// Catch some of the obvious user errors from Cobra.
// We don't want to show the usage message for every error.
// The below may be to generic. Time will show.
var userErrorRegexp = regexp.MustCompile("argument|flag|shorthand")

func isUserError(err error) bool {
	if cErr, ok := err.(commandError); ok && cErr.isUserError() {
		return true
	}

	return userErrorRegexp.MatchString(err.Error())
}

// BystackcliCmd is Bystackcli's root command.
// Every other command attached to BystackcliCmd is a child command to it.
var BystackcliCmd = &cobra.Command{
	Use:   "bystackd",
	Short: "Bystackcli is a commond line client for bytom core (a.k.a. bystackd)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.SetUsageTemplate(usageTemplate)
			cmd.Usage()
		}
	},
}

// Execute adds all child commands to the root command BystackcliCmd and sets flags appropriately.
func Execute() {

	AddCommands()
	AddTemplateFunc()

	if _, err := BystackcliCmd.ExecuteC(); err != nil {
		os.Exit(util.ErrLocalExe)
	}
}

// AddCommands adds child commands to the root command BystackcliCmd.
func AddCommands() {
	BystackcliCmd.AddCommand(createAccessTokenCmd)
	BystackcliCmd.AddCommand(listAccessTokenCmd)
	BystackcliCmd.AddCommand(deleteAccessTokenCmd)
	BystackcliCmd.AddCommand(checkAccessTokenCmd)

	BystackcliCmd.AddCommand(createAccountCmd)
	BystackcliCmd.AddCommand(deleteAccountCmd)
	BystackcliCmd.AddCommand(listAccountsCmd)
	BystackcliCmd.AddCommand(updateAccountAliasCmd)
	BystackcliCmd.AddCommand(createAccountReceiverCmd)
	BystackcliCmd.AddCommand(listAddressesCmd)
	BystackcliCmd.AddCommand(validateAddressCmd)
	BystackcliCmd.AddCommand(listPubKeysCmd)

	BystackcliCmd.AddCommand(createAssetCmd)
	BystackcliCmd.AddCommand(getAssetCmd)
	BystackcliCmd.AddCommand(listAssetsCmd)
	BystackcliCmd.AddCommand(updateAssetAliasCmd)

	BystackcliCmd.AddCommand(getTransactionCmd)
	BystackcliCmd.AddCommand(listTransactionsCmd)

	BystackcliCmd.AddCommand(getUnconfirmedTransactionCmd)
	BystackcliCmd.AddCommand(listUnconfirmedTransactionsCmd)
	BystackcliCmd.AddCommand(decodeRawTransactionCmd)

	BystackcliCmd.AddCommand(listUnspentOutputsCmd)
	BystackcliCmd.AddCommand(listBalancesCmd)

	BystackcliCmd.AddCommand(rescanWalletCmd)
	BystackcliCmd.AddCommand(walletInfoCmd)

	BystackcliCmd.AddCommand(buildTransactionCmd)
	BystackcliCmd.AddCommand(signTransactionCmd)
	BystackcliCmd.AddCommand(submitTransactionCmd)
	BystackcliCmd.AddCommand(estimateTransactionGasCmd)

	BystackcliCmd.AddCommand(getBlockCountCmd)
	BystackcliCmd.AddCommand(getBlockHashCmd)
	BystackcliCmd.AddCommand(getBlockCmd)
	BystackcliCmd.AddCommand(getBlockHeaderCmd)

	BystackcliCmd.AddCommand(createKeyCmd)
	BystackcliCmd.AddCommand(deleteKeyCmd)
	BystackcliCmd.AddCommand(listKeysCmd)
	BystackcliCmd.AddCommand(updateKeyAliasCmd)
	BystackcliCmd.AddCommand(resetKeyPwdCmd)
	BystackcliCmd.AddCommand(checkKeyPwdCmd)

	BystackcliCmd.AddCommand(signMsgCmd)
	BystackcliCmd.AddCommand(verifyMsgCmd)
	BystackcliCmd.AddCommand(decodeProgCmd)

	BystackcliCmd.AddCommand(createTransactionFeedCmd)
	BystackcliCmd.AddCommand(listTransactionFeedsCmd)
	BystackcliCmd.AddCommand(deleteTransactionFeedCmd)
	BystackcliCmd.AddCommand(getTransactionFeedCmd)
	BystackcliCmd.AddCommand(updateTransactionFeedCmd)

	BystackcliCmd.AddCommand(netInfoCmd)
	BystackcliCmd.AddCommand(gasRateCmd)

	BystackcliCmd.AddCommand(versionCmd)
}

// AddTemplateFunc adds usage template to the root command BystackcliCmd.
func AddTemplateFunc() {
	walletEnableCmd := []string{
		createAccountCmd.Name(),
		listAccountsCmd.Name(),
		deleteAccountCmd.Name(),
		updateAccountAliasCmd.Name(),
		createAccountReceiverCmd.Name(),
		listAddressesCmd.Name(),
		validateAddressCmd.Name(),
		listPubKeysCmd.Name(),

		createAssetCmd.Name(),
		getAssetCmd.Name(),
		listAssetsCmd.Name(),
		updateAssetAliasCmd.Name(),

		createKeyCmd.Name(),
		deleteKeyCmd.Name(),
		listKeysCmd.Name(),
		resetKeyPwdCmd.Name(),
		checkKeyPwdCmd.Name(),
		signMsgCmd.Name(),

		buildTransactionCmd.Name(),
		signTransactionCmd.Name(),

		getTransactionCmd.Name(),
		listTransactionsCmd.Name(),
		listUnspentOutputsCmd.Name(),
		listBalancesCmd.Name(),

		rescanWalletCmd.Name(),
		walletInfoCmd.Name(),
	}

	cobra.AddTemplateFunc("WalletEnable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return true
			}
		}
		return false
	})

	cobra.AddTemplateFunc("WalletDisable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return false
			}
		}
		return true
	})
}
