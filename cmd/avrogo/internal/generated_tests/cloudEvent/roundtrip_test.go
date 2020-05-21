// Code generated by generatetestcode.go; DO NOT EDIT.

package cloudEvent

import (
	"testing"

	"github.com/heetch/avro/cmd/avrogo/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "com.heetch.someDomain.someEvent",
                "type": "record",
                "fields": [
                    {
                        "name": "Metadata",
                        "type": {
                            "name": "Metadata",
                            "type": "record",
                            "fields": [
                                {
                                    "name": "CloudEvent",
                                    "type": {
                                        "name": "CloudEvent",
                                        "type": "record",
                                        "fields": [
                                            {
                                                "name": "id",
                                                "type": "string"
                                            },
                                            {
                                                "name": "source",
                                                "type": "string",
                                                "doc": "* source holds the\n\t\t * source of the message."
                                            },
                                            {
                                                "name": "specversion",
                                                "type": "string"
                                            },
                                            {
                                                "name": "time",
                                                "type": {
                                                    "type": "long",
                                                    "logicalType": "timestamp-micros"
                                                }
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    },
                    {
                        "name": "other",
                        "type": "string",
                        "doc": "other documentation"
                    }
                ],
                "heetchmeta": {
                    "status": "active",
                    "commentary": "This Schema describes version v9.9.99 of the event someEvent from the domain someDomain.",
                    "topickey": "someDomain.someEvent.v9.9.99",
                    "partitions": 1
                }
            }`,
	GoType: new(Message),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "Metadata": {
                            "CloudEvent": {
                                "time": 1580392724000000,
                                "id": "id1",
                                "source": "source1",
                                "specversion": "someversion"
                            }
                        },
                        "other": "some other data"
                    }`,
		OutDataJSON: `{
                        "Metadata": {
                            "CloudEvent": {
                                "time": 1580392724000000,
                                "id": "id1",
                                "source": "source1",
                                "specversion": "someversion"
                            }
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
