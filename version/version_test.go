package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require := require.New(t)

	testcases := map[string]struct{
		ver string
		err error
	}{
		"success, single": {
			ver: "1",
		},
		"success, levelled": {
			ver: "1.1.0",
		},
		"error - end with .": {
			ver: "1.1.0.",
			err: InvalidSeparatorUse,
		},
		"error - start with .": {
			ver: ".1.1.0",
			err: InvalidSeparatorUse,
		},
		"error - missing number": {
			ver: "1..0",
			err: InvalidSeparatorUse,
		},
		"error - alpha": {
			ver: "1.1.a",
			err: InvalidElement,
		},
	}

	for tn, tc := range testcases {
		t.Run(tn, func(t *testing.T){
			v, err := NewVersion(tc.ver)
			if tc.err != nil {
				require.Nil(v, "a version should not be returned")
				require.Error(err, "should return error when expected")
				require.Equal(tc.err, err, "should return the expected error")
			} else {
				require.NotNil(v, "a Version should be returned")
				require.Equal(tc.ver, v.String(), "Version should have the expected string")
			}
		})
	}
}

func TestCompare(t *testing.T) {
	require := require.New(t)

	testCases := map[string]struct{
		v1   string
		v2   string
		rslt int
		lt   bool
		gt   bool
		eq   bool
	}{
		"v1 equal v2": {
			v1: "1.1.1",
			v2: "1.1.1",
			rslt: 0,
			lt: false,
			gt: false,
			eq: true,
		},
		"v1 longer v2": {
			v1: "1.1.1.0",
			v2: "1.1.1",
			rslt: 1,
			lt: false,
			gt: true,
			eq: false,
		},
		"v1 shorter v2": {
			v1: "1.1.1",
			v2: "1.1.1.0",
			rslt: -1,
			lt: true,
			gt: false,
			eq: false,
		},
		"v1 less than v2": {
			v1: "1.1.1",
			v2: "1.2.1",
			rslt: -1,
			lt: true,
			gt: false,
			eq: false,
		},
		"v1 greater than v2": {
			v1: "1.1.3",
			v2: "1.1.1",
			rslt: 1,
			lt: false,
			gt: true,
			eq: false,
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			ver1, _ := NewVersion(tc.v1)
			require.NotNil(ver1, "should create Version")
			ver2, _ := NewVersion(tc.v2)
			require.NotNil(ver2, "should create Version")

			require.Equal(tc.rslt, ver1.Compare(ver2), "Compare should give the expected result")
			require.Equal(tc.lt, ver1.LessThan(ver2), "LessThan should give the expected result")
			require.Equal(tc.gt, ver1.GreaterThan(ver2), "GreaterThan should give the expected result")
			require.Equal(tc.eq, ver1.Equal(ver2), "Equal should give the expected result")
		})
	}
}

func TestPart(t *testing.T) {
	require := require.New(t)

	ver, _ := NewVersion("1.2.3.4") 
	
	for i := 0; i < 5; i++ {
		p, err := ver.Part(i)
		if i < 4 {
			require.NoError(err, "should return version part")
			require.Equal(i+1, p, "should return the expected value")
		} else {
			require.Error(err, "should return error for invalid index")
			require.Equal(InvalidIndex, err, "should return the expected error")
		}
	}
}