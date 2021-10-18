package api

type MemberInfosResponse struct {
	MemberInfos	[]MemberInfoResponse `json:"members"`
}

type MemberInfoResponse struct {
	MemberId	int		`json:"user_id"`
	MemberName	string	`json:"user_name" binding:"required,min=1"`
	Birthday	string	`json:"birthday"`
	Height		string	`json:"height"`
	BloodType	string	`json:"blood_type"`
	BlogURL		string	`json:"blog_url"`
	ImgURL		string	`json:"img_url"`
}
