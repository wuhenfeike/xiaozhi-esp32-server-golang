package msg

const (
	MDeviceMockPubTopicPrefix = "device-server"
	MDeviceMockSubTopicPrefix = "null"
	MDeviceSubTopicPrefix     = "/p2p/device_sub/"
	MDevicePubTopicPrefix     = "/p2p/device_public/"
	MServerSubTopicPrefix     = "/p2p/device_public/#"
	MServerPubTopicPrefix     = MDeviceSubTopicPrefix
)

// 消息类型常量
const (
	MessageTypeHello  = "hello"  // 握手消息
	MessageTypeAbort  = "abort"  // 中止消息
	MessageTypeListen = "listen" // 监听消息
	MessageTypeIot    = "iot"    // 物联网消息
	MessageTypeMcp    = "mcp"    // MCP消息
)

// 服务器消息类型常量
const (
	ServerMessageTypeHello = "hello" // 握手消息
	ServerMessageTypeStt   = "stt"   // 语音转文本
	ServerMessageTypeTts   = "tts"   // 文本转语音
	ServerMessageTypeIot   = "iot"   // 物联网消息
	ServerMessageTypeLlm   = "llm"   // 大语言模型
	ServerMessageTypeText  = "text"  // 文本消息
)

// 消息状态常量
const (
	MessageStateStart         = "start"          // 开始状态
	MessageStateSentenceStart = "sentence_start" // 句子开始状态
	MessageStateSentenceEnd   = "sentence_end"   // 句子结束状态
	MessageStateStop          = "stop"           // 停止状态
	MessageStateDetect        = "detect"         // 检测状态
	MessageStateAbort         = "abort"          // 中止状态
	MessageStateSuccess       = "success"        // 成功状态
)
