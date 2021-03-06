# Amazon CLI

The CLI interface to you favourite retailer written in [Go](http://golang.org/).

![Amazon CLI](https://raw.githubusercontent.com/matthistuff/amazon/master/amazon.gif)

## Installation

If you have Go >= 1.4 installed already

```
go get github.com/matthistuff/amazon
```

(Mac) OS X >= 10.7 and Linux (x86, amd64) users can get binaries from the [releases page](https://github.com/matthistuff/amazon/releases). Currently I can not test other OSes than OS X and Ubuntu Linux, feel welcome to help!

Currently you'll need an AWS access key and secret key to use the Amazon Product Advertisement API. Amazon CLI expects the two env variables `AMAZON_ACCESS_KEY` and `AMAZON_SECRET_KEY` to be present in order to work.