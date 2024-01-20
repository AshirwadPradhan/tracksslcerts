package helpers

import (
	"testing"
	"time"
)

func TestCalcDaysDiff(t *testing.T) {
	d := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	u := time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC)

	actual := CalcDaysDiff(d, u)
	exp := int64(1)
	if actual != exp {
		t.Fatalf("Expected %d got %d", exp, actual)
	}

	d = time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	u = time.Date(2024, 1, 21, 0, 0, 0, 0, time.UTC)

	actual = CalcDaysDiff(d, u)
	exp = int64(-1)
	if actual != exp {
		t.Fatalf("Expected %d got %d", exp, actual)
	}

}
