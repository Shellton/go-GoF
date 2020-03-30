package abstractfactory

type TranslateFactory interface {
	CreateReceiver() Receiver
	CreateSender() Sender
	CreateWatchDog() WatchDog
}

type Receiver interface {
	StartReceive(func([]byte))
}

type Sender interface {
	Send([]byte)
}

type WatchDog interface {
	StartWatch(func())
}
