package abstractfactory

type TcpTransporter struct {
}

func (*TcpTransporter) CreateReceiver() Receiver {
	return &TcpReceiver{}
}

func (*TcpTransporter) CreateSender() Sender {
	return &TcpSender{}
}

func (*TcpTransporter) CreateWatchDog() WatchDog {
	return &TcpWatchDog{}
}

type TcpReceiver struct{}

func (*TcpReceiver) StartReceive(callback func([]byte)) {

}

type TcpSender struct{}

func (*TcpSender) Send([]byte) {

}

type TcpWatchDog struct {
}

func (*TcpWatchDog) StartWatch(callback func()) {
}
