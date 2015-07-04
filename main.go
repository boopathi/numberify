package main

import (
  "os"
  "fmt"
  "github.com/boopathi/numberify/numberify"
  "github.com/codegangsta/cli"
)

func Action(c *cli.Context) {
  fmt.Println(numberify.Numberify("Text to check with a \"quot\""))
}

func main() {
  app := cli.NewApp()
  app.Name = "numberify"
  app.Usage = "numberify \"text to numberify\""
  app.Action = Action
  app.Run(os.Args)
}
