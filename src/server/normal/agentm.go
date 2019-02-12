package normal

import (
	"sync"

	"github.com/name5566/leaf/gate"
	llog "github.com/name5566/leaf/log"
)

// AgentMap 连接管理器
type AgentMap struct {
	agents sync.Map
}

// NewAgentMap 构造一个AgentMap
func NewAgentMap() *AgentMap {
	return &AgentMap{}
}

// Add 添加数据(uid-agent)
func (a *AgentMap) Add(key int, agent *gate.Agent) {
	if _, ok := a.agents.LoadOrStore(key, agent); ok {
		llog.Release("key=%d already added.", key)
	}
}

// Remove 根据uid移除对应的agent
func (a *AgentMap) Remove(key int) {
	a.agents.Delete(key)
}

// Get 根据uid获取agent，如果存在ok=true
func (a *AgentMap) Get(key int) (agent *gate.Agent, ok bool) {
	tmp, is := a.agents.Load(key)
	if is {
		agent = tmp.(*gate.Agent)
		ok = is
	}
	return
}
