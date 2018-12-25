package ppanalysis

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

type AnalysisBody = map[string]map[string][]AnalysisMessage

type AnalysisMessage struct {
	Posn    string `json:"posn"`
	Message string `json:"message"`
}

func Parse(r io.Reader) (AnalysisBody, error) {
	body := AnalysisBody{}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "parse: ioutil.ReadAll")
	}
	err = json.Unmarshal(b, &body)
	if err != nil {
		return nil, errors.Wrap(err, "parse: json.Unmarshal")
	}

	return body, nil
}
