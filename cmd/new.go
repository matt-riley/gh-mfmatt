package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// templateRepos maps repo types to GitHub template repo slugs.
// Add entries here as new template repos are created.
var templateRepos = map[string]string{
	"node-api": "matt-riley/elysia-crud-template",
}

var newCmd = &cobra.Command{
	Use:   "new <name>",
	Short: "Create and clone a new GitHub repo",
	Args:  cobra.ExactArgs(1),
	RunE:  runNew,
}

func runNew(_ *cobra.Command, args []string) error {
	name := args[0]

	var repoType, visibility string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Repo type").
				Options(
					huh.NewOption("Go service", "go-service"),
					huh.NewOption("Node API (elysia-crud-template)", "node-api"),
					huh.NewOption("Astro site", "astro-site"),
					huh.NewOption("Bare (no template)", "bare"),
				).
				Value(&repoType),
			huh.NewSelect[string]().
				Title("Visibility").
				Options(
					huh.NewOption("Private", "private"),
					huh.NewOption("Public", "public"),
				).
				Value(&visibility),
		),
	)

	if err := form.Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}

	ghArgs := []string{"repo", "create", "matt-riley/" + name, "--" + visibility}

	if tmpl, ok := templateRepos[repoType]; ok {
		ghArgs = append(ghArgs, "--template", tmpl)
	}

	ghArgs = append(ghArgs, "--clone")

	c := exec.Command("gh", ghArgs...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	return c.Run()
}
