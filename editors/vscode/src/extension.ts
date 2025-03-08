// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';
import { LanguageClient, LanguageClientOptions, ServerOptions } from 'vscode-languageclient/node';

let lc: LanguageClient;

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
	const serverOptions: ServerOptions = {
		command: 'unpoly-lsp',
	};
	const clientOptions: LanguageClientOptions = {
		documentSelector: [
			{language: 'html'},
			{language: "templ"},
			{language: "erb"},
		],
	};
	lc = new LanguageClient('unpoly-lsp', 'Unpoly LSP', serverOptions, clientOptions);
	lc.start();
}

// This method is called when your extension is deactivated
export function deactivate() {
	if (!lc) {
		return undefined;
	}
	return lc.stop();
}
