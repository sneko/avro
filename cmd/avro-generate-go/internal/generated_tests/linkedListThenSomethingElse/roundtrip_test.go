// Code generated by generatetestcode.go; DO NOT EDIT.

package linkedListThenSomethingElse

import (
	"testing"

	"github.com/heetch/avro/cmd/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "L",
                        "type": {
                            "name": "List",
                            "type": "record",
                            "fields": [
                                {
                                    "name": "Item",
                                    "type": "int"
                                },
                                {
                                    "name": "Next",
                                    "type": [
                                        "null",
                                        "List"
                                    ],
                                    "default": null
                                }
                            ]
                        }
                    },
                    {
                        "name": "M",
                        "type": "int"
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "M": 1234,
                        "L": {
                            "Item": 1234,
                            "Next": {
                                "List": {
                                    "Item": 9999,
                                    "Next": null
                                }
                            }
                        }
                    }`,
		OutDataJSON: `{
                        "M": 1234,
                        "L": {
                            "Item": 1234,
                            "Next": {
                                "List": {
                                    "Item": 9999,
                                    "Next": null
                                }
                            }
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}