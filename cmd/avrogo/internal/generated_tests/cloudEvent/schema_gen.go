// Code generated by avrogen. DO NOT EDIT.

package cloudEvent

import (
	"github.com/heetch/avro/avrotypegen"
	"time"
)

type CloudEvent struct {
	Id string `json:"id"`

	// source holds the
	// source of the message.
	Source      string    `json:"source"`
	Specversion string    `json:"specversion"`
	Time        time.Time `json:"time"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (CloudEvent) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"id","type":"string"},{"doc":"* source holds the\n\t\t * source of the message.","name":"source","type":"string"},{"name":"specversion","type":"string"},{"name":"time","type":{"logicalType":"timestamp-micros","type":"long"}}],"name":"com.heetch.CloudEvent","type":"record"}`,
		Required: []bool{
			0: true,
			1: true,
			2: true,
			3: true,
		},
	}
}

type Message struct {
	Metadata Metadata
}

// AvroRecord implements the avro.AvroRecord interface.
func (Message) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"Metadata","type":{"fields":[{"name":"CloudEvent","type":{"fields":[{"name":"id","type":"string"},{"doc":"* source holds the\n\t\t * source of the message.","name":"source","type":"string"},{"name":"specversion","type":"string"},{"name":"time","type":{"logicalType":"timestamp-micros","type":"long"}}],"name":"CloudEvent","type":"record"}}],"name":"Metadata","type":"record"}}],"name":"com.heetch.Message","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}

type Metadata struct {
	CloudEvent CloudEvent
}

// AvroRecord implements the avro.AvroRecord interface.
func (Metadata) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"CloudEvent","type":{"fields":[{"name":"id","type":"string"},{"doc":"* source holds the\n\t\t * source of the message.","name":"source","type":"string"},{"name":"specversion","type":"string"},{"name":"time","type":{"logicalType":"timestamp-micros","type":"long"}}],"name":"CloudEvent","type":"record"}}],"name":"com.heetch.Metadata","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}
