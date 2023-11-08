package response

type ListMemberResponse struct {
	MemberID      string `json:"member_id"`
	MemberName    string `json:"member_name"`
	MemberEmail   string `json:"member_email"`
	MemberAddress string `json:"member_address"`
	MemberPhone   string `json:"member_phone"`
}
