package parseduration

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseUnit(t *testing.T) {
	testCases := []struct {
		tCase string
		expM  int64
		expP  string
		err   error
	}{
		{
			"", 0, "", ErrInvalidTimeUnit,
		},
		{
			"1d", 1, "d", nil,
		},
		{
			"24h", 24, "h", nil,
		},
		{
			"-60s", -60, "s", nil,
		},
		{
			"124ms", 124, "ms", nil,
		},
		{
			"-32w", -32, "w", nil,
		},
		{
			"204us", 204, "us", nil,
		},
		{
			"1a", 0, "", ErrInvalidTimeUnit,
		},
		{
			"123", 0, "", ErrInvalidTimeUnit,
		},
		{
			"-d", -1, "d", nil,
		},
		{
			"-", 0, "", ErrInvalidTimeUnit,
		},
		{
			"--", 0, "", ErrInvalidTimeUnit,
		},
	}
	for _, tc := range testCases {
		m, p, e := parseUnit(tc.tCase)
		require.Equal(t, tc.expM, m)
		require.Equal(t, tc.expP, p)
		require.Equal(t, tc.err, e)
	}

}

func TestParseDuration(t *testing.T) {
	testCases := []struct {
		tCase string
		expected  time.Duration
		err   error
	}{
		{"1d", time.Duration(time.Hour * 24), nil},
		{"+d", time.Duration(time.Hour * 24), nil},
		{"24h", time.Duration(time.Hour * 24), nil},
		{"2w", time.Duration(time.Hour * 168 * 2), nil},
		{"30m", time.Duration(time.Minute * 30), nil},
		{"22s", time.Duration(time.Second * 22), nil},
		{"-124ms", time.Duration(time.Millisecond * -124), nil},
		{"34us", time.Duration(time.Microsecond * 34), nil},
		{"1ns", time.Duration(time.Nanosecond), nil},
		{"2w3d12h32m60s172ms1us74ns", time.Duration(time.Hour * 168 * 2 + time.Hour * 24 * 3 + time.Hour * 12 + time.Minute * 32 + time.Second * 60 + time.Millisecond * 172 + time.Microsecond * 1 + time.Nanosecond * 74), nil},
		{"123", 0, ErrInvalidTimeUnit},
		{"2d2d", 0, ErrDuplicatedTimeUnit},
		{"3w4d3h2s3h", 0, ErrDuplicatedTimeUnit},
		{"-d", time.Duration(time.Hour * -24), nil},
		{"-", 0, ErrInvalidTimeUnit},
		{"--", 0, ErrInvalidTimeUnit},
		{"", 0, ErrInvalidTimeUnit},
	}
	for _, tc := range testCases {
		d, e := ParseDuration(tc.tCase)
		require.Equal(t, tc.expected, d)
		require.Equal(t, tc.err, e)
	}
}