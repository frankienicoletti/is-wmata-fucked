package wmata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BusIncindent represents a rail incident.
type BusIncindent struct {
	DateUpdated    time.Time `json:"DateUpdated"`
	Description    string    `json:"Description,omitempty"`
	IncidentID     string    `json:"IncidentID"`
	IncidentType   string    `json:"IncidentType"`
	RoutesAffected []string  `json:"RoutesAffected"`
}

// UnmarshalJSON is a custom unmarshaler.
func (r *BusIncindent) UnmarshalJSON(b []byte) error {
	v := struct {
		Date           string   `json:"DateUpdated"`
		Description    string   `json:"Description,omitempty"`
		IncidentID     string   `json:"IncidentID"`
		IncidentType   string   `json:"IncidentType"`
		RoutesAffected []string `json:"RoutesAffected"`
	}{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	dateUpdated, err := time.Parse("2006-01-02T15:04:05", v.Date)
	if err != nil {
		return err
	}
	*r = BusIncindent{
		DateUpdated:    dateUpdated,
		Description:    v.Description,
		IncidentID:     v.IncidentID,
		IncidentType:   v.IncidentType,
		RoutesAffected: v.RoutesAffected,
	}
	return nil
}

// GetBusIncidents retrieves bus incidents.
func GetBusIncidents() ([]BusIncindent, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s?api_key=%s", baseURL, "/Incidents.svc/json/BusIncidents", apiKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	v := struct {
		BusIncidents []BusIncindent
	}{}
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, err
	}
	return v.BusIncidents, nil
}
