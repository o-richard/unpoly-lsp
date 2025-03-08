package lsp

import (
	"encoding/json"
	"unicode/utf16"
)

const (
	plaintextMarkupKind = "plaintext"
	markdownMarkupKind  = "markdown"
)

type textDocumentItem struct {
	Version int    `json:"version"`
	URI     string `json:"uri"`
	Text    string `json:"text"`
}

type textDocumentIdentifier struct {
	URI string `json:"uri"`
}

type versionedTextDocumentIdentifier struct {
	URI     string `json:"uri"`
	Version int    `json:"version"`
}

type position struct {
	Line      int `json:"line"`      // zero-based. utf-16.
	Character int `json:"character"` // zero-based. utf-16. if the character value is greater than the line length it defaults back to the line length.
}

type positionRange struct {
	Start position `json:"start"`
	End   position `json:"end"` // exclusive
}

type markupContent struct {
	Kind  string `json:"kind"` // plaintext. markdown.
	Value string `json:"value"`
}

type didOpenTextDocumentNotification struct {
	baseRequestMessage
	Params didOpenTextDocumentParams `json:"params"`
}

type didOpenTextDocumentParams struct {
	TextDocument textDocumentItem `json:"textDocument"`
}

func processDidOpenTextDocumentNotification(content []byte, state *State) {
	var notification didOpenTextDocumentNotification
	if err := json.Unmarshal(content, &notification); err != nil {
		return
	}
	state.update(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text, notification.Params.TextDocument.Version)
}

type didChangeTextDocumentNotification struct {
	baseRequestMessage
	Params didChangeTextDocumentParams `json:"params"`
}

type didChangeTextDocumentParams struct {
	TextDocument   versionedTextDocumentIdentifier       `json:"textDocument"`
	ContentChanges []didChangeTextDocumentContentChanges `json:"contentChanges"`
}

type didChangeTextDocumentContentChanges struct {
	Text string `json:"text"`
}

func processDidChangeTextDocumentNotification(content []byte, state *State) {
	var notification didChangeTextDocumentNotification
	if err := json.Unmarshal(content, &notification); err != nil {
		return
	}
	text := notification.Params.ContentChanges[len(notification.Params.ContentChanges)-1].Text
	state.update(notification.Params.TextDocument.URI, text, notification.Params.TextDocument.Version)
}

type didCloseTextDocumentNotification struct {
	baseRequestMessage
	Params didCloseTextDocumentParams `json:"params"`
}

type didCloseTextDocumentParams struct {
	TextDocument textDocumentIdentifier `json:"textDocument"`
}

func processDidCloseTextDocumentNotification(content []byte, state *State) {
	var notification didCloseTextDocumentNotification
	if err := json.Unmarshal(content, &notification); err != nil {
		return
	}
	state.close(notification.Params.TextDocument.URI)
}

func isAlphanumeric(character uint16) bool {
	return (character >= '0' && character <= '9') || (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')
}

func inHTMLElementTag(lines []string, pos position) (positionRange, bool) {
	var quotes int
	for lineIdx := len(lines) - 1; lineIdx >= 0; lineIdx-- {
		characters := utf16.Encode([]rune(lines[lineIdx]))

		characterIdx := len(characters) - 1
		if pos.Line == lineIdx {
			characterIdx = pos.Character
		}

		for ; characterIdx >= 0; characterIdx-- {
			if characters[characterIdx] == '\'' || characters[characterIdx] == '"' {
				quotes++
			}
			if characters[characterIdx] == '<' || characters[characterIdx] == '>' {
				nextIdx := characterIdx + 1
				isElementTag := characters[characterIdx] == '<' && nextIdx < len(characters) && isAlphanumeric(characters[nextIdx]) && quotes%2 == 0
				if isElementTag {
					return positionRange{
						Start: position{Line: lineIdx, Character: characterIdx}, End: pos,
					}, true
				}
				return positionRange{}, false
			}
		}
	}
	return positionRange{}, false
}
