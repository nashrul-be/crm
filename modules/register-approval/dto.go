package register_approval

type Representation struct {
	Username []string `json:"username"`
}

type ApproveRequest struct {
	Approved []string `json:"approved"`
	Rejected []string `json:"rejected"`
}

type ApproveResponse struct {
	Success []string `json:"success"`
	Fail    []string `json:"fail"`
}
