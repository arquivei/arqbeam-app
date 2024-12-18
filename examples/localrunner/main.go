package main

import (
	app "github.com/arquivei/arqbeam-app"
)

var (
	version = "development"
	config  struct {
		app.Config

		Source string `default:"gs://apache-beam-samples/shakespeare/kinglear.txt"`
		Output string `default:"wordcounts.txt"`
	}
)

func main() {
	app.Bootstrap(version, &config)

	p := getPipeline()

	app.Run(p)
}
