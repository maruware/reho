# reho

A tool for replacing host of URL

## Usage

```sh
$ reho https://www.google.com/search?q=hello example.com
https://example.com/search?q=hello
```

```sh
$ reho https://www.google.com/search?q=hello http://localhost:8080
http://localhost:8080/search?q=hello
```

## Install

See [Releases](/releases)

or

```
go install github.com/maruware/reho@latest
```
