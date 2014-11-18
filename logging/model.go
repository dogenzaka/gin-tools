package logging

import "time"

type logInfo struct {
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

type accessLog struct {
	logInfo
	Error error `json:"error,omitempty"`
}

type activityLog struct {
	logInfo
	RequestBody map[string]interface{} `json:"requestBody,omitempty"`
	Extra       interface{}            `json:"extra,omitempty"`
}
