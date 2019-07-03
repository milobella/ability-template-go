# {{cookiecutter.project_title}}
{{cookiecutter.project_description}}

## Features
>
> TODO: describe features
>

## Prerequisites

- Having access to [gitlab.milobella.com](https://gitlab.milobella.com/milobella)
- Having ``golang`` installed [instructions](https://golang.org/doc/install)
- Having ``go dep`` installed [instructions](https://golang.github.io/dep/docs/installation.html)

## Build

```bash
$ dep ensure
$ go build -o bin/ability cmd/ability/main.go
```

## Run

```bash
$ bin/ability
```

## Requests example

>
> TODO: put some examples
>

## CHANGELOGS
- [Application changelog](CHANGELOG.md)
- [Helm chart changelog]({{cookiecutter.repository_name}}/helm/{{cookiecutter.helm_package}}/CHANGELOG.md)