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

Version provide a representation of version numbers to allow easy comparison, reporting if the version is less than, greater than, or equal to another version.

Create a new Version by passing in the version string:
```golang
  ver, err := NewVersion("1.2.3.4")
```

Compare version with the provided methods on the Version:
```golang
LessThan(*Version) bool
Equal(*Version) bool
GreaterThan(*Version) bool
Compare(*Version) int
```

`Compare` function return
- -1 if the target Version is less than the provided Version
- 0 if the target Version is equal to the provided Version
- 1 if the target Version is greater than the provided Version

Example usage:
```golang
  ver, _ := NewVersion("1.1")

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
