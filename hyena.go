package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path"
	// "fmt"
	// "bytes"
	"log"

	"github.com/codegangsta/cli"
	// "github.com/codeskyblue/go-sh"
	// "github.com/kardianos/osext"
	"github.com/sosuke-k/hyena/app/acrobat"
	"github.com/sosuke-k/hyena/app/chrome"
	"github.com/sosuke-k/hyena/app/kobito"
	"github.com/sosuke-k/hyena/util/pm"
)

var projects = []string{}
var hyenaPath string  //e.g. 'Users/name/.config/hyena'
var configPath string //e.g. 'Users/name/.config/hyena/config.json'
var srcDir string     //e.g. 'Users/name/go/~~~'

func validError(errs ...error) error {
	for i := range errs {
		if errs[i] != nil {
			return errs[i]
		}
	}
	return nil
}

// this function not correspond to 'hyena init' command
// call this function before all command
func init() {
	hyenaPath = os.Getenv("HYENA_DIR_PATH")
	if hyenaPath == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		hyenaPath = path.Join(usr.HomeDir, ".config/hyena")
	}
	configPath = path.Join(hyenaPath, "config.json")
}

func save(projectName string) {
	projectPath := path.Join(hyenaPath, projectName)
	chrome.Save(path.Join(projectPath, "chrome.json"))
	acrobat.Save(path.Join(projectPath, "acrobat.json"))
	kobito.Save(path.Join(projectPath, "kobito.json"))
}

func restore(projectName string) {
	projectPath := path.Join(hyenaPath, projectName)
	chrome.Restore(path.Join(projectPath, "chrome.json"))
	acrobat.Restore(path.Join(projectPath, "acrobat.json"))
}

func main() {
	app := cli.NewApp()
	// app.EnableBashCompletion = true
	app.Name = "hyena"
	app.Usage = "see help"
	app.Action = func(c *cli.Context) {
		println("This is the tool like hyena...")
		println("to get more info, command 'hyena help'")
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize hyena",
			Action: func(c *cli.Context) {
				println("Do you create " + configPath + " ? y/[n]")

				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					println("reading standard input:", err)
				}
				a := scanner.Text()
				if a == "y" {
					pm.Init(configPath)
					println("config file was created")
				} else {
					println("please input y")
				}
			},
		}, // end init action definition
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "show the list",
			Action: func(c *cli.Context) {
				projects = pm.Load(configPath)
				projectString := ""
				for _, v := range projects {
					projectString += v + "\t"
				}
				fmt.Fprintln(os.Stdout, projectString)
			},
		}, // end ls action definition
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add the project",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if name == "" {
					println("please input project name")
				} else {
					pm.Add(configPath, name)
					println("created new project named " + name)
				}
			},
		}, // end add action definition
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "save the project",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if name == "" {
					println("please input project name")
				} else {
					save(name)
				}
			},
		}, // end save action definition
		{
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "restore the project",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if name == "" {
					println("please input project name")
				} else {
					restore(name)
				}
			},
		}, // end restore action definition
	}

	app.Run(os.Args)
}
