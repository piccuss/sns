package notify

import "sns/internal/pkg"

type Notifier interface {
	sendEmail(reciever string, data []pkg.Stock)
}

var notifiers = make(map[string]Notifier)

func registerNotifiers(name string, notifier Notifier) {
	if notifiers[name] != nil {
		pkg.Sugar().Warnf("try to add notifier with same name. name=%s, notifier=%s", name, notifier)
		return
	}
	notifiers[name] = notifier
	pkg.Sugar().Infof("add notifier, name=%s", name)
}

func SendEmail(reciever, notifyVendor string, data []*pkg.Stock) error {
	//TODO implement sendEmail
	return nil
}
