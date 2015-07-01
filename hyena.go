package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/codegangsta/cli"
	"github.com/sosuke-k/hyena/app/acrobat"
	"github.com/sosuke-k/hyena/app/atom"
	"github.com/sosuke-k/hyena/app/chrome"
	"github.com/sosuke-k/hyena/app/kobito"
	"github.com/sosuke-k/hyena/util/git"
	"github.com/sosuke-k/hyena/util/jxa"
	"github.com/sosuke-k/hyena/util/log"
	"github.com/sosuke-k/hyena/util/pm"
)

var hyenaPath string  //e.g. 'Users/name/.config/hyena'
var configPath string //e.g. 'Users/name/.config/hyena/config.json'

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
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("====================")
	hyenaLogger.Println("  Starting hyena... ")
	hyenaLogger.Println("====================")
}

func hyenaInit(c *cli.Context) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run init command")
	println("Do you create " + configPath + " ? y/[n]")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		hyenaLogger.Printf("reading standard input:: %v", err)
		hyenaLogger.Println("")
	}
	a := scanner.Text()
	if a == "y" {
		hyenaLogger.Println("y input")
		jxa.Compile()
		pm.Init(configPath)
		println("config file was created")
	} else {
		hyenaLogger.Println("not y input")
		println("please input y")
	}
	hyenaLogger.Println("finished init command")
}

func hyenaLs(c *cli.Context) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run ls command")
	projects := pm.Load(configPath)
	projectString := ""
	for _, v := range projects {
		projectString += v + "\t"
	}
	fmt.Fprintln(os.Stdout, projectString)
	hyenaLogger.Println("finished ls command")
}

func hyenaAdd(c *cli.Context) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run add command")
	name := c.Args().First()
	if name == "" {
		hyenaLogger.Println("scan empty as input")
		println("please input project name")
	} else {
		hyenaLogger.Println("scan " + name + " as input")
		pm.Add(configPath, name)
		println("created new project named " + name)
		git.Init(path.Join(hyenaPath, name))
	}
	hyenaLogger.Println("finished add command")
}

func save(projectName string) {
	projectPath := path.Join(hyenaPath, projectName)
	chrome.Save(path.Join(projectPath, "chrome.json"))
	acrobat.Save(path.Join(projectPath, "acrobat.json"))
	kobito.Save(path.Join(projectPath, "kobito.json"))
	atom.Save(path.Join(projectPath, "atom.json"))
	git.Commit(projectPath, "hyena save "+projectName, false)
}

func hyenaSave(c *cli.Context) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run save command")
	name := c.Args().First()
	if name == "" {
		hyenaLogger.Println("scan empty as input")
		println("please input project name")
	} else {
		hyenaLogger.Println("scan " + name + " as input")
		save(name)
	}
	hyenaLogger.Println("finished save command")
}

func restore(projectName string) {
	projectPath := path.Join(hyenaPath, projectName)
	chrome.Restore(path.Join(projectPath, "chrome.json"))
	acrobat.Restore(path.Join(projectPath, "acrobat.json"))
	kobito.Restore(path.Join(projectPath, "kobito.json"))
	atom.Restore(path.Join(projectPath, "atom.json"))
	git.Commit(projectPath, "hyena restore "+projectName, true)
}

func hyenaRestore(c *cli.Context) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run restore command")
	name := c.Args().First()
	if name == "" {
		hyenaLogger.Println("scan empty as input")
		println("please input project name")
	} else {
		hyenaLogger.Println("scan " + name + " as input")
		restore(name)
	}
	hyenaLogger.Println("finished restore command")
}

func main() {
	app := cli.NewApp()
	app.Name = "hyena"
	app.Usage = "see help"
	app.Action = func(c *cli.Context) {
		hyenaLogger := logger.GetInstance()
		hyenaLogger.Println("to run only hyena command")
		println("This is the tool like hyena...")
		println("to get more info, command 'hyena help'")
		hyenaLogger.Println("finished hyena command")
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize hyena",
			Action:  hyenaInit,
		}, {
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "show the list",
			Action:  hyenaLs,
		}, {
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add the project",
			Action:  hyenaAdd,
		}, {
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "save the project",
			Action:  hyenaSave,
		}, {
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "restore the project",
			Action:  hyenaRestore,
		},
	}

	app.Run(os.Args)
}
