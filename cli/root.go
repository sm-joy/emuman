package cli

import (
	//"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/sm-joy/emuman/tui"
	"github.com/spf13/cobra"
)

func initroot() cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "emuman",
		Short: "An android emulator manager.",
		Long:  "emuman is a simple wrapper over the Android SDK an emulator management.",
		Run: func(cmd *cobra.Command, args []string) {
			tui.Start()
		},
	}

	var sdkCmd = &cobra.Command{
		Use:   "sdk",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var listCmd = &cobra.Command{
		Use: "list",
		Short: "",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {
			Cmd := exec.Command("sdkmanager.bat", "--list", "--channel=0")
			Cmd.Stdout = os.Stdout
			Cmd.Stderr = os.Stderr
			Cmd.Stdin = os.Stdin

			err := Cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	sdkCmd.AddCommand(listCmd)

	rootCmd.AddCommand(sdkCmd)

	return *rootCmd
}

func Execute() {
	var rootCmd = initroot()
	if error := rootCmd.Execute(); error != nil {
		os.Exit(1)
	}
}
