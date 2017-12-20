package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	csvFile1, err := os.Open("tools/data.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile1.Close()

	r := csv.NewReader(bufio.NewReader(csvFile1))

	out, err := os.Create("zip.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fmt.Fprintf(out, "package ZipDb\n")
	//	fmt.Fprintf(out, "import(\"log\")\n")
	fmt.Fprintf(out, "type GeoLocation struct {\n")
	fmt.Fprintf(out, "Latitude  float32\n"+
		"Longitude float32\n"+
		"Name string\n"+
		"State string\n"+
		"}\n")

	fmt.Fprintf(out, "var m map[string]GeoLocation\n")

	fmt.Fprintf(out, "func init(){\n")
	fmt.Fprintf(out, "m = make(map[string]GeoLocation)\n")

	result, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	log.Printf("Lines: %v", len(result))

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

		fmt.Fprintf(out, "m[\"%s\"] = GeoLocation{ Latitude: %f, Longitude: %f, Name: \"%s\", State: \"%s\"}\n",
			zip, lat, lon, name, state)
	}
	fmt.Fprintf(out, "}\n")

	out.Sync()

}
