package core

import "time"

type Result struct {
	ServiceName       string        `json:"servicename"`
	URL               string        `json:"url"`
	Timestamp         time.Time     `json:"timestamp"`
	HTTPStatusCode    int           `json:"httpstatuscode"`
	Duration          time.Duration `json:"duration"`
	CertificateExpiry time.Duration `json:"certificateexpiry"`
}
