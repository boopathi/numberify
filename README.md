# numberify

Inspired from this tweet by Rob Pike, https://twitter.com/rob_pike/status/510944805292351489

Converts,

```
A quick brown fox jumps over the "lazy" dog!.
```

to

```
A q3k b3n f1x j3s o2r t1e "l2y" d1g.
```

## Install

```sh
$ go get github.com/boopathi/numberify
```

## Usage

```
$ cat MyText.txt | numberify
```

or

```
$ numberify < my_awesome_text.txt
```
