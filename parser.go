package ppanalysis

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type AnalysisBody = map[string]map[string][]AnalysisMessage

type AnalysisMessage struct {
	Posn    string `json:"posn"`
	Message string `json:"message"`
}

func Parse(r io.Reader) (AnalysisBody, error) {
	var body AnalysisBody
	if err := json.NewDecoder(r).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "parse: json.Decoder.Decode")
	}
	return body, nil
}
