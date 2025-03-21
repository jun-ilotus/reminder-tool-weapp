package wxnotice

const (
	ReminderTemplateID = "T6iprDxSmNa_hmMQDSrfJAGxTulxZh3dkBTycKWpXlI"
	_                  = iota
	TypeReminder
)

type (
	Item struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}

	// MessageLotteryStart 提醒的消息内容
	MessageReminder struct {
		ReminderContent Item `json:"thing1"`
		ReminderTime    Item `json:"time2"`
	}
)

func (m *MessageReminder) Type() int {
	return TypeReminder
}

func (m *MessageReminder) TemplateId() string {
	return ReminderTemplateID
}
