package logging

import "time"

// LogInfo is a base log information
type LogInfo struct {
	ClientIP    string        `json:"ip"`
	Date        string        `json:"date"`
	Method      string        `json:"method"`
	RequestURI  string        `json:"uri"`
	Referer     string        `json:"referer,omitempty"`
	HTTPVersion string        `json:"httpVersion"`
	Size        int           `json:"size"`
	Status      int           `json:"status"`
	UserAgent   string        `json:"userAgent"`
	Latency     time.Duration `json:"latency"`
}

// AccessLog is a log information of user access
type AccessLog struct {
	LogInfo
	Error error `json:"error,omitempty"`
}

// ActivityLog is a log information of user action
type ActivityLog struct {
	LogInfo
	RequestBody map[string]interface{} `json:"requestBody,omitempty"`
	Extra       interface{}            `json:"extra,omitempty"`
}
