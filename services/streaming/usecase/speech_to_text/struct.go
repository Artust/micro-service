package speech_to_text

type Speaker string

var (
	SpeakerOperator Speaker = "operator"
	SpeakerGuest    Speaker = "guest"
)

type SpeechToTextMessage struct {
	Speaker     Speaker `json:"speaker"`
	Content     string  `json:"content"`
	SendingTime string  `json:"sendingTime"`
}
