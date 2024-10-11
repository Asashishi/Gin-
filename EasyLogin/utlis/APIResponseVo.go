package utlis

type Succeed struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
