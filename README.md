# gh-mfmatt

A personal `gh` CLI extension for bootstrapping new GitHub repos with my conventions.

## Install

```bash
gh extension install matt-riley/gh-mfmatt
```

## Usage

```bash
gh mfmatt new <repo-name>
```

Interactively prompts for:
- **Repo type** — `go-api` (go-api-template), `bun-api` (elysia-crud-template), `neovim-plugin` (nvim-plugin-template), `astro-site` (astro-site-template), `ts-package` (ts-package-template), or `bare`
- **Visibility** — `private` or `public`

Then creates the repo under `matt-riley/<repo-name>` and clones it locally.

## Adding new template repos

Edit `cmd/new.go` and add an entry to `templateRepos`:

```go
var templateRepos = map[string]string{
    "bun-api":       "matt-riley/elysia-crud-template",
    "neovim-plugin": "matt-riley/nvim-plugin-template",
    "astro-site":    "matt-riley/astro-site-template",
    "go-api":        "matt-riley/go-api-template",
    "ts-package":    "matt-riley/ts-package-template",
}
```

## Development

```bash
go build ./...
gh extension install .
gh mfmatt new my-test-repo
```
