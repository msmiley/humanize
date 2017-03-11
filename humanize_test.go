package humanize

import (
	"testing"
	"time"
)


func TestAddCommas(t *testing.T) {
	actual := AddCommas(1234)
	expected := "1,234"
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
	actual = AddCommas(123456789)
	expected = "123,456,789"
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestNumber(t *testing.T) {
	actual := Number(1234, 2)
	expected := "1.23k"
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestSize(t *testing.T) {
	actual := Size(1024, 2)
	expected := "1.00K"
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestIsASCII1(t *testing.T) {
	actual := IsASCII("jldsfdskl", 0.30)
	expected := true
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestIsASCII2(t *testing.T) {
	actual := IsASCII("jkk\x03\xfd\xef\x12\x00\x02", 0.30)
	expected := false
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestHexDump(t *testing.T) {
	b := []byte{2,3,4,5,6,54,43,23,2,3,6,8,4,7,8,65,7,43,4,3,7,88,64,45,234}
	actual := HexDump(b)
	expected := `00000  02 03 04 05 06 36 2B 17  02 03 06 08 04 07 08 41  .....6+.  .......A
00016  07 2B 04 03 07 58 40 2D  EA                       .+...X@-  .       ` + "\n"
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestISO8601ms(t *testing.T) {
	actual := ISO8601ms(time.Date(2009, time.November, 10, 23, 1, 2, 3000000, time.UTC))
	expected := "2009-11-10T23:01:02.003Z"
	if expected != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}