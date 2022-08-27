package version

import (
	"fmt"
	"strconv"
	"strings"
)

// Version - stores a version string for comparison
// - Limits version string to be in the format: \d+[.\d+]*
// - Version string can contain any number of levels

type Version struct {
	asString string
	asArray  []int
}

var (
	InvalidVersion = fmt.Errorf("Invalid Version string, cannot be an empty string")
	InvalidElement = fmt.Errorf("Invalid Version string, non-numeric element")
	InvalidSeparatorUse = fmt.Errorf("Invalid Version string, possible missing element")
	InvalidIndex = fmt.Errorf("Version does not contain requested element")
)

func NewVersion(s string) (*Version, error) {
	if s == "" {
		return nil, InvalidVersion
	}

	v := &Version{
		asString: s,
		asArray:  []int{},
	}
	strArray := strings.Split(s, ".")

	for _, str := range strArray {
		if str == "" {
			return nil, InvalidSeparatorUse
		}
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, InvalidElement
		}
		v.asArray = append(v.asArray, val)
	}

	return v, nil
}

// String - print the version as a string
func (v Version) String() string {
	return v.asString
}

// Len - return the number of levels in the version
func (v Version) Len() int {
	return len(v.asArray)
}

// Part - return the part of the version for the given index.
// If index doesn't exist it returns an error
func (v Version) Part(index int) (int, error) {
	if index >= v.Len() {
		return -1, InvalidIndex
	}
	return v.asArray[index], nil
}

// LessThan - check if version is less than another version
func (v Version) LessThan(v2 *Version) bool {
	if v.Compare(v2) == -1 {
		return true
	}
	return false
}

// GreaterThan - check if version is greater than another version
func (v Version) GreaterThan(v2 *Version) bool {
	if v.Compare(v2) == 1 {
		return true
	}
	return false
}

// Equal - check if version is equal to another version
func (v Version) Equal(v2 *Version) bool {
	if v.Compare(v2) == 0 {
		return true
	}
	return false
}

// Compare - compare two version, return 0 if they are equal, -1 if v < v2, or 1 if v > v2
//   Will compare up to the level where one version is greater or less than the other
//   If one version string is longer than the other, and they are equal up to that level,
//     the shorter version will be considered the lower, ie 1.1.1 < 1.1.1.0
func (v Version) Compare(v2 *Version) int {
	// short circuit compare if whole version strings match 
	if v.String() == v2.String() {
		return 0
	}

	for i, val1 := range v.asArray {
		val2, err := v2.Part(i)
		if err != nil {
			// if v2 does not have part i, it is the lower
			return 1
		}
		if val1 < val2 {
			return -1
		} else if val1 > val2 {
			return 1
		}
	}

	// so far versions are the same, so if v2 is longer than v, it will be the later version
	if v.Len() < v2.Len() {
		return -1
	}

	// if we get here, the versions match
	return 0
}
