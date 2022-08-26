# version
Compare version numbers

## Install

```golang
go get github.com/ishando/sojourn/version

// use in your .go code
import (
    "github.com/ishando/sojourn/version"
)
```

## Usage

Version compares version numbers, reporting the version is less than, greater than or equal to a reference version.

Create a new Version by passing in the version string:
```golang
  ver, err := NewVersion("1.2.3.4")
```

Compare version with the provided methods on the Version:
```golang
  if ver.LessThan("1.2") {...}
  if ver.Equal("1.2") {...}
  if ver.GreaterThan("1.2") {...}

  comp, err := ver.Compare("1.3")
```

## Limitations / Assumptions

The version string
- consists of only numbers and '.' as a seperator
- is of the format: `\d+(.\d+)*`
- cannot start or end with '.' or have consecutive '.'s
- can have any number of levels - assuming a reasonable usage, have not tested the limits of this

## Possible Enhancements

- Allow `-alpha`, `-beta` or other character components of the version number
