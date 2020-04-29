package entity

type DownloadFileInfo struct {
	VersionId          int `json: "VersionId"`
	TextVersion        string
	FiasCompleteXmlUrl string
	FiasDeltaXmlUrl    string
}
