package ap

import (
	"encoding/json"
	"fmt"
)

type OpCode int32

const (
	//system
	OP_host_started OpCode = iota
	OP_host_notify
	OP_host_agentTask
	OP_host_callAgent
	OP_host_agentList

	//user
	OP_host_getUserTask
	OP_host_bindThreadID

	//task
	OP_host_userTaskCreate
	OP_host_userTaskList

	//thread
	OP_host_threadHistory

	OP_agent_callAgent
	OP_agent_output
)

var OpCodeName = map[OpCode]string{
	OP_host_started:        "OP_host_started",
	OP_host_notify:         "OP_host_notify",
	OP_host_callAgent:      "OP_host_callAgent",
	OP_host_agentTask:      "OP_host_agentTask",
	OP_host_agentList:      "OP_host_agentList",
	OP_host_threadHistory:  "OP_host_threadHistory",
	OP_host_getUserTask:    "OP_host_getUserTask",
	OP_host_bindThreadID:   "OP_host_bindThreadID",
	OP_host_userTaskCreate: "OP_host_userTaskCreate",
	OP_host_userTaskList:   "OP_host_userTaskList",
	OP_agent_callAgent:     "OP_agent_callAgent",
	OP_agent_output:        "OP_agent_output",
}

var OpCodeMap = map[string]OpCode{
	"OP_host_started":       OP_host_started,
	"OP_host_notify":        OP_host_notify,
	"OP_host_callAgent":     OP_host_callAgent,
	"OP_host_agentTask":     OP_host_agentTask,
	"OP_host_agentList":     OP_host_agentList,
	"OP_host_threadHistory": OP_host_threadHistory,

	"OP_host_getUserTask":    OP_host_getUserTask,
	"OP_host_bindThreadID":   OP_host_bindThreadID,
	"OP_host_userTaskCreate": OP_host_userTaskCreate,
	"OP_host_userTaskList":   OP_host_userTaskList,
	"OP_agent_callAgent":     OP_agent_callAgent,
	"OP_agent_output":        OP_agent_output,
}

type NotifyParams struct {
	ThreadID  string `json:"threadID"`
	Role      string `json:"role"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	AgentName string `json:"agentName"`
	AgentID   string `json:"agentID"`
	Data      any    `json:"data"`
}
type ModelConfig struct {
	Name     string `json:"name"`
	Stream   bool   `json:"stream,omitempty"`
	Provider string `json:"provider,omitempty"`
	APIKey   string `json:"apiKey"`
	Thinking bool   `json:"thinking,omitempty"`
	Timeout  int64  `json:"timeout,omitempty"`
	URL      string `json:"url"`
}

type MCPType string

const (
	Stdio      MCPType = "stdio"
	HTTPStream MCPType = "httpstream"
)

type MCPServer struct {
	Name    string   `json:"name"`
	Type    MCPType  `json:"type"`
	URL     string   `json:"url"`
	Command []string `json:"command,omitempty"`
}
type AgentTools struct {
	AllowedMCPServers  []string `json:"allowedMCPServers"`
	ExcludedMCPServers []string `json:"excludedMCPServers"`
	AllowedTools       []string `json:"allowedTools"`
	ExcludedTools      []string `json:"excludedTools"`
}

type AgentConfig struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Bio         string       `json:"bio"`
	Sysprompt   string       `json:"systemPrompt"`
	Tags        []string     `json:"tags"`
	ModelConfig *ModelConfig `json:"model"`

	ConnType   string       `json:"connType"`
	Command    []string     `json:"command"`
	URL        string       `json:"url"`
	MCPServers []*MCPServer `json:"mcpServers"`
	Tools      *AgentTools  `json:"tools"`
	OpCodes    []OpCode     `json:"opCodes"`
}

type UserTask struct {
	UserID    string   `bson:"userID" json:"userID"`
	TaskID    string   `bson:"taskID" json:"taskID"`
	TaskName  string   `bson:"taskName" json:"taskName"`
	ThreadIDs []string `bson:"threadIDs" json:"threadIDs"`
}
type ThreadMemory struct {
	ThreadID     string   `json:"threadID" bson:"threadID"`
	Name         string   `json:"name" bson:"name"`
	AgentTaskIDs []string `json:"agentTaskIDs" bson:"agentTaskIDs"`
	CreatedAt    int64    `json:"createdAt" bson:"createdAt"`
	UpdatedAt    int64    `json:"updatedAt" bson:"updatedAt"`
}

func (c *OpCode) String() string {
	return OpCodeName[*c]
}

// UnmarshalJSON implements custom JSON unmarshaling for OpCode
func (c *OpCode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		// Try to unmarshal as int32
		var i int32
		if err := json.Unmarshal(data, &i); err != nil {
			return fmt.Errorf("OpCode must be a string or int32: %w", err)
		}
		*c = OpCode(i)
		return nil
	}

	code, ok := OpCodeMap[s]
	if !ok {
		return fmt.Errorf("invalid OpCode value: %s", s)
	}
	*c = code
	return nil
}

// MarshalJSON implements custom JSON marshaling for OpCode
func (c OpCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// OP_agent_callAgent message
type OpAgentCallAgentMessage struct {
	AgentID string `json:"agentID"`
	Content string `json:"content"`
}

func (x *OpAgentCallAgentMessage) isOpMessage() {}

// OP_host_callAgent message
type OpHostCallAgentMessage struct {
	AgentID string `json:"agentID"`
	Content string `json:"content"`
}

func (x *OpHostCallAgentMessage) isOpMessage() {}

// OP_host_getUserTask message
type OpHostGetUserTaskMessage struct {
	TaskID string `json:"taskID"`
}
type OpHostGetUserTaskResp struct {
	TaskID   string          `json:"taskID"`
	TaskName string          `json:"taskName"`
	Threads  []*ThreadMemory `json:"threads"`
}

func (x *OpHostGetUserTaskMessage) isOpMessage() {}

// OP_host_bindThreadID message
type OpHostBindThreadIDMessage struct {
	TaskID   string `json:"taskID"`
	ThreadID string `json:"threadID"`
}

func (x *OpHostBindThreadIDMessage) isOpMessage() {}

type OpHostUserTaskCreateMessage struct {
	UserID   string `json:"userID"`
	TaskName string `json:"taskName"`
}

func (x *OpHostUserTaskCreateMessage) isOpMessage() {}

type OpHostUserTaskListMessage struct {
	UserID string `json:"userID"`
}

func (x *OpHostUserTaskListMessage) isOpMessage() {}

type OpHostThreadHistoryMessage struct {
	ThreadID string `json:"threadID"`
}

func (x *OpHostThreadHistoryMessage) isOpMessage() {}
