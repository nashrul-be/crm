package register_approval

type Representation struct {
	Username []string `json:"username"`
}

type ApproveRequest struct {
	Approve []string `json:"approve"`
	Reject  []string `json:"reject"`
}

type ApproveResponse struct {
	Success []string `json:"success"`
	Fail    []string `json:"fail"`
}
