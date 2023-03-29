package gandi

import (
	"fmt"
	"strings"
	"time"
)

type Domain struct {
	FQDN              string `json:"fqdn"`
	DomainHref        string `json:"domain_href"`
	DomainRecordsHref string `json:"domain_records_href"`
}

type Record struct {
	Name   string   `json:"rrset_name"`
	TTL    int      `json:"rrset_ttl"`
	Type   string   `json:"rrset_type"`
	Values []string `json:"rrset_values"`
	Href   string   `json:"rrset_href,omitempty"`
}

func (r Record) String() string {
	var values []string
	for i := range r.Values {
		values = append(values, fmt.Sprintf("%s\t%d\t%s\t%s", r.Name, r.TTL, r.Type, r.Values[i]))
	}
	return strings.Join(values, "\n")
}

type Response struct {
	Message string `json:"message"`
}

type Snapshot struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Automatic bool      `json:"automatic"`
	Href      string    `json:"href"`
	ZoneData  []Record  `json:"zone_data"`
}

func (r Snapshot) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%v", r.ID, r.Name, r.CreatedAt, r.Automatic)
}

type CreateSnapshot struct {
	Name string `json:"name"`
}
