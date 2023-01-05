# GoSwanson

A CLI tool to interact with https://ron-swanson-quotes.herokuapp.com

## Build and Run
```
~$ go build
~$ ./goswanson -h
```

## Usage
```
A CLI tool to interact with https://ron-swanson-quotes.herokuapp.com

Usage:
./goswanson [Flag]

Flags:
    -h, --help             This help text
    -s, --search string    Search Quote
    -No flag-              Get Random Quote
```

## Example
```
~$ ./goswanson --search cholesterol
.. * What's cholesterol?

~$ ./goswanson
.. Fishing relaxes me. It's like yoga, except I still get to kill something.
```