package main

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
)

/*
[
    {
        "name": "supported_file_type",
        "help": "",
        "type": 1,
        "metric": [
            {
                "label": [
                    {
                        "name": "file_ext",
                        "value": "docx"
                    }
                ],
                "gauge": {
                    "value": 1
                }
            },
            {
                "label": [
                    {
                        "name": "file_ext",
                        "value": "txt"
                    }
                ],
                "gauge": {
                    "value": 5
                }
            }
        ]
    }
]
*/

func dumpToJson(gather prometheus.Gatherer) (string, error) {
	mfs, err := gather.Gather()
	if err != nil {
		panic(err)
	}
	b, e := json.MarshalIndent(mfs, "", "    ")
	return string(b), e
}
