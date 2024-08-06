package publicholidays

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Holiday represents a single public holiday
type Holiday struct {
	Date        string   `json:"date"`
	LocalName   string   `json:"localName"`
	Name        string   `json:"name"`
	CountryCode string   `json:"countryCode"`
	Fixed       bool     `json:"fixed"`
	Global      bool     `json:"global"`
	Counties    []string `json:"counties"`
	LaunchYear  int      `json:"launchYear"`
	Types       []string `json:"types"`
}

// Holidays represents a slice of Holiday
type Holidays []Holiday

// Constants for holiday types
const (
	Public      = "Public"
	Bank        = "Bank"        // Bank holiday, banks and offices are closed
	School      = "School"      // School holiday, schools are closed
	Authorities = "Authorities" // Authorities are closed
	Optional    = "Optional"    // Majority of people take a day off
	Observance  = "Observance"  // Optional festivity, no paid day off
)

// validTypes holds all valid holiday types
var validTypes = map[string]bool{
	Public:      true,
	Bank:        true,
	School:      true,
	Authorities: true,
	Optional:    true,
	Observance:  true,
}

// FilterByType filters and returns only holidays with defined types
func (holidays Holidays) FilterByType(hTypes ...string) (Holidays, error) {
	var filteredHolidays Holidays
	typeSet := make(map[string]bool)

	// Validate and populate the map with desired types
	for _, t := range hTypes {
		if err := ValidateType(t); err != nil {
			return nil, err
		}
		typeSet[t] = true
	}

	// Iterate through holidays and filter by the desired types
	for _, holiday := range holidays {
		for _, t := range holiday.Types {
			if typeSet[t] {
				filteredHolidays = append(filteredHolidays, holiday)
				break
			}
		}
	}
	return filteredHolidays, nil
}

// GetHolidays fetches public holidays for a given year and country code
func GetHolidays(year int, countryCode string) (Holidays, error) {
	// Construct the URL
	url := fmt.Sprintf("https://date.nager.at/api/v3/PublicHolidays/%d/%s", year, countryCode)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	// Decode the JSON response
	var holidays []Holiday
	if err := json.NewDecoder(resp.Body).Decode(&holidays); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return holidays, nil
}

// GetCurrentYear fetches the current year
func GetCurrentYear() int {
	return time.Now().Year()
}

// ValidateType checks if the provided type is valid
func ValidateType(hType string) error {
	if !validTypes[hType] {
		return fmt.Errorf("invalid holiday type: %s", hType)
	}
	return nil
}
