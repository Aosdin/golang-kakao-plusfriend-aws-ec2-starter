package models



type Keyboard struct {
	Type string `json:"type"`
}

type KakaoRequest struct {
	User_key        string `json:"user_key"`
	Type        string `json:"type"`
	Content        string `json:"content"`
}
type KakaoMessage struct {
	Message     map[string]string  `json:"message"`
	Keyboard    map[string]string `json:"keyboard"`
}
func KakaoKeyboardModel() Keyboard{
	return Keyboard{
		Type: "text",
	}
}
func KakaoMessageModel(request KakaoRequest) KakaoMessage {

	return KakaoMessage{
		Message: map[string]string{"text": "User_key="+request.User_key+" Type="+request.Type+" Content="+request.Content + " entity.Name="},
		Keyboard: map[string]string{"type": "text"},
	}
}