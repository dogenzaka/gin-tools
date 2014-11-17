package logging

import "time"

type logInfo struct {
	ClientIP    string        `json:"ip"`
	Date        string        `json:"date"`
	Method      string        `json:"method"`
	Path        string        `json:"path"`
	RawQuery    string        `json:"rawQuery,omitempty"`
	HTTPVersion string        `json:"httpVersion"`
	Size        int           `json:"size"`
	Status      int           `json:"status"`
	UserAgent   string        `json:"userAgent"`
	Latency     time.Duration `json:"latency"`
}

type accessLog struct {
	logInfo
	Error error `json:"error,omitempty"`
}

type activityLog struct {
	logInfo
	UserID      string                 `json:"userId"`
	RequestBody map[string]interface{} `json:"requestBody"`
}
