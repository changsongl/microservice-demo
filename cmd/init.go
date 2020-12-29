package cmd

import (
	"fmt"
	"github.com/changsongl/microservice/logic/project"
	"github.com/changsongl/microservice/logic/project/web"
	"github.com/changsongl/microservice/vars"
	"github.com/spf13/cobra"
)

type initFlag struct {
	projectName string
}

func init() {
	flags := initFlag{}

	var initCmd = &cobra.Command{
		Use:     "init [web|worker] [-n=project-name]",
		Short:   "Init project",
		Long:    "Init web or work project",
		Version: vars.KcliInitVersion,
		Run:     initRunFunc(&flags.projectName),
	}

	initCmd.Flags().StringVarP(&flags.projectName, "name", "n", "default", "project name")

	kcliCmd.AddCommand(initCmd)
}

func initRunFunc(name *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var i project.Initializer
		if len(args) == 0 {
			i = web.NewInitializer()
		} else {
			t := args[0]
			if t == "web" {
				i = web.NewInitializer()
			} else if t == "worker" {
				fmt.Println("worker!!")
				return
			} else {
				fmt.Printf("invalid `%s`, should be web or worker", t)
				return
			}
		}

		fmt.Println(*name)
		err := project.Init(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
