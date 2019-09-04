package main

import (
  "log"

  "github.com/arschles/go-in-5-minutes/episode21/actions"
)

func main() {
  app := actions.App()
  if err := app.Serve(); err != nil {
    log.Fatal(err)
  }
}