package lsp

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"unicode/utf16"
)

type hoverRequest struct {
	baseRequestMessage
	Params hoverRequestParams `json:"params"`
}

type hoverRequestParams struct {
	TextDocument textDocumentIdentifier `json:"textDocument"`
	Position     position               `json:"position"`
}

type hoverResponse struct {
	baseResponseMessage
	Result *hoverResult `json:"result"`
}

type hoverResult struct {
	Contents markupContent `json:"contents"`
	Range    positionRange `json:"range"`
}

func processHoverRequest(content []byte, state *State) []byte {
	var request hoverRequest
	if err := json.Unmarshal(content, &request); err != nil {
		result, _ := json.Marshal(responseMessageError{
			baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID},
			Error:               responseError{Code: jsonParseError, Message: fmt.Sprintf("JSON Parse Error: %v", err.Error())},
		})
		return result
	}

	text := state.read(request.Params.TextDocument.URI)
	lines := strings.Split(text, "\n")
	if request.Params.Position.Line >= len(lines) {
		result, _ := json.Marshal(hoverResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}

	characters := utf16.Encode([]rune(lines[request.Params.Position.Line]))
	if request.Params.Position.Character >= len(characters) {
		result, _ := json.Marshal(hoverResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}
	if character := characters[request.Params.Position.Character]; slices.Contains([]uint16{' ', '\t', '=', '>'}, character) {
		result, _ := json.Marshal(hoverResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}

	leftIdx := request.Params.Position.Character
	for ; leftIdx >= 0 && !slices.Contains([]uint16{' ', '\t'}, characters[leftIdx]); leftIdx-- {
	}
	rightIdx := request.Params.Position.Character
	for ; rightIdx < len(characters) && !slices.Contains([]uint16{' ', '\t', '=', '>'}, characters[rightIdx]); rightIdx++ {
	}
	hoveredWord := string(utf16.Decode(characters[leftIdx+1 : rightIdx]))

	if _, ok := inHTMLElementTag(lines[:request.Params.Position.Line+1], request.Params.Position); !ok {
		result, _ := json.Marshal(hoverResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}

	attribute, ok := attributes[hoveredWord]
	if !ok {
		result, _ := json.Marshal(hoverResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}
	documentation := attribute.documentation
	if documentation == "" {
		documentation = attribute.detail
	}
	result, _ := json.Marshal(hoverResponse{
		baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID},
		Result: &hoverResult{
			Contents: markupContent{Kind: markdownMarkupKind, Value: documentation},
			Range: positionRange{
				Start: position{Line: request.Params.Position.Line, Character: leftIdx + 1},
				End:   position{Line: request.Params.Position.Line, Character: rightIdx},
			},
		},
	})
	return result
}
