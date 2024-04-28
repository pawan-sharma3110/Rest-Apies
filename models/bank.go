package models

type Bank struct {
	BranchName string `json:"branch_name"`
	IfscCode   string `json:"ifsc_code"`
	Address    string `json:"address"`
	Accounts   []Account `json:"accounts"`
	// Employee   []Employee `json:"employee"`
}