package message

const (
	PanicLevel = "panic"
	FatalLevel = "fatal"
	ErrorLevel = "error_response"
	WarnLevel  = "warn"
	InfoLevel  = "info"
	DebugLevel = "debug"
	TraceLevel = "trace"
)

// LogMessage ...
type LogMessage struct {
	Message     string          `json:"message"`
	FullMessage *string         `json:"full_message,omitempty"`
	Host        *string         `json:"host,omitempty"`
	Timestamp   *float64        `json:"timestamp,omitempty"`
	Facility    *string         `json:"facility,omitempty"`
	Extra       *map[string]any `json:"extra,omitempty"`
}

// NewMessage ...
func NewMessage(message string) *LogMessage {
	return &LogMessage{
		Message: message,
	}
}
