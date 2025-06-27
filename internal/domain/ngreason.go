package domain

type NGReason string

func (rcv NGReason) IsNG() bool {
	return rcv != ""
}

func (rcv NGReason) IsOK() bool {
	return !rcv.IsNG()
}

func (rcv NGReason) String() string {
	return string(rcv)
}
