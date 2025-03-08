package lsp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/o-richard/unpoly-lsp/lsp"
)

const (
	rpcVersion = "2.0"

	// Requests
	initializeMethod             = "initialize"
	shutdownMethod               = "shutdown"
	textDocumentCompletionMethod = "textDocument/completion"
	textDocumentHoverMethod      = "textDocument/hover"

	// Notifications
	exitMethod             = "exit"
	textDocumentOpenMethod = "textDocument/didOpen"

	htmlFileURI = "file:///home/test/index.html"
)

type baseRequest struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type textDocumentIdentifier struct {
	URI string `json:"uri"`
}

type position struct {
	Line      int `json:"line"`      // zero-based. utf-16.
	Character int `json:"character"` // zero-based. utf-16. if the character value is greater than the line length it defaults back to the line length.
}

type positionRange struct {
	Start position `json:"start"`
	End   position `json:"end"` // exclusive
}

type didOpenTextDocumentNotification struct {
	RPC    string                    `json:"jsonrpc"`
	Method string                    `json:"method"`
	Params didOpenTextDocumentParams `json:"params"`
}

type didOpenTextDocumentParams struct {
	TextDocument textDocumentItem `json:"textDocument"`
}

type textDocumentItem struct {
	Version int    `json:"version"`
	URI     string `json:"uri"`
	Text    string `json:"text"`
}

type exitNotification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

var applicationState *lsp.State

var documents map[string]string = map[string]string{
	/*
		HOVER TESTS

		Line 2, Character 3 - Valid Attribute Inside HTML Tag. Ranges from 1 - 12
		Line 3, Character 2 - Space
		Line 3, Character 29 - Invalid Attribute Inside HTML Tag
		Line 5, Character 8 - Valid Attribute Outside HTML Tag

		COMPLETION TESTS


	*/
	htmlFileURI: `
<span
	up-clickable role="button"
    up-background="true" hidden
	aria-haspopup="true"
>words up-clickable words</span>
`,
}

func prepareRequest(data []byte) []byte {
	return fmt.Appendf(nil, "Content-Length: %d\r\n\r\n%s", len(data), data)
}

func checkResponseError(data []byte) error {
	_, body, _ := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		return fmt.Errorf("JSON Unmarshall Error: %w", err)
	}
	if payload["error"] != nil {
		return fmt.Errorf("Response Error: %v", payload["error"])
	}
	return nil
}

func TestMain(m *testing.M) {
	applicationState = lsp.NewState()

	initRequest := baseRequest{RPC: rpcVersion, Method: initializeMethod}
	initRequestBytes, _ := json.Marshal(initRequest)
	response, _ := lsp.HandleRequestMessage(prepareRequest(initRequestBytes), applicationState)
	if err := checkResponseError([]byte(response)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for uri, content := range documents {
		openNotification := didOpenTextDocumentNotification{
			RPC: rpcVersion, Method: textDocumentOpenMethod,
			Params: didOpenTextDocumentParams{
				TextDocument: textDocumentItem{URI: uri, Version: 1, Text: content},
			},
		}
		openNotificationBytes, _ := json.Marshal(openNotification)
		response, _ := lsp.HandleRequestMessage(prepareRequest(openNotificationBytes), applicationState)
		if response != "" {
			fmt.Printf("Expected: empty response, Got: %v\n", response)
			os.Exit(1)
		}
	}

	exitCode := m.Run()

	shutdownRequest := baseRequest{RPC: rpcVersion, Method: shutdownMethod}
	shutdownRequestBytes, _ := json.Marshal(shutdownRequest)
	response, _ = lsp.HandleRequestMessage(prepareRequest(shutdownRequestBytes), applicationState)
	if err := checkResponseError([]byte(response)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	exitNotification := exitNotification{RPC: rpcVersion, Method: exitMethod}
	exitNotificationBytes, _ := json.Marshal(exitNotification)
	_, shouldExit := lsp.HandleRequestMessage(prepareRequest(exitNotificationBytes), applicationState)
	if !shouldExit {
		fmt.Println("Expected `ShoulExit` to be true")
		os.Exit(1)
	}

	os.Exit(exitCode)
}
