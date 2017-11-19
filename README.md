# license
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/license)](https://goreportcard.com/report/github.com/leighmcculloch/license)

License is a tool that provides license text for open source licenses.

## Install

### Source

```
go get 4d63.com/license
```

## Usage

```
$ license
License provides license text for 94 open source lisence.

Usage:

  license <license>

Example:

  license mit > LICENSE

Licenses:

  aal                 afl-3.0             agpl-3.0            apache-1.1
  apache-2.0          apl-1.0             apsl-2.0            artistic-1.0
  artistic-2.0        bsd-2               bsd-3               bsd-4
  bsl-1.0             catosl-1.1          cddl-1.0            cecill-2.1
  cnri-python         cpal-1.0            cpl-1.0             cua-opl-1.0
  cvw                 ecl-1.0             ecl-2.0             efl-1.0
  efl-2.0             entessa             epl-1.0             eudatagrid
  eupl-1.1            fair                frameworx-1.0       gfdl-1.2
  gfdl-1.3            gpl-1.0             gpl-2.0             gpl-3.0
  hpnd                intel               ipa                 ipl-1.0
  isc                 jabberpl            lgpl-2.0            lgpl-2.1
  lgpl-3.0            liliq-p-1.1         liliq-r&#43;        liliq-r-1.1
  lpl-1.0             lpl-1.02            lppl-1.3c           miros
  mit                 motosoto            mpl-1.0             mpl-1.1
  mpl-2.0             ms-pl               ms-rl               multics
  nasa-1.3            naumen              ncsa                ngpl
  nokia               nposl-3.0           ntp                 oclc-2.0
  ofl-1.1             ogtsl               opl-2.1             osl-1.0
  osl-2.1             osl-3.0             php-3.0             postgresql
  python-2.0          qpl-1.0             rpl-1.1             rpl-1.5
  rpsl-1.0            rscpl               simple-2.0          sissl
  sleepycat           spl-1.0             upl                 vsl-1.0
  w3c                 watcom-1.0          wxwindows           xnet
  zlib                zpl-2.0
```

## Example

```
$ license mit
Copyright (c) <YEAR>, <OWNER>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

```
$ license mit > LICENSE
```
