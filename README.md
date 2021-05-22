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

- use Location header

```sh
$ reho -L https://g.co/cast/power example.com
https://example.com/chromecast/answer/10087405
```

## Install

See [Releases](https://github.com/maruware/reho/releases)

or

```
go install github.com/maruware/reho@latest
```
