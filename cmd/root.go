package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog"

	"github.com/rp-magrathea/cybertron/pkg/theme"
)

var (
	version       string
	versionCommit string
	inputFile     string
	stylesheet    string
	name          string
	wrapper       string
	macros        string
	templates     string
)

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().StringVarP(&name, "name", "n", "theme", "The name of the theme")
	buildCmd.PersistentFlags().StringVarP(&stylesheet, "stylesheet", "s", "assets/stylesheet.css", "The location of the stylesheet")
	buildCmd.PersistentFlags().StringVarP(&wrapper, "wrapper", "w", "wrapper.html", "The location of the wrapper.html")
	buildCmd.PersistentFlags().StringVarP(&macros, "macros-folder", "m", "macros", "The folder with the macros in it.")
	buildCmd.PersistentFlags().StringVarP(&templates, "templates-folder", "t", "html-templates", "The folder with the templates in it.")

	rootCmd.AddCommand(readCmd)
	readCmd.PersistentFlags().StringVarP(&inputFile, "input-file", "f", "", "The file to input")

	klog.InitFlags(nil)
	flag.Parse()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

var rootCmd = &cobra.Command{
	Use:   "cybertron",
	Short: "cybertron",
	Long:  `A tool to compile Jcink XML files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must specify a sub-command.")
		err := cmd.Help()
		if err != nil {
			klog.Error(err)
		}
		os.Exit(1)
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	Long:  `Builds a jcink xml theme from a set of files`,
	Run: func(cmd *cobra.Command, args []string) {

		output, err := theme.Construct(name, stylesheet, wrapper, macros, templates)
		if err != nil {
			klog.Error(err)
			os.Exit(1)
		}
		if output == nil {
			fmt.Println("got a nil return. something has gone very wrong.")
			os.Exit(1)
		}

		outputFile := fmt.Sprintf("%s.xml", name)
		f, err := os.Create(outputFile)
		if err != nil {
			klog.Fatal(err)
		}
		defer f.Close()
		_, err = f.WriteString(*output)
		if err != nil {
			klog.Fatal(err)
		}
		fmt.Printf("Successfully exported to %s\n", outputFile)
	},
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read",
	Long:  `Parses a theme file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := theme.Read(inputFile)
		if err != nil {
			klog.Error(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

// Execute the stuff
func Execute(VERSION string, COMMIT string) {
	version = VERSION
	versionCommit = COMMIT
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
