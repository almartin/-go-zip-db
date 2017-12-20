package ZipDb

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

func TestLookup(t *testing.T) {
	const zip string = "06269"

	l, err := Lookup(zip)
	if err != nil {
		t.FailNow()
	}
	if l.State != "CT" {
		t.FailNow()
	}
}

func TestLookupFail(t *testing.T) {
	const zip string = "000000"

	l, err := Lookup(zip)
	if err == nil {
		t.Logf("%+v", l)
		t.FailNow()
	}
	t.Log(err)

}

func TestData(t *testing.T) {
	csvFile1, err := os.Open("tools/data.csv")
	if err != nil {
		t.Fail()
	}
	defer csvFile1.Close()

	r := csv.NewReader(bufio.NewReader(csvFile1))
	result, err := r.ReadAll()
	if err != nil {
		t.Fail()
	}
	t.Logf("Lines: %v", len(result))

	for i := range result {
		zip := result[i][1]

		lat64, err := strconv.ParseFloat(result[i][6], 32)
		if err != nil {
			continue
		}

		lon64, err := strconv.ParseFloat(result[i][7], 32)
		if err != nil {
			continue
		}

		lat := float32(lat64)
		lon := float32(lon64)

		name := result[i][3]
		state := result[i][4]

		l, err := Lookup(zip)
		if err != nil {
			t.Logf("Zip: %s", zip)
			t.Log(err)
			t.FailNow()
		}
		if lat != l.Latitude || lon != l.Longitude || name != l.Name || state != l.State {
			t.Logf("Expected Zip: %s, lat: %f, lon: %f, name: %s", zip, lat, lon, name)
			t.Logf("Got %+v", l)
			t.FailNow()
		}

	}

}
