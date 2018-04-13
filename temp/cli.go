package main

import (
  "log"
  "os"
  "fmt"

  "github.com/urfave/cli"
)

func main() {
  var language string
  var name string

  app := cli.NewApp()
  app.Name = "hstore"
  app.Usage = "Distributed key-value store"
  app.Version = "0.1.0"

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "name, n",
      Value:       "hoanh",
      Usage:       "name to greet",
      Destination: &name,
    },
    cli.StringFlag{
      Name:        "lang, l",
      Value:       "english",
      Usage:       "language for the greeting",
      Destination: &language,
    },
  }

  app.Action = func(c *cli.Context) error {
    if language == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
