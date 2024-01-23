package ViewModels

import "wan-api-spam-report/DTO"

type SpamReportViewModel struct {
	Result       string              `json:"result"`       //Result : success, unsuccess
	IsExisted    bool                `json:"isExisted"`    // Existence of the News
	ErrorMessage string              `json:"errorMessage"` // The query error
	Guid         string              `json:"guid"`
	SpamReportID int64               `json:"spamReportID"`
	SpamReport   *DTO.SpamReportsDTO `json:"spamReport"`
}

type ListSpamReportsViewModel struct {
	Result         string                `json:"result"`       //Result : success, unsuccess
	ErrorMessage   string                `json:"errorMessage"` // The query error
	PageIndex      int64                 `json:"pageIndex"`
	PageSize       int64                 `json:"pageSize"`
	TotalCount     int64                 `json:"totalCount"`  // The number of News
	ListSpamReport *[]DTO.SpamReportsDTO `json:"spamReports"` // List of the News
}
