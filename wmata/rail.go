package wmata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// RailIncindent represents a rail incident.
type RailIncindent struct {
	DateUpdated   time.Time `json:"Updated"`
	Description   string    `json:"Description,omitempty"`
	IncidentID    string    `json:"IncidentID"`
	IncidentType  string    `json:"IncidentType"`
	LinesAffected []string
}

// UnmarshalJSON is a custom unmarshaler.
func (r *RailIncindent) UnmarshalJSON(b []byte) error {
	v := struct {
		Date              string `json:"DateUpdated"`
		Description       string `json:"Description,omitempty"`
		IncidentID        string `json:"IncidentID"`
		IncidentType      string `json:"IncidentType"`
		LinesAffectedList string `json:"LinesAffected,omitempty"`
	}{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	linesAffected := strings.Split(v.LinesAffectedList, "; ")
	dateUpdated, err := time.Parse("2006-01-02T15:04:05", v.Date)
	if err != nil {
		return err
	}
	*r = RailIncindent{
		DateUpdated:   dateUpdated,
		Description:   v.Description,
		IncidentID:    v.IncidentID,
		IncidentType:  v.IncidentType,
		LinesAffected: linesAffected,
	}
	return nil
}

// GetRailIncidents retrieves rail incidents.
func GetRailIncidents() ([]RailIncindent, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s?api_key=%s", baseURL, "/Incidents.svc/json/Incidents", apiKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	v := struct {
		Incidents []RailIncindent
	}{}
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, err
	}
	return v.Incidents, nil
}
