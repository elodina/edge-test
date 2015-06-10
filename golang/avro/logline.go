package avro

import "github.com/stealthly/go-avro"

type LogLine struct {
	Line      interface{}       `avro:"line"`
	Source    interface{}       `avro:"source"`
	Tag       map[string]string `avro:"tag"`
	Logtypeid interface{}       `avro:"logtypeid"`
	Timings   []*Timing         `avro:"timings"`
	Size      interface{}       `avro:"size"`
}

func NewLogLine() *LogLine {
	return &LogLine{}
}

func (this *LogLine) Schema() avro.Schema {
	if _LogLine_schema_err != nil {
		panic(_LogLine_schema_err)
	}
	return _LogLine_schema
}

type Timing struct {
	EventName string      `avro:"eventName"`
	Value     int64       `avro:"value"`
	Ntpstatus interface{} `avro:"ntpstatus"`
}

func NewTiming() *Timing {
	return &Timing{}
}

func (this *Timing) Schema() avro.Schema {
	if _Timing_schema_err != nil {
		panic(_Timing_schema_err)
	}
	return _Timing_schema
}

// Generated by codegen. Please do not modify.
var _LogLine_schema, _LogLine_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "avro",
    "name": "logLine",
    "fields": [
        {
            "name": "line",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "source",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "tag",
            "default": null,
            "type": [
                "null",
                {
                    "type": "map",
                    "values": "string"
                }
            ]
        },
        {
            "name": "logtypeid",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        },
        {
            "name": "timings",
            "default": null,
            "type": [
                "null",
                {
                    "type": "array",
                    "items": {
                        "type": "record",
                        "name": "Timing",
                        "fields": [
                            {
                                "name": "eventName",
                                "type": "string"
                            },
                            {
                                "name": "value",
                                "type": "long"
                            },
                            {
                                "name": "ntpstatus",
                                "default": null,
                                "type": [
                                    "null",
                                    "long"
                                ]
                            }
                        ]
                    }
                }
            ]
        },
        {
            "name": "size",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Timing_schema, _Timing_schema_err = avro.ParseSchema(`{
    "type": "record",
    "name": "Timing",
    "fields": [
        {
            "name": "eventName",
            "type": "string"
        },
        {
            "name": "value",
            "type": "long"
        },
        {
            "name": "ntpstatus",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        }
    ]
}`)
