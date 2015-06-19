package main

import (
  "os"
  "os/user"
  // "os/exec"
	"path"
  "bufio"
  // "fmt"
  "strconv"
  "log"
  "bytes"
  "github.com/codegangsta/cli"
  "github.com/bitly/go-simplejson"
  "github.com/codeskyblue/go-sh"
)

var projects = []string{}
var hyena_path string
var config_path string

var tasks = []string{"ls", "clean", "laundry", "eat", "sleep", "code"}

func checkError( err ) {
  if e != nil {
    panic(e)
  }
}

func init() {
  hyena_path = os.Getenv("HYENA_DIR_PATH")
  if hyena_path == "" {
    usr, err := user.Current()
    if err != nil {
      log.Fatal(err)
    }
    hyena_path = path.Join(usr.HomeDir, ".config/hyena")
  }
  config_path = path.Join(hyena_path, "config.json")
}

func createConfig() {
  println(config_path)
  if err := os.MkdirAll(hyena_path, 0777); err != nil {
    println(err)
  }
  js := simplejson.New()
  js.Set("length", 0)
  js.Set("projects", simplejson.New())
  w, err := os.Create(config_path)
  if err != nil {
    log.Fatal(err)
    log.Fatal(config_path)
  }
  defer w.Close()
  o, _ := js.EncodePretty()
  w.Write(o)
}

func loadConfig() (js simplejson.Json) {

  r, err := os.Open(config_path)
  if err != nil {
    log.Fatal(config_path + " is not found")
  } else {
    js, err := simplejson.NewFromReader(r)
    if err != nil {
      log.Fatal("cannot read " + config_path)
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

func load() {
  // sh.Command("echo", "hello\tworld").Command("cut", "-f2").Run()
  // sh.Command("cd ~/go/src/example/").Run()
  // sh.Dir("~/go/src/example/").Run()
  // sh.Command("pwd", sh.Dir(os.Getenv("HOME") + "/go/src/example/")).Run()
  // sh.Command("cd /;ls -al").Run()
  session := sh.NewSession()
  var sessionOut bytes.Buffer
  session.Stdout = &sessionOut
  // session.SetEnv("BUILD_ID", "123")
  session.SetDir("/Users/katososuke/Documents/COMPUTER_SCIENCE/DATABASES/INFORMATION_RETRIEVAL/IIR/")
  // # then call cmd
  // session.Command("echo", "hello").Run()
  session.Command("ls").Run()
  // session.Command("open", "./FnTIR-Press-Kelly.pdf").Run()
  // applescript := "display dialog 'test'"
  session.Command("osascript", "/Users/katososuke/go/src/example/greet/sample.scpt").Run()
  println(sessionOut.String())
  // session.Command("/Applications/Preview.app/Contents/MacOS/Preview", "/Users/katososuke/Documents/COMPUTER_SCIENCE/DATABASES/INFORMATION_RETRIEVAL/IIR/FnTIR-Press-Kelly.pdf").Run()
  // # set ShowCMD to true for easily debug
  // session.ShowCmd = true


  // session.Stdout = &sessionOut
  // w, err := os.Create(configpath)
  // defer w.Close()
  // // o, _ := js.MarshalJSON()
  // // o, _ := js.Encode()
  // o, _ := js.EncodePretty()
  // w.Write(o)


  // exec.Command("cd", "/Users/katososuke/Documents/COMPUTER_SCIENCE/DATABASES/INFORMATION_RETRIEVAL/IIR/").Run()
  // cmd := exec.Command("ls", "-al")
  // var out bytes.Buffer
  // cmd.Stdout = &out
  // cmd_err := cmd.Run()
  // if cmd_err != nil {
  // 	log.Fatal(cmd_err)
  // }
  //
  // println(out.String())
  // exec.Command("open", "-a Preview", "./FnTIR-Press-Kelly.pdf").Run()
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
        println("Do you create " + config_path + " ? y/[n]")

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
          w, err := os.Create(config_path)
          if err != nil {
            log.Fatal(err)
            log.Fatal(config_path)
          }
          defer w.Close()
          o, _ := js.EncodePretty()
          w.Write(o)
          println("created new project named " + name)
        }
      },
    }, // end add action definition
  }

  app.Run(os.Args)
}
