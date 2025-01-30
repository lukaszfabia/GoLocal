package location_test

import (
	"backend/internal/location"
	"backend/internal/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestParseCoords(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]any
		want    models.Coords
		wantErr bool
	}{
		{
			name: "valid coords",
			input: map[string]any{
				"lon": "21.0122",
				"lat": "52.2297",
			},
			want: models.Coords{
				Longitude: 21.0122,
				Latitude:  52.2297,
			},
			wantErr: false,
		},
		{
			name: "invalid lon type",
			input: map[string]any{
				"lon": 21.0122,
				"lat": "52.2297",
			},
			want:    models.Coords{},
			wantErr: true,
		},
		{
			name: "invalid lat type",
			input: map[string]any{
				"lon": "21.0122",
				"lat": 52.2297,
			},
			want:    models.Coords{},
			wantErr: true,
		},
		{
			name: "invalid lon value",
			input: map[string]any{
				"lon": "invalid",
				"lat": "52.2297",
			},
			want:    models.Coords{},
			wantErr: true,
		},
		{
			name: "invalid lat value",
			input: map[string]any{
				"lon": "21.0122",
				"lat": "invalid",
			},
			want:    models.Coords{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := location.ParseCoords(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCoords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Latitude != tt.want.Latitude {
				t.Errorf("ParseCoords() = %v, want %v", got.Latitude, tt.want.Latitude)
			}

			if got.Longitude != tt.want.Longitude {
				t.Errorf("ParseCoords() = %v, want %v", got.Longitude, tt.want.Longitude)
			}
		})
	}
}

func TestParseAddr(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]any
		want    models.Address
		wantErr bool
	}{
		{
			name: "valid address",
			input: map[string]any{
				"address": map[string]any{
					"road":         "Main Street",
					"house_number": "123",
				},
				"display_name": "Main Street 123, Warsaw, Poland",
			},
			want: models.Address{
				Street:         "Main Street",
				StreetNumber:   "123",
				AdditionalInfo: "Main Street 123, Warsaw, Poland",
			},
			wantErr: false,
		},
		{
			name: "missing street",
			input: map[string]any{
				"address": map[string]any{
					"house_number": "123",
				},
				"display_name": "Main Street 123, Warsaw, Poland",
			},
			want:    models.Address{},
			wantErr: true,
		},
		{
			name: "missing house_number",
			input: map[string]any{
				"address": map[string]any{
					"road": "Main Street",
				},
				"display_name": "Main Street, Warsaw, Poland",
			},
			want: models.Address{
				Street:         "Main Street",
				StreetNumber:   "",
				AdditionalInfo: "Main Street, Warsaw, Poland",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := location.ParseAddr(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchLocation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"lon": "21.0122",
			"lat": "52.2297",
			"address": {
				"road": "Main Street",
				"house_number": "123",
				"city": "Warsaw",
				"postcode": "00-001",
				"country": "Poland"
			},
			"display_name": "Main Street 123, Warsaw, Poland"
		}`))
	}))
	defer server.Close()

	originalGetUrl := location.GetUrl
	location.GetUrl = func(lon, lat string, format string) (*url.URL, error) {
		return url.Parse(server.URL)
	}
	defer func() { location.GetUrl = originalGetUrl }()

	ch := location.FetchLocation("21.0122", "52.2297")

	result := <-ch
	if result.Err != nil {
		t.Errorf("FetchLocation() error = %v", result.Err)
		return
	}

	loc := result.Location

	expectedCords := &models.Coords{
		Longitude: 21.0122,
		Latitude:  52.2297,
	}

	expectedAddress := &models.Address{
		Street:         "Main Street",
		StreetNumber:   "123",
		AdditionalInfo: "Main Street 123, Warsaw, Poland",
	}
	if !reflect.DeepEqual(loc.Address, expectedAddress) {
		t.Errorf("FetchLocation() Address = %v, want %v", loc.Address, expectedAddress)
	}
	if loc.Coords.Longitude != expectedCords.Longitude || loc.Coords.Latitude != expectedCords.Latitude {
		t.Errorf("Bad coords: expected (%f, %f), got (%f, %f)", expectedCords.Longitude, expectedCords.Latitude, loc.Coords.Longitude, loc.Coords.Latitude)
	}

	if loc.City != "Warsaw" {
		t.Errorf("FetchLocation() Coords = %v, want %v", loc.City, "Warsaw")
	}

	if loc.Zip != "00-001" {
		t.Errorf("FetchLocation() Zip = %v, want %v", loc.Zip, "00-001")
	}
}

func TestGetUrl(t *testing.T) {
	var lon, lat string = "23.23", "32.12"

	have, err := location.GetUrl(lon, lat, "json")

	if err != nil {
		t.Error("Error occured during getting url")
	}

	expected, err := url.Parse(fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%s&lon=%s", lat, lon))
	if err != nil {
		t.Errorf("Error parsing expected URL: %v", err)
	}

	if expected.String() != have.String() {
		t.Errorf("GetUrl() want %s, have %s", expected.String(), have.String())
	}
}
