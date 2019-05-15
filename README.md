# go-hwloc

[![Build Status](https://travis-ci.com/c0mm4nd/go-hwloc.svg?branch=master)](https://travis-ci.com/c0mm4nd/go-hwloc)
[![Go Doc](https://godoc.org/godoc.org/github.com/c0mm4nd/go-hwloc?status.svg)](https://godoc.org/github.com/c0mm4nd/go-hwloc)

## Introduction

Golang binding for static hwloc

## Ready

Install hwloc from your system package manager or compile from the source in https://github.com/open-mpi/hwloc

For example:
```bash
# Static Build
git clone https://github.com/open-mpi/hwloc && cd hwloc
./autogen.sh
./configure --enable-static --disable-shared LDFLAGS="-static"
make LDFLAGS=-all-static
sudo make install
sudo ldconfig /usr/local/lib
```

## Usage

```go
import hwloc "github.com/c0mm4nd/go-hwloc"
```

## Extra

- Q: Why not assemble hwloc into go lib?

A: hwloc is a very powerful and platform-based native library, so comparing to provide a simple version, I prefer to let developers compile it by their own hands or the binary for their platform.

- Q: Why build fails when using bazel?

A: Bazel compiling project within a sandbox, which disallow linker find the "external" library, hwloc, so you should write `BUILD.bazel` and `WORKSPACE` for hwloc before using this go binding lib. Actually, it is what tensorflow does, and you can learn from [their practice](https://github.com/tensorflow/tensorflow/tree/master/third_party/hwloc)

# Hwloc

## Introduction

The Hardware Locality (hwloc) software project aims at easing the process of
discovering hardware resources in parallel architectures. It offers
command-line tools and a C API for consulting these resources, their locality,
attributes, and interconnection. hwloc primarily aims at helping
high-performance computing (HPC) applications, but is also applicable to any
project seeking to exploit code and/or data locality on modern computing
platforms.

hwloc is actually made of two subprojects distributed together:

  * The original hwloc project for describing the internals of computing nodes.
 It is described in details starting at section Hardware Locality (hwloc)
 Introduction.
  * The network-oriented companion called netloc (Network Locality), described
 in details starting with section Network Locality (netloc).

See also the Related pages tab above for links to other sections.

Netloc may be disabled, but the original hwloc cannot. Both hwloc and netloc
APIs are documented after these sections.
