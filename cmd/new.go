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
	"bun-api":       "matt-riley/elysia-crud-template",
	"neovim-plugin": "matt-riley/nvim-plugin-template",
	"astro-site":    "matt-riley/astro-site-template",
	"go-api":        "matt-riley/go-api-template",
	"ts-package":    "matt-riley/ts-package-template",
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
					huh.NewOption("Go API (go-api-template)", "go-api"),
					huh.NewOption("Bun API (elysia-crud-template)", "bun-api"),
					huh.NewOption("Neovim plugin (nvim-plugin-template)", "neovim-plugin"),
					huh.NewOption("Astro site (astro-site-template)", "astro-site"),
					huh.NewOption("TS package (ts-package-template)", "ts-package"),
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
