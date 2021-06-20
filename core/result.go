package core

import "time"

type Result struct {
	URL               string        `json:"url"`
	Timestamp         time.Time     `json:"timestamp"`
	HTTPStatusCode    int           `json:"httpstatuscode"`
	Duration          time.Duration `json:"duration"`
	CertificateExpiry time.Duration `json:"certificateexpiry"`
}
