package humanize

import (
	"regexp"
	"strconv"
	"math"
	"bytes"
	"fmt"
	"time"
)

// Regex for peeling off groups of 3 digits from the end of a string
var commasRe = regexp.MustCompile(`(\d+)(\d{3})`)

//
// AddCommas returns a string representation of the given integer with
// commas added
//
func AddCommas(n int) string {
	working := strconv.Itoa(n)
	for {
		// normally, we would use the age-old regex: (\d)(?=(\d{3})+$),
		// but Go's rudimentary regex engine doesn't support (?=)
		// so we will use this loop to incrementally run a simplified regex
		// and add commas until the regex doesn't add any more, then return
		replaced := commasRe.ReplaceAllString(working, "$1,$2")
		if replaced == working {
			return replaced
		}
		working = replaced
	}
}

//
// Number returns a humanized string representation of the given int with
// a human-readable suffix.
//
func Number(n, decimals int) string {
	suffixes := [...]string{"", "k", "M", "B", "T"}
	if n < 1000 {
		return strconv.Itoa(n)
	}
	i := math.Floor(math.Log(float64(n)) / math.Log(1000))
	return strconv.FormatFloat(float64(n) / math.Pow(1000, i), 'f', decimals, 64) + suffixes[int(i)]
}


//
// Size returns a humanized string representation of the given int which
// is assumed to be a byte count.
//
func Size(n, decimals int) string {
	suffixes := [...]string{"", "K", "M", "G", "T", "P"}
	if n < 1024 {
		return strconv.Itoa(n)
	}
	i := math.Floor(math.Log(float64(n)) / math.Log(1024))
	return strconv.FormatFloat(float64(n) / math.Pow(1024, i), 'f', decimals, 64) + suffixes[int(i)]
}

//
// IsASCII counts the number of un-ASCII chars in the given string and compares
// that count to the given ratio threshold.
//
func IsASCII(str string, ratio float32) bool {
	// count unprintable chars
	cnt := 0
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 127 {
			cnt++
		}
	}
	return float32(cnt) / float32(len(str)) < ratio
}

//
// Render a byte-slice to a string with hex-editor like formatting.
//
func HexDump(slice []byte) string {
	var ret bytes.Buffer
	for i := 0; i < len(slice); i += 16 {
		// index
		ret.WriteString(fmt.Sprintf("%05d  ", i))

		// hex representation
		for j := 0; j < 16; j++ {
			if i+j < len(slice) {
				ret.WriteString(fmt.Sprintf("%02X ", slice[i+j]))
			} else {
				ret.WriteString("   ") // spacer for empty hex cells (2 hex + space)
			}
			if j == 7 {
				ret.WriteString(" ") // middle gutter (ends up 2-wide)
			}
		}
		ret.WriteString(" ")

		// char representation
		for j := 0; j < 16; j++ {
			if i+j < len(slice) {
				c := slice[i+j]
				if c >= 32 && c < 127 {
					ret.WriteString(fmt.Sprintf("%c", c))
				} else {
					ret.WriteString(".")
				}
			} else {
				ret.WriteString(" ") // spacer for empty char cells
			}
			if j == 7 {
				ret.WriteString("  ") // middle gutter
			}
		}
		ret.WriteString("\n")
	}

	return ret.String()
}

//
// ISO8601ms renders the given Time as an ISO8601 string with millisecond resolution.
// There is no direct method in time.Time for this particular format, so this
// function hides the odd format string used by golang.
//
func ISO8601ms(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z07:00")
}