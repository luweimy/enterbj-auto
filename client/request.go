package client

type Resp struct {
	RespCode string `json:"rescode"`
	RespDesc  string `json:"resdes"`
}

type PersonInfo struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Sex             string `json:"sex"`
	HeadUrl         string `json:"headurl"`
	DriverLicenseNo string `json:"driverlicenseno"`
	Resp
}
