# Go Licenses

`go-licenses` is a small CLI which allows you to export and check the licenses of your dependencies from your `go.mod` file.

## Installation

```sh
go install github.com/aboutyou/go-licenses
```

## Usage

The project uses [licenseclassifier](github.com/google/licenseclassifier) to classify the licenses
of each package.

### Exporting the licenses

To export all modules and their packages run one of the following commands:

```sh
go-licenses export csv > licenses.csv
```

or

```sh
go-licenses export json > licenses.json
```

### Check the licenses

The `check` command validates that all the modules in the project use compatible licenses.

```sh
go-licenses check
```

By default the following licenses are allowed:

- Apache20
- BSD2Clause
- BSD3Clause
- MIT
- Facebook3Clause
- Ruby
- PHP301
- Python20

## Ideas

In Go 1.18 there will be a new package [debug/buildinfo](https://pkg.go.dev/debug/buildinfo@master) which will allow us to offer the commands also for compiled Go programms.

This will allow to run license compliance terms on your shipped binaries.


## Maintainers

- [Henri Beck](https://github.com/HenriBeck)