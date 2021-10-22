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
	Generation	string	`json:"generation"`
	BlogURL		string	`json:"blog_url"`
	ImgURL		string	`json:"img_url"`
}

type FormationResponse struct {
	Single		string		`json:"single"`
	Title		string		`json:"title"`
	Positions	[] Position `json:"positions"`
}

type Position struct {
	MemberName	string	`json:"member_name"`
	Position	string	`json:"position"`
	IsCenter	bool	`json:"is_center"`
}

type GetSongsResponse struct {
	Single	string		`json:"single"`
	Title	string		`json:"title"`
	Center	[]string	`json:"center"`
}

type GetPositionsResponse struct {
	MemberName	string	`json:"member_name"`
	ImgURL		string	`json:"img_url"`
	Position	string	`json:"position"`
	IsCenter	bool	`json:"is_center"`
}
