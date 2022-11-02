package notify

import (
	"fmt"

	"github.com/juunini/simple-go-line-notify/notify"
)

func SimpleNotifyPunchStatus(sid string, lineToken string, punchStatus string) {
	accessToken := lineToken
	message := sid + " : " + punchStatus

	if err := notify.SendText(accessToken, message); err != nil {
		fmt.Println(err)
	}

}
