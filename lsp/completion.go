package lsp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"unicode/utf16"
)

const (
	valueCompletionItemKind = 12
	unitCompletionItemKind  = 11

	deprecatedCompletionItemTag = 1

	plainTextInsertTextFormat = 1
	snippetInsertTextFormat   = 2

	asIsInsertTextMode              = 1
	adjustIndentationInsertTextMode = 2
)

type completionRequest struct {
	baseRequestMessage
	Params completionRequestParams `json:"params"`
}

type completionRequestParams struct {
	TextDocument textDocumentIdentifier `json:"textDocument"`
	Position     position               `json:"position"`
}

type completionResponse struct {
	baseResponseMessage
	Result []completionItem `json:"result"`
}

type completionItem struct {
	Label            string         `json:"label"`
	Kind             int            `json:"kind"`
	Tags             []int          `json:"tags,omitempty"`
	Detail           string         `json:"detail,omitempty"`
	Documentation    *markupContent `json:"documentation,omitempty"`
	InsertText       string         `json:"insertText"`
	InsertTextFormat int            `json:"insertTextFormat"`
	InsertTextMode   int            `json:"insertTextMode"`
}

func checkCompletionUnits(characters []uint16, characterIdx int) string {
	minLeftIdx := characterIdx - 4  // account for  x="
	minRightIdx := characterIdx + 1 // account for "
	if minLeftIdx < 0 || minRightIdx >= len(characters) {
		return ""
	}
	inQuotations := isAlphanumeric(characters[characterIdx-3]) &&
		characters[characterIdx-2] == '=' &&
		characters[characterIdx-1] == '"' &&
		characters[minRightIdx] == '"'
	if !inQuotations {
		return ""
	}
	for ; minLeftIdx >= 0 && !slices.Contains([]uint16{' ', '\t'}, characters[minLeftIdx]); minLeftIdx-- {
	}
	word := utf16.Decode(characters[minLeftIdx+1 : characterIdx-2])
	return string(word)
}

func checkCompletionValues(characters []uint16, characterIdx int) string {
	minLeftIdx := characterIdx - 1
	for ; minLeftIdx >= 0 && !slices.Contains([]uint16{' ', '\t'}, characters[minLeftIdx]); minLeftIdx-- {
	}
	word := utf16.Decode(characters[minLeftIdx+1 : characterIdx])
	return string(word)
}

func obtainCompletionUnits(word string, requestId *int) []byte {
	choices := attributes[word].choices
	items := make([]completionItem, 0, len(choices))
	for _, choice := range choices {
		items = append(items, completionItem{
			Label: choice, Kind: unitCompletionItemKind, InsertText: choice,
			InsertTextFormat: plainTextInsertTextFormat, InsertTextMode: asIsInsertTextMode,
		})
	}
	if len(items) == 0 {
		items = nil
	}
	result, _ := json.Marshal(completionResponse{
		baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: requestId},
		Result:              items,
	})
	return result
}

func shouldExcludeCompletionValue(text []byte, label string, lines []string, positionRange *positionRange) bool {
	var startIndex int
	for i := 0; i < positionRange.Start.Line; i++ {
		startIndex += len(lines[i]) + 1
	}
	characters := utf16.Encode([]rune(lines[positionRange.Start.Line]))
	word := string(utf16.Decode(characters[0 : positionRange.Start.Character+1]))
	startIndexIncrement := len(word)
	startIndex += startIndexIncrement

	var endIndex int
	for i := positionRange.Start.Line; i < positionRange.End.Line; i++ {
		endIndex += len(lines[i]) + 1
	}
	characters = utf16.Encode([]rune(lines[positionRange.End.Line]))
	word = string(utf16.Decode(characters[0 : positionRange.End.Character+1]))
	endIndex += (startIndex - startIndexIncrement) + len(word)

	return bytes.Contains(text[startIndex:endIndex], fmt.Appendf(nil, "%v", label))
}

func obtainCompletionValues(word string, requestId *int, text []byte, lines []string, positionRange *positionRange) []byte {
	items := make([]completionItem, 0)
	for label, attribute := range attributes {
		if !strings.HasPrefix(label, word) {
			continue
		}
		if shouldExcludeCompletionValue(text, label, lines, positionRange) {
			continue
		}
		var tags []int
		if attribute.isDepreciated {
			tags = []int{deprecatedCompletionItemTag}
		}
		items = append(items, completionItem{
			Label: label, Kind: valueCompletionItemKind, Tags: tags,
			Detail: attribute.detail, Documentation: &markupContent{Kind: markdownMarkupKind, Value: attribute.documentation},
			InsertText: attribute.insertText, InsertTextFormat: attribute.insertTextFormat, InsertTextMode: asIsInsertTextMode,
		})
	}
	if len(items) == 0 {
		items = nil
	}
	result, _ := json.Marshal(completionResponse{
		baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: requestId},
		Result:              items,
	})
	return result
}

func processCompletionRequest(content []byte, state *State) []byte {
	var request completionRequest
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
		result, _ := json.Marshal(completionResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}

	characters := utf16.Encode([]rune(lines[request.Params.Position.Line]))
	request.Params.Position.Character--
	if request.Params.Position.Character >= len(characters) {
		result, _ := json.Marshal(completionResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}

	positionRange, ok := inHTMLElementTag(lines[:request.Params.Position.Line+1], request.Params.Position)
	if !ok {
		result, _ := json.Marshal(completionResponse{baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: request.ID}})
		return result
	}

	attribute := checkCompletionUnits(characters, request.Params.Position.Character)
	if attribute != "" {
		return obtainCompletionUnits(attribute, request.ID)
	}

	label := checkCompletionValues(characters, request.Params.Position.Character)
	return obtainCompletionValues(label, request.ID, []byte(text), lines[:request.Params.Position.Line+1], &positionRange)
}
