package lsp_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/o-richard/unpoly-lsp/lsp"
)

type hoverRequest struct {
	RPC    string             `json:"jsonrpc"`
	ID     int                `json:"id"`
	Method string             `json:"method"`
	Params hoverRequestParams `json:"params"`
}

type hoverRequestParams struct {
	TextDocument textDocumentIdentifier `json:"textDocument"`
	Position     position               `json:"position"`
}

type hoverResponse struct {
	RPC    string         `json:"jsonrpc"`
	ID     int            `json:"id"`
	Error  map[string]any `json:"error"`
	Result *hoverResult   `json:"result"`
}

type hoverResult struct {
	Range positionRange `json:"range"`
}

func TestProcessHoverRequest(t *testing.T) {
	testCases := []struct {
		Name          string
		Position      position
		ContainResult bool
		ResultRange   positionRange
	}{
		{Name: "Invalid Line Number", Position: position{Line: 10_000}},
		{Name: "Whitespace", Position: position{Line: 3, Character: 2}},
		{
			Name:     "Valid Attribute Inside HTML Tag",
			Position: position{Line: 2, Character: 3}, ContainResult: true,
			ResultRange: positionRange{
				Start: position{Line: 2, Character: 1}, End: position{Line: 2, Character: 13},
			},
		},
		{Name: "Invalid Attribute Inside HTML Tag", Position: position{Line: 3, Character: 29}},
		{Name: "Valid Attribute Outside HTML Tag", Position: position{Line: 5, Character: 8}},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			request := hoverRequest{
				RPC: rpcVersion, Method: textDocumentHoverMethod,
				Params: hoverRequestParams{
					TextDocument: textDocumentIdentifier{URI: htmlFileURI}, Position: test.Position,
				},
			}
			requestBytes, _ := json.Marshal(request)
			response, _ := lsp.HandleRequestMessage(prepareRequest(requestBytes), applicationState)
			_, body, _ := bytes.Cut([]byte(response), []byte{'\r', '\n', '\r', '\n'})

			var result hoverResponse
			if err := json.Unmarshal(body, &result); err != nil {
				t.Fatalf("JSON Unmarshall Error: %v", err.Error())
			}
			if result.Error != nil {
				t.Fatalf("Response Error: %v", result.Error)
			}

			if test.ContainResult && result.Result == nil {
				t.Fatalf("Expected a result, Got nothing")
			}
			if !test.ContainResult && result.Result != nil {
				t.Fatalf("Expected no result, Got a result")
			}

			if test.ContainResult && test.ResultRange != result.Result.Range {
				t.Errorf("Expected %v, Got %v", test.ResultRange, result.Result.Range)
			}
		})
	}
}
