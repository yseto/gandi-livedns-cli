package gandi

type ReplaceRecord struct {
	ZoneName string   `json:"-"`
	Items    []Record `json:"items"`
}

type RecordMerger struct {
	Records []Record
}

func (c *RecordMerger) Merge(r Record) error {
	for i := range c.Records {
		cmp := c.Records[i]
		if cmp.Name == r.Name && cmp.Type == r.Type {
			if cmp.TTL != r.TTL {
				return ErrDifferentTTL
			}
			c.Records[i].Values = append(c.Records[i].Values, r.Values...)
			return nil
		}
	}
	c.Records = append(c.Records, r)
	return nil
}

func (c *RecordMerger) Output() []Record {
	return c.Records
}

func (c *RecordMerger) ReplaceOutput() []ReplaceRecord {
	var rr []ReplaceRecord
	for i := range c.Records {
		alreadyAppend := false

		cmp := c.Records[i]
		for j := range rr {
			if rr[j].ZoneName == cmp.Name {
				appendR := Record{
					Type:   cmp.Type,
					Values: cmp.Values,
					TTL:    cmp.TTL,
				}
				rr[j].Items = append(rr[j].Items, appendR)
				alreadyAppend = true
			}
		}
		if alreadyAppend {
			continue
		}

		appendZ := ReplaceRecord{
			ZoneName: cmp.Name,
			Items: []Record{
				{
					Type:   cmp.Type,
					Values: cmp.Values,
					TTL:    cmp.TTL,
				},
			},
		}

		rr = append(rr, appendZ)
	}

	return rr
}
