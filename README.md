# flatjson

Make JSON one level deep

## Description

flatjson is a utility to normalize JSON files. Given a JSON object or an array
it will return an array of JSON objects one level deep, enriched with
additional properties that establish the original relation.

## Installation

flatjson is available using the standard `go get` command.

Install by running:

    go get github.com/cbhvn/flatjson

## Example usage

```shell
>echo [{"a":{"b":"c"}}, {"d":"e","f":{"g":"h"}}] | flatjson
[
  {
    "_id": 2,
    "_parentid": 1,
    "_parentref": ""
  },
  {
    "_id": 3,
    "_parentid": 2,
    "_parentref": "a",
    "b": "c"
  },
  {
    "_id": 4,
    "_parentid": 1,
    "_parentref": "",
    "d": "e"
  },
  {
    "_id": 5,
    "_parentid": 4,
    "_parentref": "f",
    "g": "h"
  }
]
```

## Usage

```shell
>flatjson --help
Usage of flatjson.exe:
  -f, --file string       input file
      --id int            begin id numbering with (default 1)
  -o, --output string     output file
      --parentid string   parent id property name (default "_parentid")
      --propid string     id property name (default "_id")
      --ref string        parent ref property name (default "_parentref")
```
