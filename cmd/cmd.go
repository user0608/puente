package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	version     string
	descripcion string
	commit      string
	branch      string
)
var configExample string

type ConfigPath string

var config *string
var showVersion *bool
var exampleConfig *bool

// var help *bool

func ReadConfigPath() ConfigPath {
	var root = &cobra.Command{
		Use:  path.Base(os.Args[0]),
		Long: descripcion,
		Run:  func(cmd *cobra.Command, args []string) {},
	}

	config = root.PersistentFlags().StringP("config", "c", "", "location of the application's configuration file")
	showVersion = root.PersistentFlags().BoolP("version", "v", false, "version")
	exampleConfig = root.PersistentFlags().Bool("example", false, "generate configuration file")
	// help = root.PersistentFlags().BoolP("help", "h", false, fmt.Sprintf("help for %s", os.Args[0]))

	if err := root.Execute(); err != nil {
		slog.Error("error executing")
	}
	help, _ := root.Flags().GetBool("help")
	if help {
		os.Exit(0)
	}
	if *showVersion {
		fmt.Println(version)
		fmt.Println("Commit:", commit)
		fmt.Println("Branch:", branch)
		fmt.Println()
		os.Exit(0)
	}
	if *exampleConfig {
		fmt.Println(configExample)
		os.Exit(0)
	}
	return ConfigPath(*config)
}
