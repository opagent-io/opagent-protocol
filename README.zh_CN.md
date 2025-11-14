# 简介

OpAgent Protocol 是一个分布式智能体互操作(operation)协议，在扩展的 MCP 规范之上构建，使各类智能体能够像插件或 Web 服务一样互相调用并共享能力。

# 规范说明

### agent

```
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "agent/call",
  "params": "
}
```