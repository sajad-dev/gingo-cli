# Gingo CLI

A lightweight and modular Golang framework built with Gin and GORM for fast and scalable web development.

## Overview

Gingo CLI is a command-line tool that helps developers quickly scaffold and manage Gingo framework projects. It provides an easy way to initialize new projects with the Gingo framework template and manage project configurations.

## Features

- 🚀 **Quick Project Setup**: Initialize new Gingo projects with a single command
- 📦 **Template Management**: Clone and customize project templates from GitHub
- 🔧 **Configurable**: Support for custom repository URLs and project names
- 📋 **Version Management**: Display current CLI and framework versions
- 🎯 **Flexible**: Support for different template versions

## Installation

### Prerequisites

- Go 1.22 or higher
- Git installed and configured
- Internet connection for cloning templates

### Install from Source

```bash
git clone https://github.com/sajad-dev/gingo-cli.git
cd gingo-cli
go build -o gingo cmd/gingo/main.go
```

### Install Binary

```bash
go install github.com/sajad-dev/gingo-cli/cmd/gingo@latest
```

## Usage

### Initialize a New Project

Create a new Gingo project in the current directory:

```bash
gingo new
```

#### Options

- `-n, --name`: Project directory name (default: "gingo-project")
- `-r, --repo`: Git repository URL to use as template (default: "gingo-project")
- `-v, --version`: Template version to use (default: "latest")

#### Examples

```bash
# Create project with default settings
gingo new

# Create project with custom name
gingo new -n my-awesome-api

# Create project with specific version
gingo new -n my-project -v v0.1.0

# Create project with custom repository
gingo new -r github.com/myorg/my-template -n my-project
```

### Check Version

Display the current CLI version and requirements:

```bash
gingo version
```

### Help

Display help information:

```bash
gingo help
```

## Project Structure

After running `gingo new`, you'll get a project structure like:

```
my-project/
├── go.mod
├── go.sum
├── main.go
├── gear              # Executable script
├── config/
├── controllers/
├── middleware/
├── models/
├── routes/
└── ...
```

## Configuration

The CLI uses environment variables for configuration. Create a `.env` file in your project root:

```env
APP_NAME=MyApp
DESCRIPTION=My awesome Gingo application
AUTHOR=Your Name
DEBUG=true
```

### Available Configuration Options

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_NAME` | "GINGO" | Application name |
| `DESCRIPTION` | "A lightweight and modular Golang framework..." | App description |
| `AUTHOR` | "Sajad pourajam" | Author name |
| `DEBUG` | "true" | Debug mode |

## Requirements

- Go >= v1.22
- Gingo >= v0.1
- Gingo-helper >= v0.1

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/sajad-dev/gingo-cli.git
cd gingo-cli

# Install dependencies
go mod download

# Build the CLI
go build -o gingo cmd/gingo/main.go

# Run tests
go test ./...
```

### Project Structure

```
gingo-cli/
├── cmd/gingo/          # Main CLI entry point
├── internal/
│   ├── cli/            # CLI framework setup
│   ├── command/        # Command implementations
│   │   ├── newpr/      # New project command
│   │   └── versions/   # Version command
│   └── config/         # Configuration management
├── go.mod
├── go.sum
└── requirements.txt
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/command/newpr/
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Color](https://github.com/fatih/color) - Colored terminal output
- [GoDotEnv](https://github.com/joho/godotenv) - Environment variable loading
- [Testify](https://github.com/stretchr/testify) - Testing toolkit

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Sajad Pourajam**

- GitHub: [@sajad-dev](https://github.com/sajad-dev)

## Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/sajad-dev/gingo-cli/issues) page
2. Create a new issue with detailed information
3. Join our community discussions

## Changelog

### v0.1.0

- Initial release
- Basic project scaffolding
- Template cloning functionality
- Version management
- Configuration system

---

**Happy coding with Gingo! 🚀**
