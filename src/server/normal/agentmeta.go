package normal

// AgentMeta agent数据
type AgentMeta struct {
	// UID agent对应Uid
	UID int
}

// NewAgentMeta AgentMeta构造函数
func NewAgentMeta() *AgentMeta {
	return &AgentMeta{
		UID: 0,
	}
}
