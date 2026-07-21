package respones

type DataTableUserResp struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Avatar    string `json:"image"`
	Status    string `json:"status"`
	Custom    string `json:"custom"`
}
