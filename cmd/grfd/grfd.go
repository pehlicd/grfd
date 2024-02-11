package main

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"grfd/internal"
	"os"
	"text/template"
	"time"
)

var (
	log           = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}).With().Timestamp().Logger()
	versionString = "development"
	buildDate     string
	buildCommit   string
)

func exitWithError(msg string, err error) {
	log.Error().Msgf("%s: %v", msg, err)
	os.Exit(1)
}

func main() {
	var cfg internal.Config

	flag.StringVar(&cfg.GitLabAddr, "gitlab-addr", "", "GitLab address.")
	flag.StringVar(&cfg.GitLabToken, "gitlab-token", "", "GitLab API token.")
	flag.IntVar(&cfg.LastDays, "last-days", 1, "Last days to check for issues. Default is 1 day.")
	flag.StringVar(&cfg.TemplateFile, "template-file", "", "Template file. Should be a valid Go template file.")
	flag.StringVar(&cfg.OutputDir, "output-dir", "./", "Output directory. Default is current directory.")
	flag.BoolVar(&cfg.Insecure, "insecure", false, "Insecure mode to access GitLab.")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Verbose mode.")
	flag.Bool("version", false, "Print version and exit.")

	flag.Parse()

	if flag.Lookup("version").Value.String() == "true" {
		fmt.Printf("grfd version: %s\n", versionString)
		fmt.Printf("Build date: %s\n", buildDate)
		fmt.Printf("Build commit: %s\n", buildCommit)
		os.Exit(0)
	}

	client, err := internal.NewClient(cfg)
	if err != nil {
		exitWithError("error occurred while creating new client", err)
	}

	log.Info().Msgf("Starting Get Ready For the Daily...")

	cfg, err = client.GetCurrentUser(cfg)
	if err != nil {
		exitWithError("error occurred while getting current user", err)
	}

	log.Info().Msgf("Fetching last %d days of data for user %s from Gitlab...", cfg.LastDays, cfg.GitlabUsername)

	data, err := client.DataGenerator(cfg)
	if err != nil {
		exitWithError("error occurred while generating data", err)
	}

	templateFile := cfg.TemplateFile
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		exitWithError("error occurred while parsing template file", err)
	}

	outputFile := cfg.OutputDir + "/" + data.Date + ".md"
	file, err := os.Create(outputFile)
	if err != nil {
		exitWithError("error occurred while creating output file", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		exitWithError("error occurred while executing template", err)
	}

	log.Info().Msgf("You are ready for the daily!")
}
