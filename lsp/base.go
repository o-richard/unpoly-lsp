package lsp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"sync/atomic"
)

const (
	rpcVersion = "2.0"

	// JSON-RPC Error Codes (-32099 to -32000)
	jsonParseError           = -32700
	jsonInvalidRequest       = -32600
	jsonMethodNotFound       = -32601
	jsonInvalidParams        = -32602
	jsonInternalError        = -32603
	jsonServerNotInitialized = -32002
	jsonUnknownErrorCode     = -32001

	// LSP Error Codes (-32899 to -32800)
)

const (
	initializeMethod             = "initialize"
	shutdownMethod               = "shutdown"
	textDocumentCompletionMethod = "textDocument/completion"
	textDocumentHoverMethod      = "textDocument/hover"

	// Notifications
	initializedMethod        = "initialized"
	exitMethod               = "exit"
	textDocumentOpenMethod   = "textDocument/didOpen"
	textDocumentChangeMethod = "textDocument/didChange"
	textDocumentCloseMethod  = "textDocument/didClose"
)

type baseRequestMessage struct {
	RPC    string `json:"jsonrpc"`
	ID     *int   `json:"id"`
	Method string `json:"method"`
}

type baseResponseMessage struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id"`
}

type responseMessageError struct {
	baseResponseMessage
	Error responseError `json:"error"`
}

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type responseMessageSuccess struct {
	baseResponseMessage
	Result map[string]any `json:"result"`
}

const (
	uninitializeState int64 = iota
	initializedState
	shutdownState
)

var (
	applicationState atomic.Int64
)

func decodeRequestMessage(data []byte) (baseInfo baseRequestMessage, content []byte, jsonErr error) {
	header, body, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return baseRequestMessage{}, nil, fmt.Errorf("header & body separator not found")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil || contentLength < 0 || contentLength > len(body) {
		return baseRequestMessage{}, nil, fmt.Errorf("content length should be a valid positive integer")
	}

	var baseRequest baseRequestMessage
	if err := json.Unmarshal(body[:contentLength], &baseRequest); err != nil {
		return baseRequestMessage{}, nil, err
	}
	return baseRequest, body[:contentLength], nil
}

func checkApplicationState(id *int, method string) []byte {
	if applicationState.Load() == uninitializeState && !slices.Contains([]string{initializeMethod, exitMethod}, method) {
		result, _ := json.Marshal(responseMessageError{
			baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: id},
			Error:               responseError{Code: jsonServerNotInitialized, Message: "Server Not Initialized: Initialization request was not received."},
		})
		return result
	}
	if applicationState.Load() == shutdownState && method != exitMethod {
		result, _ := json.Marshal(responseMessageError{
			baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: id},
			Error:               responseError{Code: jsonInvalidRequest, Message: "Invalid Request: Shutdown notification was recieved."},
		})
		return result
	}
	return nil
}

// The caller should check whether the response is empty before writing to the sender.
func HandleRequestMessage(data []byte, state *State) (response string, shouldExit bool) {
	baseInfo, content, err := decodeRequestMessage(data)
	if err != nil {
		result, _ := json.Marshal(responseMessageError{
			baseResponseMessage: baseResponseMessage{RPC: rpcVersion},
			Error:               responseError{Code: jsonParseError, Message: fmt.Sprintf("JSON Parse Error: %v", err.Error())},
		})
		return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(result), result), false
	}

	if result := checkApplicationState(baseInfo.ID, baseInfo.Method); len(result) != 0 {
		return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(result), result), false
	}

	var result []byte
	switch baseInfo.Method {
	case initializeMethod:
		result = processInitializeRequest(baseInfo.ID)
	case textDocumentOpenMethod:
		processDidOpenTextDocumentNotification(content, state)
	case textDocumentChangeMethod:
		processDidChangeTextDocumentNotification(content, state)
	case textDocumentCloseMethod:
		processDidCloseTextDocumentNotification(content, state)
	case textDocumentHoverMethod:
		result = processHoverRequest(content, state)
	case textDocumentCompletionMethod:
		result = processCompletionRequest(content, state)
	case shutdownMethod:
		applicationState.Store(shutdownState)
		result, _ = json.Marshal(responseMessageSuccess{
			baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: baseInfo.ID},
		})
	case exitMethod:
		return "", true
	default:
		// Ignore notifications
		if baseInfo.ID != nil {
			result, _ = json.Marshal(responseMessageError{
				baseResponseMessage: baseResponseMessage{RPC: rpcVersion, ID: baseInfo.ID},
				Error:               responseError{Code: jsonMethodNotFound, Message: fmt.Sprintf("Unsupported Method: %v", baseInfo.Method)},
			})
		}
	}
	if len(result) == 0 {
		return "", false
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(result), result), false
}

func SplitFunc(data []byte, _ bool) (advance int, token []byte, err error) {
	header, body, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if contentLength < 0 {
		return 0, nil, fmt.Errorf("content length should be a valid positive integer")
	}

	if len(body) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
