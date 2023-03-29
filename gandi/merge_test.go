package gandi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	slices := []string{
		"www 3600 A 192.0.2.1",
		"www 3600 A 192.0.2.2",
		"www2 A 192.0.2.3",
		"www2 TXT aaaa",
		"www2 A 192.0.2.4",
	}

	var m RecordMerger
	for i := range slices {
		v, _ := ParseRecord(slices[i])
		err := m.Merge(*v)
		if err != nil {
			t.Errorf("invalid error: %v", err)
		}
	}

	actual := m.Output()
	expected := []Record{
		{
			Name: "www",
			TTL:  3600,
			Type: "A",
			Values: []string{
				"192.0.2.1",
				"192.0.2.2",
			},
		},
		{
			Name: "www2",
			TTL:  300,
			Type: "A",
			Values: []string{
				"192.0.2.3",
				"192.0.2.4",
			},
		},
		{
			Name:   "www2",
			TTL:    300,
			Type:   "TXT",
			Values: []string{"aaaa"},
		},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Output() is mismatch (-actual +expected):\n%s", diff)
	}

	rractual := m.ReplaceOutput()
	rrexpected := []ReplaceRecord{
		{
			ZoneName: "www",
			Items: []Record{
				{
					TTL:    3600,
					Type:   "A",
					Values: []string{"192.0.2.1", "192.0.2.2"},
				},
			},
		},
		{
			ZoneName: "www2",
			Items: []Record{
				{
					TTL:    300,
					Type:   "A",
					Values: []string{"192.0.2.3", "192.0.2.4"},
				},
				{
					TTL:    300,
					Type:   "TXT",
					Values: []string{"aaaa"},
				},
			},
		},
	}

	if diff := cmp.Diff(rractual, rrexpected); diff != "" {
		t.Errorf("ReplaceOutput() is mismatch (-actual +expected):\n%s", diff)
	}

}
