# arqbeam-app

An Apache Beam application wrapper using [go-app](https://github.com/arquivei/go-app).

This further extends go-app project to create Apache Beam applications.

TL;DR

```go
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/arquivei/arqbeam-app"
)

var (
	version = "development"
	config struct {
		app.Config
		YourCustomConfig string
		SomeSecretConfig int           `secret:"true"`
		Timeout          time.Duration `default:"3m"`
	}
)

func main() {
	app.Bootstrap(version, &config)

	p := getPipeline()

	app.Run(p)
}

```

All Apache Beam flags are mapped to the config struct by embedding `app.Config`. Use `-h` to see all values and their default values.

Comments, discussions, issues and pull-requests are welcomed.
