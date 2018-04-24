package model

// Build for Jenkins Job Report
type Build struct {
	JobID          int64  `json:"jobid"`
	Job            string `json:"job"`
	RepositoryURL  string `json:"repositoryurl"`
	RepositoryName string `json:"repositoryname"`
	BranchName     string `json:"branchname"`
	Status         string `json:"status"`
	Slave          string `json:"slave"`
	Author         string `json:"author"`
	Duration       int64  `json:"duration"`
}
