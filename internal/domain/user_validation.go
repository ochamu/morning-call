package domain

func (rcv *User) CanAcceptMorningCall(friendID UserID) NGReason {
	for _, ru := range rcv.RelatedUsers {
		if ru.ID == friendID {
			return ""
		}
	}
	return "フレンドではありません。"
}