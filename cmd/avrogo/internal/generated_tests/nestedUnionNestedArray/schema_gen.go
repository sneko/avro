// Code generated by avrogen. DO NOT EDIT.

package nestedUnionNestedArray

import (
	"github.com/heetch/avro/avrotypegen"
)

type R struct {
	F [][]*string
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"F","type":{"items":{"items":["null","string"],"type":"array"},"type":"array"}}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
		},
		Unions: []avrotypegen.UnionInfo{
			0: {
				Type: new([][]*string),
				Union: []avrotypegen.UnionInfo{{
					Type: nil,
				}, {
					Type: new(string),
				}},
			},
		},
	}
}
