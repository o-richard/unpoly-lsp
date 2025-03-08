package lsp_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/o-richard/unpoly-lsp/lsp"
)

type completionRequest struct {
	RPC    string                  `json:"jsonrpc"`
	ID     int                     `json:"id"`
	Method string                  `json:"method"`
	Params completionRequestParams `json:"params"`
}

type completionRequestParams struct {
	TextDocument textDocumentIdentifier `json:"textDocument"`
	Position     position               `json:"position"`
}

type completionResponse struct {
	RPC    string         `json:"jsonrpc"`
	ID     int            `json:"id"`
	Error  map[string]any `json:"error"`
	Result []any          `json:"result"`
}

func TestProcessCompletionRequest(t *testing.T) {
	testCases := []struct {
		Name          string
		Position      position
		ContainResult bool
	}{
		{Name: "Invalid Line Number", Position: position{Line: 10_000}},
		{Name: "Whitespace", Position: position{Line: 3, Character: 2}, ContainResult: true},
		{Name: "Valid Attribute Inside HTML Tag", Position: position{Line: 2, Character: 3}, ContainResult: true},
		{Name: "Invalid Attribute Inside HTML Tag", Position: position{Line: 3, Character: 29}},
		{Name: "Valid Attribute Outside HTML Tag", Position: position{Line: 5, Character: 8}},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			request := completionRequest{
				RPC: rpcVersion, Method: textDocumentCompletionMethod,
				Params: completionRequestParams{
					TextDocument: textDocumentIdentifier{URI: htmlFileURI}, Position: test.Position,
				},
			}
			requestBytes, _ := json.Marshal(request)
			response, _ := lsp.HandleRequestMessage(prepareRequest(requestBytes), applicationState)
			_, body, _ := bytes.Cut([]byte(response), []byte{'\r', '\n', '\r', '\n'})

			var result completionResponse
			if err := json.Unmarshal(body, &result); err != nil {
				t.Fatalf("JSON Unmarshall Error: %v", err.Error())
			}
			if result.Error != nil {
				t.Fatalf("Response Error: %v", result.Error)
			}

			if test.ContainResult && len(result.Result) == 0 {
				t.Fatalf("Expected a result, Got nothing")
			}
			if !test.ContainResult && len(result.Result) != 0 {
				t.Fatalf("Expected no result, Got a result")
			}
		})
	}
}
