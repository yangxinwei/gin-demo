package types


type RegisterReq struct {
	Request
}

type RegisterResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}


type LoginReq struct {
	Mobile string `json:"mobile" binding:"required"`
	Code   string `json:"code" binding:"required"`
	OpenID string `json:"openid" binding:"required"`
}

type LoginResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}


type SMSCodeReq struct {
	UserRequest
}

type SMSCodeResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}



type GetUserInfoReq struct {
	UserRequest
}

type GetUserInfoResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}



type ChangeMobileReq struct {
	UserRequest
}

type ChangeMobileResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}


type FindPasswordReq struct {
	UserRequest
}

type FindPasswordResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}



type ChangePasswordReq struct {
	UserRequest
	New string `json:"new" binding:"required"`
}

type ChangePasswordResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}


type LogoutReq struct {
	UserRequest
}

type LogoutResp struct {
	Response
	Data struct {
		UserID   string `json:"userId"`
		NickName string `json:"nickName"`
	} `json:"data"`
}
