# generate-git-commit

AI-powered Git commit message generator using Google's Gemini API. Analyzes your staged changes and automatically generates conventional commit messages following best practices.

## Features

- 🤖 AI-powered commit message generation using Gemini 2.5 Flash
- 📝 Follows conventional commit message format
- 🎯 Analyzes git diffs to understand the context of changes
- ⚡ Real-time streaming response
- 📋 Adheres to commit message style guide (imperative mood, proper length limits)

## Prerequisites

- Go 1.25.1 or higher
- Git installed and configured
- Google Gemini API key

## Installation

### Option 1: Install from source

```bash
# Clone the repository
git clone <repository-url>
cd generate-git-commit

# Install dependencies
go mod download

# Build and install globally
go install .
```

This will install the binary to your `$GOPATH/bin` (usually `~/go/bin` on Linux/Mac or `%USERPROFILE%\go\bin` on Windows).

Make sure your Go bin directory is in your PATH:
- **Windows**: Add `C:\Users\<YourUser>\go\bin` to your PATH
- **Linux/Mac**: Add `export PATH=$PATH:$(go env GOPATH)/bin` to your `.bashrc` or `.zshrc`

### Option 2: Build manually

```bash
# Build the binary
make build

# The binary will be in ./bin/generate-git-commit
# You can move it to any directory in your PATH
```

## Configuration

Set your Google Gemini API key as an environment variable:

```bash
# Linux/Mac
export GEMINI_API_KEY="your-api-key-here"

# Windows (PowerShell)
$env:GEMINI_API_KEY="your-api-key-here"

# Windows (CMD)
set GEMINI_API_KEY=your-api-key-here
```

**Note**: Currently the API key needs to be configured in `internal/gemini/apikey.go`. This will be improved to use environment variables in future versions.

## Usage

1. Stage your changes:
```bash
git add .
```

2. Run the commit message generator:
```bash
generate-git-commit
```

3. The AI will analyze your staged changes and generate a commit message following conventional commit format

## Commit Message Format

The generated commits follow this structure:

```
{type}({scope}): {subject}

{body}

{footer}
```

### Allowed Types
- `feat` → new feature
- `fix` → bug fix
- `docs` → documentation changes
- `style` → formatting, linting
- `refactor` → code restructure without behavior changes
- `test` → adding missing tests
- `chore` → maintenance tasks
- `init` → initial commit
- `rearrange` → file organization
- `update` → version updates, dependency changes

## Makefile Commands

```bash
make build          # Build the binary to ./bin/
make run            # Run directly without building
make clean          # Remove build artifacts
make build-linux    # Cross-compile for Linux
make build-windows  # Cross-compile for Windows
```

## Example

```bash
$ git add README.md
$ generate-git-commit

===================generate-git-command initialised===================
docs(readme): add comprehensive documentation

Add detailed README with installation instructions, usage guide,
and commit message format specifications. Include examples and
Makefile commands for easier project setup.
```

## Development

To modify the system prompt or commit message style guide, edit the `systemprompt` variable in `internal/gemini/streamContent.go`.