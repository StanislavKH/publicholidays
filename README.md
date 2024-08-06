
# PublicHolidays Package

The `publicholidays` package provides functionality to fetch and filter public holidays for a given year and country using the [Nager.Date API](https://date.nager.at/). It offers features like filtering holidays by type, fetching current year holidays, and more.

## Installation

To use this package, you need to install it using:

```bash
go install github.com/yourusername/publicholidays@latest
```

## Usage

Below is a brief overview of how to use the `publicholidays` package:

### Fetching Holidays

To fetch public holidays for a specific year and country code, use the `GetHolidays` function:

```go
import "github.com/yourusername/publicholidays"

holidays, err := publicholidays.GetHolidays(2024, "US")
if err != nil {
    fmt.Println("Error fetching holidays:", err)
    return
}
fmt.Println("Holidays:", holidays)
```

### Filtering Holidays by Type

You can filter holidays by type using the `FilterByType` method:

```go
filteredHolidays, err := holidays.FilterByType(publicholidays.Public, publicholidays.Bank)
if err != nil {
    fmt.Println("Error filtering holidays:", err)
    return
}
fmt.Println("Filtered Holidays:", filteredHolidays)
```

### Getting the Current Year

To get the current year, use the `GetCurrentYear` function:

```go
currentYear := publicholidays.GetCurrentYear()
fmt.Println("Current Year:", currentYear)
```

## Holiday Types

The package defines several constants for holiday types, which can be used to filter holidays:

- `Public`: Public holiday
- `Bank`: Bank holiday, banks and offices are closed
- `School`: School holiday, schools are closed
- `Authorities`: Authorities are closed
- `Optional`: Majority of people take a day off
- `Observance`: Optional festivity, no paid day off

## API

### Types

#### `Holiday`

```go
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
```

#### `Holidays`

```go
type Holidays []Holiday
```

### Functions

#### `GetHolidays`

Fetches public holidays for a given year and country code.

```go
func GetHolidays(year int, countryCode string) (Holidays, error)
```

- `year`: The year for which to fetch holidays.
- `countryCode`: The country code (e.g., "US").

#### `FilterByType`

Filters and returns only holidays with defined types.

```go
func (holidays Holidays) FilterByType(hTypes ...string) (Holidays, error)
```

- `hTypes`: A list of holiday types to filter.

#### `GetCurrentYear`

Fetches the current year.

```go
func GetCurrentYear() int
```

#### `ValidateType`

Checks if the provided type is valid.

```go
func ValidateType(hType string) error
```

- `hType`: The holiday type to validate.

