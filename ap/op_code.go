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

	//user
	OP_host_getUserTask
	OP_host_bindThreadID
	OP_host_userTaskCreate

	OP_agent_callAgent
	OP_agent_output
)

var OpCodeName = map[OpCode]string{
	OP_host_started:        "OP_host_started",
	OP_host_notify:         "OP_host_notify",
	OP_host_callAgent:      "OP_host_callAgent",
	OP_host_agentTask:      "OP_host_agentTask",
	OP_host_getUserTask:    "OP_host_getUserTask",
	OP_host_bindThreadID:   "OP_host_bindThreadID",
	OP_host_userTaskCreate: "OP_host_userTaskCreate",
	OP_agent_callAgent:     "OP_agent_callAgent",
	OP_agent_output:        "OP_agent_output",
}

var OpCodeMap = map[string]OpCode{
	"OP_host_started":   OP_host_started,
	"OP_host_notify":    OP_host_notify,
	"OP_host_callAgent": OP_host_callAgent,
	"OP_host_agentTask": OP_host_agentTask,

	"OP_host_getUserTask":    OP_host_getUserTask,
	"OP_host_bindThreadID":   OP_host_bindThreadID,
	"OP_host_userTaskCreate": OP_host_userTaskCreate,
	"OP_agent_callAgent":     OP_agent_callAgent,
	"OP_agent_output":        OP_agent_output,
}

type ThreadMemory struct {
	ThreadID string `json:"threadID" bson:"threadID"`
	Name     string `json:"name" bson:"name"`
	// History      []map[string]any `json:"history" bson:"history"`
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
