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
	"github.com/sosuke-k/hyena/app/preview"
	"github.com/sosuke-k/hyena/util/git"
	"github.com/sosuke-k/hyena/util/jxa"
	"github.com/sosuke-k/hyena/util/log"
	"github.com/sosuke-k/hyena/util/pm"
	"github.com/sosuke-k/hyena/util/server"
	"github.com/sosuke-k/hyena/util/sh"
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
		srcDir := path.Join(os.Getenv("HOME"), "src/github.com/sosuke-k/hyena")
		sh.Execute(srcDir, "bower", []string{"install"})
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

func hyenaExecWithProject(c *cli.Context) {
	cmd := c.Command.Name
	name := c.Args().First()
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run " + cmd + " command")
	if name == "" {
		hyenaLogger.Println("scan empty as input")
		println("please input project name")
	} else {
		hyenaLogger.Println("scan " + name + " as input")
		switch cmd {
		case "add":
			hyenaAdd(name)
		case "delete":
			hyenaDelete(name)
		case "save":
			hyenaSave(name)
		case "restore":
			hyenaRestore(name)
		default:
			hyenaLogger.Println("cmd not apply any case")
		}
	}
	hyenaLogger.Println("finished " + cmd + " command")
}

func hyenaAdd(name string) {
	pm.Add(configPath, name)
	projectDir := path.Join(hyenaPath, name)
	rep := gyena.Repository{Dir: projectDir}
	rep.Init()
	println("created new project named " + name)
}

func hyenaDelete(name string) {
	pm.Delete(configPath, name)
	println("deleted the project named " + name)
}

func hyenaSave(projectName string) {
	projectPath := path.Join(hyenaPath, projectName)
	chrome.Save(path.Join(projectPath, "chrome.json"))
	acrobat.Save(path.Join(projectPath, "acrobat.json"))
	kobito.Save(path.Join(projectPath, "kobito.json"))
	atom.Save(path.Join(projectPath, "atom.json"))
	preview.Save(path.Join(projectPath, "preview.json"))
	rep := gyena.Repository{Dir: projectPath}
	rep.Commit("hyena save "+projectName, false)
}

func hyenaRestore(projectName string) {
	projectPath := path.Join(hyenaPath, projectName)
	chrome.Restore(path.Join(projectPath, "chrome.json"))
	acrobat.Restore(path.Join(projectPath, "acrobat.json"))
	kobito.Restore(path.Join(projectPath, "kobito.json"))
	atom.Restore(path.Join(projectPath, "atom.json"))
	preview.Restore(path.Join(projectPath, "preview.json"))
	rep := gyena.Repository{Dir: projectPath}
	rep.Commit("hyena restore "+projectName, true)
}

func hyenaBrowser(c *cli.Context) {
	hyenaLogger := logger.GetInstance()
	hyenaLogger.Println("to run browser command")
	fmt.Fprintln(os.Stdout, "to run browser command")

	port := "8187"
	sh.Execute(os.Getenv("HOME"), "open", []string{"http://localhost:" + port})
	server.Listen(port)

	hyenaLogger.Println("finished browser command")
	fmt.Fprintln(os.Stdout, "finished browser command")
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
			Action:  hyenaExecWithProject,
		}, {
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete the project",
			Action:  hyenaExecWithProject,
		}, {
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "save the project",
			Action:  hyenaExecWithProject,
		}, {
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "restore the project",
			Action:  hyenaExecWithProject,
		}, {
			Name:    "browser",
			Aliases: []string{"b"},
			Usage:   "browsing project history",
			Action:  hyenaBrowser,
		},
	}

	app.Run(os.Args)
}
