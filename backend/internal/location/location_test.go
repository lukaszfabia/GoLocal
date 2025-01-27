package location_test

import (
	"backend/internal/location"
	"backend/internal/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
			if got != tt.want {
				t.Errorf("ParseCoords() = %v, want %v", got, tt.want)
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
				"address.road":         "Main Street",
				"address.house_number": "123",
				"display_name":         "Main Street 123, Warsaw, Poland",
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
				"address.house_number": "123",
				"display_name":         "Main Street 123, Warsaw, Poland",
			},
			want:    models.Address{},
			wantErr: true,
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
			"address.road": "Main Street",
			"address.house_number": "123",
			"display_name": "Main Street 123, Warsaw, Poland",
			"address.city": "Warsaw",
			"address.postcode": "00-001",
			"address.country": "Poland"
		}`))
	}))
	defer server.Close()

	// mockowanie
	originalGetUrl := location.GetUrl
	location.GetUrl = func(lon, lat string, format string) (*url.URL, error) {
		return url.Parse(server.URL)
	}
	defer func() { location.GetUrl = originalGetUrl }()

	location, err := location.FetchLocation("21.0122", "52.2297")
	if err != nil {
		t.Errorf("FetchLocation() error = %v", err)
		return
	}

	addr := &models.Address{
		Street:         "Main Street",
		StreetNumber:   "123",
		AdditionalInfo: "Main Street 123, Warsaw, Poland",
	}

	if fmt.Sprint(*addr) != fmt.Sprint(*location.Address) {
		t.Errorf("want %v, have %v", *addr, *location.Address)
	}

	coords := &models.Coords{
		Longitude: 21.0122,
		Latitude:  52.2297,
	}

	if fmt.Sprint(*coords) != fmt.Sprint(*location.Coords) {
		t.Errorf("want %v, have %v", *coords, *location.Coords)
	}

	// reset
	location.Address = nil
	location.Coords = nil

	expected := models.Location{
		City:    "Warsaw",
		Country: "Poland",
		Zip:     "00-001",
	}

	if fmt.Sprint(location) != fmt.Sprint(expected) {
		t.Errorf("FetchLocation() = %v\n, want %v", location, expected)
	}
}
