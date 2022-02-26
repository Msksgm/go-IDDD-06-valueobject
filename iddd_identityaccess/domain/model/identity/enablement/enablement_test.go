package enablement

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func TestNewEnablement(t *testing.T) {
	enabled := true
	startDate, err := time.ParseInLocation(timeFormat, "2020-01-01 00:00:00", jst)
	if err != nil {
		t.Fatal(err)
	}
	endDate, err := time.ParseInLocation(timeFormat, "2030-01-01 00:00:00", jst)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("hoge")
	fmt.Println(startDate)
	fmt.Println(endDate)

	got, err := NewEnablement(enabled, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}

	want := &Enablement{enabled: enabled, startDate: startDate, endDate: endDate}
	allowUnexported := cmp.AllowUnexported(Enablement{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
