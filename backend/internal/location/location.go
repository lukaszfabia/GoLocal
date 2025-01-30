package location

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// https://nominatim.openstreetmap.org/reverse?format=json&lat=51.12&lon=17.05

var format = "json"

var GetUrl = func(lon, lat string, format string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=%s&lat=%s&lon=%s", format, lat, lon))
}

func ParseCoords(b map[string]any) (models.Coords, error) {
	var strLon, strLat string
	var ok bool

	strLon, ok = b["lon"].(string)
	if !ok {
		return models.Coords{}, fmt.Errorf("failed to get lon")
	}

	strLat, ok = b["lat"].(string)
	if !ok {
		return models.Coords{}, fmt.Errorf("failed to get lat")
	}

	lon, err := strconv.ParseFloat(strLon, 64)
	if err != nil {
		return models.Coords{}, fmt.Errorf("failed to parse on float64 lon")
	}

	lat, err := strconv.ParseFloat(strLat, 64)
	if err != nil {
		return models.Coords{}, fmt.Errorf("failed to parse on float64 lat")
	}

	return models.Coords{
		Longitude: lon,
		Latitude:  lat,
		Geom:      fmt.Sprintf("POINT(%.4f %.4f)", lon, lat),
	}, nil
}

func ParseAddr(b map[string]any) (models.Address, error) {

	address, ok := b["address"].(map[string]any)
	if !ok {
		return models.Address{}, fmt.Errorf("failed to get address map")
	}

	streetNumber, _ := address["house_number"].(string)

	info, _ := b["display_name"].(string)

	street, ok := address["road"].(string)
	if !ok {
		if street, ok = address["amenity"].(string); !ok {
			return models.Address{}, fmt.Errorf("failed to get street or amenity")
		}
	}

	return models.Address{
		Street:         street,
		StreetNumber:   streetNumber,
		AdditionalInfo: info,
	}, nil
}

func ParseBody(body []byte) (models.Location, error) {
	var jsonResponse map[string]any
	err := json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return models.Location{}, err
	}

	cords, err := ParseCoords(jsonResponse)

	if err != nil {
		return models.Location{}, err
	}

	addr, err := ParseAddr(jsonResponse)

	if err != nil {
		return models.Location{}, err
	}

	address, ok := jsonResponse["address"].(map[string]any)
	if !ok {
		return models.Location{}, fmt.Errorf("failed to get address map")
	}

	city, ok := address["city"].(string)
	if !ok {
		return models.Location{}, fmt.Errorf("no city provided")
	}

	zip, ok := address["postcode"].(string)
	if !ok {
		return models.Location{}, fmt.Errorf("no zip provided")
	}

	country, ok := address["country"].(string)
	if !ok {
		return models.Location{}, fmt.Errorf("no country provided")
	}

	location := models.Location{
		City:      city,
		Country:   country,
		Zip:       zip,
		Address:   &addr,
		AddressID: addr.ID,
		Coords:    &cords,
		CoordsID:  cords.ID,
	}

	return location, nil
}

func FetchLocation(lon, lat string) <-chan Result {
	ch := make(chan Result, 1)

	go func() {
		url, err := GetUrl(lon, lat, format)
		if err != nil {
			ch <- Result{Location: models.Location{}, Err: fmt.Errorf("invalid url: %w", err)}
			return
		}

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		req, err := http.NewRequest(http.MethodGet, url.String(), nil)
		if err != nil {
			ch <- Result{Location: models.Location{}, Err: fmt.Errorf("failed to make request: %w", err)}
			return
		}

		req.Header.Set("User-Agent", "go-local-api/1.0")

		resp, err := client.Do(req)
		if err != nil {
			ch <- Result{Location: models.Location{}, Err: fmt.Errorf("failed to send request: %w", err)}
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			ch <- Result{Location: models.Location{}, Err: fmt.Errorf("not 2xx response code: %s", resp.Status)}
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ch <- Result{Location: models.Location{}, Err: fmt.Errorf("failed to read body: %w", err)}
			return
		}

		location, err := ParseBody(body)
		if err != nil {
			log.Println(err)
			ch <- Result{Location: models.Location{}, Err: fmt.Errorf("failed to parse location: %w", err)}
			return
		}

		ch <- Result{Location: location, Err: nil}
	}()

	return ch
}

type Result struct {
	Location models.Location
	Err      error
}
