package gandi

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidTTL = errors.New("invalid TTL")
var ErrInvalidType = errors.New("invalid Type")
var ErrDifferentTTL = errors.New("different TTLs")
var ErrCanNotParse = errors.New("can not parse")

// refer https://api.gandi.net/docs/livedns/
//
// API is available : https://api.gandi.net/v5/livedns/dns/rrtypes
func IsSupportedType(typeStr string) bool {
	types := []string{"A", "AAAA", "ALIAS", "CAA", "CDS", "CNAME", "DNAME", "DS",
		"KEY", "LOC", "MX", "NAPTR", "NS", "OPENPGPKEY", "PTR", "RP", "SPF",
		"SRV", "SSHFP", "TLSA", "TXT", "WKS"}

	for _, typ := range types {
		if typeStr == typ {
			return true
		}
	}
	return false
}

// refer https://api.gandi.net/docs/livedns/
func isValudTTL(ttl int) bool {
	if ttl < 300 {
		return false
	}
	if 2592000 < ttl {
		return false
	}
	return true
}

func ParseRecord(record string) (*Record, error) {
	arg := strings.TrimSpace(record)
	args := strings.Fields(arg)

	var name, typ, value string
	var ttl = 300

	if len(args) == 3 {
		name = args[0]
		typ = args[1]
		value = args[2]
	} else if len(args) == 4 {
		name = args[0]
		t, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, ErrInvalidTTL
		}
		ttl = t
		typ = args[2]
		value = args[3]
	} else {
		return nil, ErrCanNotParse
	}

	// TODO name validation???
	if !isValudTTL(ttl) {
		return nil, ErrInvalidTTL
	}
	if !IsSupportedType(typ) {
		return nil, ErrInvalidType
	}
	return &Record{
		Name:   name,
		TTL:    ttl,
		Type:   typ,
		Values: []string{value},
	}, nil
}

func ParseRecordWithoutValue(slices []string) (*Record, error) {
	var name, typ string

	if len(slices) != 2 {
		return nil, ErrCanNotParse
	}

	name = slices[0]
	typ = slices[1]

	// TODO name validation???
	if !IsSupportedType(typ) {
		return nil, ErrInvalidType
	}
	return &Record{
		Name: name,
		Type: typ,
	}, nil
}
