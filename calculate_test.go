package main

import (
	"database/sql"
	"testing"
)

func Test_calculatePopulation_empty(t *testing.T) {
	cities := []City{}
	got := calculatePopulationSumHandler(cities)
	want := map[string]int{}

	if len(got) != 0 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

func Test_calculatePopulation_oneCountry(t *testing.T) {
	cities := []City{
		{
			ID:3793,
			Name: sql.NullString{
				String: "New York",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "USA",
				Valid: true,
			},
			District: sql.NullString{
				String: "New York",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 8008278,
				Valid: true,
			},
		},
		{
			ID:3795,
			Name: sql.NullString{
				String: "Chicago",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "USA",
				Valid: true,
			},
			District: sql.NullString{
				String: "Illinois",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 2896016,
				Valid: true,
			},
		},
		{
			ID:3812,
			Name: sql.NullString{
				String: "Boston",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "USA",
				Valid: true,
			},
			District: sql.NullString{
				String: "Massachusetts",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 589141,
				Valid: true,
			},
		},
	}
	got := calculatePopulationSumHandler(cities)
	want := map[string]int{
		"USA":11493435,
	}

	if len(got) != 1 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["USA"] != want["USA"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

func Test_calculatePopulation_multiCountry(t *testing.T) {
	cities := []City{
		{
			ID:1564,
			Name: sql.NullString{
				String: "Omiya",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "JPN",
				Valid: true,
			},
			District: sql.NullString{
				String: "Saitama",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 441649,
				Valid: true,
			},
		},
		{
			ID:3795,
			Name: sql.NullString{
				String: "Chicago",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "USA",
				Valid: true,
			},
			District: sql.NullString{
				String: "Illinois",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 2896016,
				Valid: true,
			},
		},
		{
			ID:3812,
			Name: sql.NullString{
				String: "Boston",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "USA",
				Valid: true,
			},
			District: sql.NullString{
				String: "Massachusetts",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 589141,
				Valid: true,
			},
		},
		{
			ID:1812,
			Name: sql.NullString{
				String: "Toronto",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "CAN",
				Valid: true,
			},
			District: sql.NullString{
				String: "Ontario",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 688275,
				Valid: true,
			},
		},
		{
			ID:1538,
			Name: sql.NullString{
				String: "Kobe",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "JPN",
				Valid: true,
			},
			District: sql.NullString{
				String: "Hyogo",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 1425139,
				Valid: true,
			},
		},
	}
	got := calculatePopulationSumHandler(cities)
	want := map[string]int{
		"USA":3485157,
		"JPN":1866788,
		"CAN":688275,
	}

	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["USA"] != want["USA"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["JPN"] != want["JPN"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["CAN"] != want["CAN"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

func Test_calculatePopulation_invalid(t *testing.T) {
	cities := []City{
		{
			ID:3793,
			Name: sql.NullString{
				String: "New York",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "USA",
				Valid: true,
			},
			District: sql.NullString{
				String: "New York",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 8008278,
				Valid: true,
			},
		},
		{
			ID:3812,
			Name: sql.NullString{
				String: "Boston",
				Valid: true,
			},
			CountryCode: sql.NullString{
				String: "",
				Valid: false,
			},
			District: sql.NullString{
				String: "Massachusetts",
				Valid: true,
			},
			Population: sql.NullInt64{
				Int64: 589141,
				Valid: true,
			},
		},
	}
	got := calculatePopulationSumHandler(cities)
	want := map[string]int{
		"USA":8008278,
	}

	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["USA"] != want["USA"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}