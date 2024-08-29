/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"
	"log"
	"github.com/spf13/cobra"
)
var listAllContainersFlag bool
var listAllRunningContainersFlag bool
var listAllStoppedContainersFlag bool
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if listAllContainersFlag{
			listAllContainers()
		}else if listAllRunningContainersFlag{
			listAllRunningContainers()
		} else if listAllStoppedContainersFlag{
			listAllStoppedContainers()
		} else{
			fmt.Println("Please specify a flag to list containers")
		}


	},
}
func listAllContainers(){
	cmd := exec.Command("docker", "ps", "-a")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(out))
}

func listAllRunningContainers(){
	cmd := exec.Command("docker","ps")
	out,err := cmd.Output()
	if err!=nil{
		log.Fatalln(err)
		return
	}
	fmt.Println(string(out))
}
func listAllStoppedContainers(){
	cmd := exec.Command("docker", "ps", "-f", "status=exited")
	out,err := cmd.Output()
	if err!=nil{
		log.Fatalln(err)
		return
	}
	fmt.Println(string(out))
}
func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&listAllContainersFlag, "all", "a", false, "List all containers",)
	lsCmd.Flags().BoolVarP(&listAllRunningContainersFlag, "running", "r", false, "List running containers",)
	lsCmd.Flags().BoolVarP(&listAllStoppedContainersFlag, "stopped", "s", false, "List stopped containers",)
}
