# fakegen: A single binary CLI for generating fake JSON data

## Installation

```bash
$ go install github.com/multiprocessio/fakegen@latest
```

## Usage

Pass the number of rows and columns you want and `fakegen` will give
you a JSON array of objects with that many rows and unique columns.

```
$ fakegen --rows 10 --cols 1000
[
  ... 10 fake objects with 1000 fake columns ...
]
```
