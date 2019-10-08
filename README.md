# groupext

## What it does?

It sorts files in the working directory into directories named after file
extension.

## How it works?

```
$ cd /tmp
$ mkdir example; cd example
$ touch example.js example.go example.jpg
$ tree .
.
├── example.go
├── example.jpg
└── example.js

0 directories, 3 files
$ groupext .
mkdir "go"
mv "example.go" "go/example.go"
mkdir "jpg"
mv "example.jpg" "jpg/example.jpg"
mkdir "js"
mv "example.js" "js/example.js"
$ tree .
.
├── go
│   └── example.go
├── jpg
│   └── example.jpg
└── js
    └── example.js

3 directories, 3 files
```

## Installation

`git clone`, then `go install`.
