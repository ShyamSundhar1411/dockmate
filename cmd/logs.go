/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/container/v1"
)

var fllowFlag bool;
// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "A brief description of your command",
	Args: cobra.MinimumNArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0];
		fmt.Println("logs called");
	},
}

func init() {
	
	logsCmd.Flags().BoolVarP(&fllowFlag, "follow", "f", false, "Follow the log output");
	rootCmd.AddCommand(logsCmd);

}

func getContainerLogs(containerName string, fllowFlag bool){
	cmdArgs := []string{"logs"};
	if fllowFlag{
		cmdArgs = append(cmdArgs, "-f");
	}
	cmdArgs = append(cmdArgs, containerName);
	cmd := exec.Command("docker",cmdArgs...);
	out,err := cmd.Output();
	if err != nil{
		log.Fatalln(err)
		return;
	}
	fmt.Println(string(out));
}
