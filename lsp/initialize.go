package lsp

import (
	"encoding/json"
)

type initializeResponse struct {
	baseResponseMessage
	Result initializeResult `json:"result"`
}

type initializeResult struct {
	ServerInfo   initializeResultServerInfo         `json:"serverInfo"`
	Capabilities initializeResultServerCapabilities `json:"capabilities"`
}

type initializeResultServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type initializeResultServerCapabilities struct {
	TextDocumentSync   int            `json:"textDocumentSync"`
	HoverProvider      bool           `json:"hoverProvider"`
	CompletionProvider map[string]any `json:"completionProvider"`
}

func processInitializeRequest(id *int) []byte {
	applicationState.Store(initializedState)
	response := initializeResponse{
		baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: id},
		Result: initializeResult{
			Capabilities: initializeResultServerCapabilities{
				TextDocumentSync: 1, HoverProvider: true, CompletionProvider: map[string]any{},
			},
			ServerInfo: initializeResultServerInfo{Name: "unpoly-lsp", Version: "0.1.0"},
		},
	}
	data, _ := json.Marshal(response)
	return data
}
