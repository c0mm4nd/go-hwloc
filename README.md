# go-hwloc

[![Build Status](https://travis-ci.com/maoxs2/go-hwloc.svg?branch=master)](https://travis-ci.com/maoxs2/go-hwloc)
[![Go Doc](https://godoc.org/godoc.org/github.com/maoxs2/go-hwloc?status.svg)](https://godoc.org/github.com/maoxs2/go-hwloc)

## Introduction
Golang binding for hwloc

## Installation

```bash
$ cd MyName/MyProject
$ git clone https://github.com/maoxs2/go-hwloc --recursive
$ cd go-hwloc/hwloc
$ ./autogen.sh
$ ./configure
$ make
$ sudo make install
$
$ # then you can `import hwloc "github.com/MyName/MyProject/go-hwloc" ` from other go files in your project

Warning: if you wanna build a static application/libary based on `go-hwloc`, [StaticBuild](https://github.com/open-mpi/hwloc/wiki/StaticBuild) of hwloc would be required.

```

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
