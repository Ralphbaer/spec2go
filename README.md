# Spec2Go - 🚧 Work in Progress (WIP) 🚧

**spec2go is currently under development. Features may be incomplete and documentation is subject to change. Use in production environments is not recommended until the official release.**

## What is Spec2Go?

**Spec2Go** is powerful command-line tool automates the creation of Golang applications directly from OpenAPI 3.0 and Go pkg.go.dev/text/template specifications. With its default setting to generate code following the principles of Hexagonal Architecture, Spec2Go ensures high modularity and ease of testing, offering a solid foundation for building robust applications.

## Key Features

- **Direct Integration with OpenAPI 3.0**: Convert your OpenAPI 3.0 specifications into well-structured Golang code seamlessly.
- **Hexagonal Architecture**: Begin your projects with an architecture that supports and enhances maintainability and scalability.
- **Customizable Code Generation**: Utilize Go pkg.go.dev/text/template templates to shape the output code to better fit your project requirements.
- **CLI-Based Utility**: Run as a standalone command-line tool, allowing integration into any part of your development pipeline.

## Installation

Spec2Go is easy to install with Go. Ensure you have Go installed on your system (version 1.16 or later is recommended). You can install Spec2Go globally using the following command:

```sh
go install github.com/Ralphbaer/spec2go@latest
```

This command will install spec2go globally on your machine, allowing you to run it from anywhere within your terminal.

## Quick Start Guide

### Setting Up Your Project
Ensure your OpenAPI 3.0 YAML files are ready. For instance, you might have an API specification for a ledger component within a financial application located at ./path/to/your/openai/file.yml.

### Generating Code
To generate Golang code using Spec2Go, use the following syntax:

```sh
spec2go generate -i ./path/to/your/openai/file.yml -o .
```

This command will instruct Spec2Go to:

- -i (--input): Path to your OpenAPI specification file.
- -o (--output): Output directory where the generated Golang code will be placed.

### Customizing Output with Templates
To customize the output, you can modify or create new Go tmpl templates.

A simple Go tmpl template might look like this:

```
{{define "main"}}
package main

import (
	"log"
	"net/http"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/{{.SourceFolder}}"
	//
	{{.PackageName}} "./{{.SourceFolder}}"
)

func main() {
	log.Println("Server started on :3000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
{{end}}
```

Adjust the template to include any specific configurations or coding standards your project requires.

## Advanced Configuration
You can further configure Spec2Go through various command-line flags to adapt the tool to your specific workflows and preferences. For example, adding verbose output, configuring API versioning, or specifying different template paths.

## Contributing
We welcome contributions from the community! If you have improvements or bug fixes, please fork the repository and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
Thanks to the OpenAPI Initiative for maintaining the specification.
Contributors who have helped shape Spec2Go into a reliable tool for developers.
We hope Spec2Go will make your development process more efficient and enjoyable!

