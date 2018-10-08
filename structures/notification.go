package structures

type NotificationList struct {
	Id uint `json:"id"`
	Image string `json:"image"`
	Message string `json:"message"`
	Header string `json:"header"`
	Priority uint `json:"priority"`
	Expired string `json:"expired"`
	Button string `json:"button"`
	Link string `json:"link"`
	Company string `json:"company"`
}

type NotificationPublic struct {
	Id uint `json:"id"`
	Image string `json:"image"`
	Message string `json:"message"`
	Header string `json:"header"`
	Priority uint `json:"priority"`
	Expired string `json:"expired"`
	Button string `json:"button"`
	Link string `json:"link"`
}