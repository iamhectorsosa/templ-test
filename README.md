# Deploy CV to GitHub Pages using Go

This project is a Go-based static site generator that builds and deploys a resume to GitHub Pages using data from a YAML file. It utilizes the templ tool for templating and tailwindcss for styling.

## Features

* Converts a YAML resume configuration into a static HTML page.
* Deploys the resume to GitHub Pages.
* TailwindCSS for responsive and minimal styling.
* Supports local development and preview modes.

## Rerequisites

* Go (1.23.2 or later)
* Node (with pnpm for `tailwindcss`)
* `templ` tool for Go (for templating)

## Installation

```bash
gh repo clone iamhectorsosa/templ-test
cd templ-test
```

## Commands

### Development

```bash
make dev
```

This will:
* Generate the static site using `templ` (with `--watch`).
* Watch for changes in TailwindCSS and source files.
* Run a local development server on port 3001.

### Preview

To build and preview on port 3000:

```bash
make preview
```

### Build

This will generate the `index.html` in the `.dist` folder:

```bash
make build
```

## Deployment

GitHub actions will build and deploy the contents of the `.dist` folder to GitHub Pages.

## Usage

The `data.yml` file contains all the information to be displayed in your resume. Example:
