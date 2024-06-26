# README

This codebase has been generated by [Autostrada](https://autostrada.dev/).

## Before Getting Started

Create a file go.sum without which $ go mod tidy would never work

## Getting started

```
$ go mod tidy
$ go run ./cmd/web
```

Then visit [http://localhost:4444](http://localhost:4444) in your browser.

You can also start the application with live reload support by using the `run` task in the `Makefile`:

```
$ make run
```

## Project structure

Everything in the codebase is designed to be editable. Feel free to change and adapt it to meet your needs.

|                       |                                                            |
| --------------------- | ---------------------------------------------------------- |
| **`assets`**          | Contains the non-code assets for the application.          |
| `↳ assets/static/`    | Contains static UI files (images, CSS etc).                |
| `↳ assets/templates/` | Contains HTML templates.                                   |
| `↳ assets/efs.go`     | Declares an embedded filesystem containing all the assets. |

|                           |                                                                                                                                                                                        |
| ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`cmd/web`**             | Your application-specific code (handlers, routing, middleware, helpers) for dealing with HTTP requests and responses.                                                                  |
| `↳ cmd/web/errors.go`     | Contains helpers for managing and responding to error conditions.                                                                                                                      |
| `↳ cmd/web/handlers.go`   | Contains your application HTTP handlers.                                                                                                                                               |
| `↳ cmd/web/helpers.go`    | Contains helper functions for common tasks.                                                                                                                                            |
| `↳ cmd/web/main.go`       | The entry point for the application. Responsible for parsing configuration settings initializing dependencies and running the server. Start here when you're looking through the code. |
| `↳ cmd/web/middleware.go` | Contains your application middleware.                                                                                                                                                  |
| `↳ cmd/web/routes.go`     | Contains your application route mappings.                                                                                                                                              |
| `↳ cmd/web/server.go`     | Contains a helper functions for starting and gracefully shutting down the server.                                                                                                      |

|                         |                                                                                          |
| ----------------------- | ---------------------------------------------------------------------------------------- |
| **`internal`**          | Contains various helper packages used by the application.                                |
| `↳ internal/database/`  | Contains your database-related code (setup, connection and queries).                     |
| `↳ internal/funcs/`     | Contains custom template functions.                                                      |
| `↳ internal/request/`   | Contains helper functions for decoding HTML forms, JSON requests, and URL query strings. |
| `↳ internal/response/`  | Contains helper functions for rendering HTML templates and sending JSON responses.       |
| `↳ internal/validator/` | Contains validation helpers.                                                             |
| `↳ internal/version/`   | Contains the application version number definition.                                      |

## Application version

The application version number is defined in a `Get()` function in the `internal/version/version.go` file. Feel free to change this as necessary.

```
package version

func Get() string {
    return "0.0.1"
}
```

## Changing the module path

The module path is currently set to `github.com/codebyyogesh/lens_view`. If you want to change this please find and replace all instances of `github.com/codebyyogesh/lens_view` in the codebase with your own module path.
