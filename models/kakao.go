package models


type Keyboard struct {
	Type string `json:"type"`
}
type KakaoMessageRequest struct {
	User_key        string `json:"user_key"`
	Type        string `json:"type"`
	Content        string `json:"content"`
}
type KakaoMessage struct {
	Message     map[string]string  `json:"message"`
	Keyboard    map[string]string `json:"keyboard"`
}
type DialogflowQuery struct {
	Query	string `json:"query"`
	Lang	string `json:"lang"`
	SessionId	string `json:"sessionId"`
}