# Minimal Word Count examples
This is the official [Minimal Word Count](https://github.com/apache/beam/blob/master/sdks/go/examples/minimal_wordcount/minimal_wordcount.go) example.

The example was modified with:

- Changed `main()` function to `getPipeline()` and made it return the pipeline instead of executing it. This way it can be passed to `app.Run()`;
- Some fixed strings were transformed into configuration, just as an example;

How to run:

``` sh
cd examples/localrunner/
go run . -app-log-level=debug -app-log-human -with-metrics
```
