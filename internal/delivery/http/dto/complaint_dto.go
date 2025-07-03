package dto

import "time"

type ComplaintResponse struct {
	ID               string    `json:"id"`
	ComplaintID      string    `json:"complaint_id"`
	ReportedName     string    `json:"reported_name"`
	Email            string    `json:"email"`
	InsidentLocation string    `json:"insident_location"`
	InsidentTime     time.Time `json:"insident_time"`
	Description      string    `json:"description"`
	EvidenceUrl      string    `json:"evidence_url"`
	ReporterName     string    `json:"reporter_name"`
	ReporterPhone    string    `json:"reporter_phone"`
	Status           string    `json:"status"`
	ComplaintType    string    `json:"complaint_type"`
}

type ComplaintGetQueryParams struct {
	Key    string `query:"key"`
	Status string `query:"status"`
	Type   string `query:"type"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

type ComplaintCreateRequest struct {
	ReportedName     string    `json:"reported_name" form:"reported_name"`
	EvidenceFile     string    `json:"evidence_file" form:"evidence_file"`
	InsidentLocation string    `json:"insident_location" form:"insident_location"`
	InsidentTime     time.Time `json:"insident_time" form:"insident_time"`
	Description      string    `json:"description" form:"description"`
	ReporterName     string    `json:"reporter_name" form:"reporter_name"`
	ReporterEmail    string    `json:"reporter_email" form:"reporter_email"`
	ReporterPhone    string    `json:"reporter_phone" form:"reporter_phone"`
	ComplaintType    string    `json:"complaint_type" form:"complaint_type"`
}
