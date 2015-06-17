package main

import (
  "os"
  "os/user"
  // "os/exec"
	"path"
  "log"
  "bytes"
  "github.com/codegangsta/cli"
  "github.com/bitly/go-simplejson"
  "github.com/codeskyblue/go-sh"
)

type configuration struct {
	APIKey   string
	City     string
	Imperial bool
	Lang     string
}

func main() {
  app := cli.NewApp()
  app.Name = "greet"
  app.Usage = "fight the loneliness!"
  app.Action = func(c *cli.Context) {
    var config configuration
    configpath := os.Getenv("SLOTH_CONFIG_PATH")
  	if configpath == "" {
  		usr, err := user.Current()
  		if err != nil {
  			log.Fatal(err)
  		}
  		configpath = path.Join(usr.HomeDir, ".config/sloth/config.json")
  	}
  	config.APIKey = ""
  	config.City = "New York"
  	config.Imperial = false
  	config.Lang = "en"
    // println("configpath", configpath)
    // home_path := os.Getenv("HOME")
    // config_path := home_path + "/.config/sloth/config.json"
    //
    r, err := os.Open(configpath)
    if err != nil {
        log.Fatal(err)
    }
    js, err := simplejson.NewFromReader(r)
    //
    // b, err := js.Get("hoge").Bool()
    // m, _ := js.Get("piyo").Map()
    // a := js.Get("piyo").Get("foo").MustArray()
    // s := js.GetPath("piyo", "bar").MustString()
    // println(b, m, a, s)
    //
    //
    // println("Hello friend!")
    //
    // projects := js.Get("projects").MustArray()
    // println(len(projects))
    // name, err := js.Get("projects").GetIndex(0).Get("name").String()
    // println(name)
    json, e := simplejson.NewJson([]byte("[1,2]"))
    // json.set("name", "sample2")
    // projects.append(json)
    if e != nil {
      panic(e)
    }
    js.Set("name", json)
    // for i, v := range js.Get("projects").MustArray() {
      // name, err := js.Get("projects").GetIndex(i).String("name")
      // if err != nil {
      //   panic(err)
      // }
      // println(name)
    //   // var json Json
    //   // simplejson.Unmarshal([]byte(jsonSrc), json)
    //   json, e := simplejson.NewJson(v.([]byte))
    //   name := json.Get("name").MustString()
	  //   println(i, name)
    //   println(e)
    // }
    // println(err)
    // println(js)
    // println(home_path)
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
    session.Stdout = &sessionOut
    w, err := os.Create(configpath)
    defer w.Close()
    // o, _ := js.MarshalJSON()
    // o, _ := js.Encode()
    o, _ := js.EncodePretty()
    w.Write(o)
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

  app.Run(os.Args)
}
