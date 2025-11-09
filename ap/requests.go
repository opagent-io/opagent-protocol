// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// This file holds the request types.

package ap

type (
	CallToolRequest                   = AgentRequest[*CallToolParamsRaw]
	OpAgentRequest                    = AgentRequest[*OpAgentParamsRaw]
	CompleteRequest                   = AgentRequest[*CompleteParams]
	GetPromptRequest                  = AgentRequest[*GetPromptParams]
	InitializedRequest                = AgentRequest[*InitializedParams]
	ListPromptsRequest                = AgentRequest[*ListPromptsParams]
	ListResourcesRequest              = AgentRequest[*ListResourcesParams]
	ListResourceTemplatesRequest      = AgentRequest[*ListResourceTemplatesParams]
	ListToolsRequest                  = AgentRequest[*ListToolsParams]
	ProgressNotificationServerRequest = AgentRequest[*ProgressNotificationParams]
	ReadResourceRequest               = AgentRequest[*ReadResourceParams]
	RootsListChangedRequest           = AgentRequest[*RootsListChangedParams]
	SubscribeRequest                  = AgentRequest[*SubscribeParams]
	UnsubscribeRequest                = AgentRequest[*UnsubscribeParams]
)

type (
	CreateMessageRequest = ClientRequest[*CreateMessageParams]
	OpHostRequest        = ClientRequest[*OpHostParams]
	// CallAgentRequest                   = ClientRequest[*CallAgentMessageParams]
	ElicitRequest                      = ClientRequest[*ElicitParams]
	initializedClientRequest           = ClientRequest[*InitializedParams]
	InitializeRequest                  = ClientRequest[*InitializeParams]
	ListRootsRequest                   = ClientRequest[*ListRootsParams]
	LoggingMessageRequest              = ClientRequest[*LoggingMessageParams]
	ProgressNotificationClientRequest  = ClientRequest[*ProgressNotificationParams]
	PromptListChangedRequest           = ClientRequest[*PromptListChangedParams]
	ResourceListChangedRequest         = ClientRequest[*ResourceListChangedParams]
	ResourceUpdatedNotificationRequest = ClientRequest[*ResourceUpdatedNotificationParams]
	ToolListChangedRequest             = ClientRequest[*ToolListChangedParams]
)
