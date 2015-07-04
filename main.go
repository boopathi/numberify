package main

import (
  "os"
  "fmt"
  "github.com/boopathi/numberify/numberify"
  "github.com/codegangsta/cli"
  "bufio"
)

func Action(c *cli.Context) {
  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    fmt.Println(numberify.Numberify(string(line)))
  }
}

func main() {
  app := cli.NewApp()
  app.Name = "numberify"
  app.Usage = "numberify \"text to numberify\""
  app.Action = Action
  app.Run(os.Args)
}
