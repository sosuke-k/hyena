package main

import (
  "os"
  "os/user"
  // "os/exec"
	"path"
  "path/filepath"
  "bufio"
  // "fmt"
  "strconv"
  "log"
  // "bytes"
  "github.com/codegangsta/cli"
  "github.com/bitly/go-simplejson"
  // "github.com/codeskyblue/go-sh"
  // "github.com/kardianos/osext"
)

var projects = []string{}
var hyenaPath string //e.g. 'Users/name/.config/hyena'
var configPath string //e.g. 'Users/name/.config/hyena/config.json'
var srcDir string //e.g. 'Users/name/go/~~~'

func validError(errs ...error) error {
  for i, _ := range errs {
    if errs[i] != nil {
      return errs[i]
    }
  }
  return nil
}

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

  srcDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  } else {
    println(srcDir)
  }
}

func createConfig() {
  println(configPath)
  if err := os.MkdirAll(hyenaPath, 0777); err != nil {
    println(err)
  }
  js := simplejson.New()
  js.Set("length", 0)
  js.Set("projects", simplejson.New())
  w, err := os.Create(configPath)
  if err != nil {
    log.Fatal(err)
    log.Fatal(configPath)
  }
  defer w.Close()
  o, _ := js.EncodePretty()
  w.Write(o)
}

func loadConfig() (js simplejson.Json) {

  r, err := os.Open(configPath)
  if err != nil {
    log.Fatal(configPath + " is not found")
  } else {
    js, err := simplejson.NewFromReader(r)
    if err != nil {
      log.Fatal("cannot read " + configPath)
    } else {
      length, _ := js.Get("length").Int()
      for i := 0; i < length; i++ {
        project, _ := js.Get("projects").Get(strconv.Itoa(i)).String()
        projects = append(projects, project)
      }
    }
  }
  return
}

func save(projectName string) {
  // session := sh.NewSession()
  // var sessionOut bytes.Buffer
  // session.Stdout = &sessionOut
  //
  // chromeSrcDir := path.Join(srcDir, "chrome")
  // session.SetDir(chromeSrcDir)
  // session.Command("osascript", "-l JavaScript", "chrome_store_tabs.applescript").Run()
  //
  // println(sessionOut.String())

  println("still not implemented")
}

func load() {

}

func main() {
  app := cli.NewApp()
  // app.EnableBashCompletion = true
  app.Name = "hyena"
  app.Usage = "see help"
  app.Action = func(c *cli.Context) {
    println("This is the tool like hyena...")
    println("to get ore info, command 'hyena help'")
  }

  app.Commands = []cli.Command{
    {
      Name:      "init",
      Aliases:     []string{"i"},
      Usage:     "initialize hyena",
      Action: func(c *cli.Context) {
        println("Do you create " + configPath + " ? y/[n]")

        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        if err := scanner.Err(); err != nil {
            println(os.Stderr, "reading standard input:", err)
        }
        a := scanner.Text()
        if a == "y" {
          createConfig()
          println("config file was created")
        } else {
          println("please input y")
        }
      },
    }, // end init action definition
    {
      Name:      "ls",
      Aliases:     []string{"l"},
      Usage:     "show the list",
      Action: func(c *cli.Context) {
        loadConfig()
        for _, v := range projects {
          println(v)
        }
      },
    }, // end ls action definition
    {
      Name:      "add",
      Aliases:     []string{"a"},
      Usage:     "add the project",
      Action: func(c *cli.Context) {
        loadConfig()
        name := c.Args().First()
        if name == "" {
          println("please input project name")
        } else {
          projects = append(projects, name)
          js := simplejson.New()
          projectJson := simplejson.New()

          for i, v := range projects {
            projectJson.Set(strconv.Itoa(i), v)
        	}
          js.Set("length", len(projects))
          js.Set("projects", projectJson)
          w, err := os.Create(configPath)
          if err != nil {
            log.Fatal(err)
            log.Fatal(configPath)
          }
          defer w.Close()
          o, _ := js.EncodePretty()
          w.Write(o)
          println("created new project named " + name)
        }
      },
    }, // end add action definition
    {
      Name:      "save",
      Aliases:     []string{"s"},
      Usage:     "save the project",
      Action: func(c *cli.Context) {
        loadConfig()
        name := c.Args().First()
        if name == "" {
          println("please input project name")
        } else {
          save(name)
        }
      },
    }, // end save action definition
  }

  app.Run(os.Args)
}
