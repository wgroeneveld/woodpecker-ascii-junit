package main

import (
	"codeberg.org/woodpecker-plugins/go-plugin"
	"context"
	"errors"
	"fmt"
	"github.com/joshdk/go-junit"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Settings struct {
	Path string
}

type Plugin struct {
	*plugin.Plugin
	Settings *Settings
}

func (p *Plugin) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "a glob path where the JUnit XML files are located",
			Sources:     cli.EnvVars("PLUGIN_PATH"),
			Destination: &p.Settings.Path,
		},
	}
}

func (p *Plugin) Execute(ctx context.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		log.Err(err).Msg("Cannot get working dir")
		return err
	}

	log.Debug().
		Str("Path", p.Settings.Path).
		Str("CWD", cwd).
		Msg("Starting ASCII JUnit Plugin")

	fullPath := p.fullPath(cwd)
	files, err := p.pathToFiles(err, fullPath)
	if err != nil {
		log.Err(err).Str("path", fullPath).Msg("Cannot retrieve files from path")
		return err
	}

	suites, err := junit.IngestFiles(files)
	if err != nil {
		log.Err(err).Msg("Cannot convert XML files to JUnit format")
		return err
	}

	p.processJunitSuites(suites)

	return nil
}

func (p *Plugin) fullPath(cwd string) string {
	path := cwd + "/" + p.Settings.Path
	if strings.HasPrefix(p.Settings.Path, "/") {
		path = cwd + p.Settings.Path
	}
	return path
}

func (p *Plugin) pathToFiles(err error, path string) ([]string, error) {
	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("no files found for path")
	}
	return files, nil
}

func (p *Plugin) processJunitSuites(suites []junit.Suite) {
	passed := 0
	failed := 0
	errored := 0
	skipped := 0
	total := 0
	var totalTime time.Duration = 0

	for _, suite := range suites {
		passed += suite.Totals.Passed
		failed += suite.Totals.Failed
		errored += suite.Totals.Error
		total += suite.Totals.Skipped
		passed += suite.Totals.Tests
		skipped += suite.Totals.Skipped
		totalTime += suite.Totals.Duration
	}

	totalTime = totalTime.Round(1 * time.Millisecond)

	fmt.Printf("JUnit Test Results: %d Test Suites Found\n", len(suites))
	fmt.Println("----------------------------------------")
	fmt.Println()
	fmt.Println("| Passed | Failed | Errored | Skipped | Total |")
	fmt.Println("_______________________________________________")
	fmt.Printf("| %s | %s | %s | %s | %s | \n", pad(6, passed), pad(6, failed), pad(7, errored), pad(7, skipped), pad(5, total))
	fmt.Println()
	fmt.Printf("Total time: %s\n", totalTime.String())
}

func pad(max int, nr int) string {
	padded := strconv.Itoa(nr)
	nrLen := len(padded)
	for i := 0; i < max-nrLen; i++ {
		padded += " "
	}
	return padded
}

func main() {
	p := &Plugin{
		Settings: &Settings{},
	}

	// Zerolog log levels are set by built-in setting "log_level": https://codeberg.org/woodpecker-plugins/go-plugin
	p.Plugin = plugin.New(plugin.Options{
		Name:        "woodpecker-ascii-junit",
		Description: "Woodpecker ASCII Junit XML Reporter",
		Flags:       p.Flags(),
		Execute:     p.Execute,
	})

	p.Run()
}
