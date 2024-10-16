package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func main() {
	start := time.Now()

	// Open the file
	measurementsFile, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	// Close the file when we're done
	defer measurementsFile.Close()

	// Create a map to hold the measurements
	measurements := make(map[string]Measurement)

	// Create a scanner to read the file
	scanner := bufio.NewScanner(measurementsFile)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		semicolon := strings.Index(line, ";")

		// Skip the line if there's no semicolon (invalid line)
		if semicolon == -1 {
			continue
		}

		city := line[:semicolon]
		rawTemp := line[semicolon+1:]

		// Convert the temperature to a float
		temp, err := strconv.ParseFloat(rawTemp, 64)
		if err != nil {
			panic(err)
		}

		// Create a new measurement if the city is not in the map
		measurement, ok := measurements[city]
		if !ok {
			measurement = Measurement{Min: temp, Max: temp, Sum: temp, Count: 1}
		} else {
			// Add the temperature to the existing measurement
			measurement.Sum += temp
			measurement.Count++
			measurement.Min = math.Min(measurement.Min, temp)
			measurement.Max = math.Max(measurement.Max, temp)
		}

		measurements[city] = measurement
	}

	locations := make([]string, 0, len(measurements))
	for location := range measurements {
		locations = append(locations, location)
	}

	sort.Strings(locations)

	fmt.Print("{")
	for idx, location := range locations {
		measurement := measurements[location]
		mean := measurement.Sum / float64(measurement.Count)
		fmt.Printf("%s=%.1f/%.1f/%.1f", location, measurement.Min, mean, measurement.Max)

		if idx < len(locations)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")

	fmt.Printf(" took %s\n", time.Since(start))
}
