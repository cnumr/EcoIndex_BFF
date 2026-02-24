package services

import (
	"encoding/json"
	"log"

	_ "embed"
)

//go:embed ecoindex_reference.json
var ecoindexReferenceJSON []byte

type ecoindexReference struct {
	Grades []struct {
		Grade string `json:"grade"`
		Color string `json:"color"`
	} `json:"grades"`
}

var gradeColorMap map[string]string

func init() {
	var ref ecoindexReference

	if err := json.Unmarshal(ecoindexReferenceJSON, &ref); err != nil {
		log.Fatalf("failed to parse ecoindex_reference.json: %v", err)
	}

	gradeColorMap = make(map[string]string, len(ref.Grades))
	for _, g := range ref.Grades {
		gradeColorMap[g.Grade] = g.Color
	}
}

