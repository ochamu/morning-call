package domain

// CanAcceptMorningCall checks if the user can accept a morning call from the specified friend
func (rcv *User) CanAcceptMorningCall(friendID UserID) NGReason {
	for _, ru := range rcv.RelatedUsers {
		if ru.ID == friendID {
			if ru.Status == RelatedUserStatusApproved {
				return ""
			}
			if ru.Status == RelatedUserStatusBlocked {
				return NGReasonBlocked
			}
			return NGReasonNotFriend
		}
	}
	return NGReasonNotFriend
}

// CanAddFriend checks if the user can add a friend
func (rcv *User) CanAddFriend(targetUserID UserID) NGReason {
	// 自分自身は追加できない
	if rcv.ID == targetUserID {
		return NGReasonSelfOperation
	}

	// 既存の関係をチェック
	for _, ru := range rcv.RelatedUsers {
		if ru.ID == targetUserID {
			switch ru.Status {
			case RelatedUserStatusApproved:
				return NGReasonAlreadyFriend
			case RelatedUserStatusPending:
				return NGReasonAlreadyRequested
			case RelatedUserStatusBlocked:
				return NGReasonBlockedByUser
			case RelatedUserStatusRejected:
				// 拒否された場合は再申請可能
				return ""
			}
		}
	}

	return ""
}

// CanApproveFriend checks if the user can approve a friend request
func (rcv *User) CanApproveFriend(requestingUserID UserID) NGReason {
	for _, ru := range rcv.RelatedUsers {
		if ru.ID == requestingUserID {
			if ru.Status == RelatedUserStatusPending {
				return ""
			}
			if ru.Status == RelatedUserStatusApproved {
				return NGReasonAlreadyFriend
			}
			if ru.Status == RelatedUserStatusBlocked {
				return NGReasonBlocked
			}
			return NGReasonInvalidStatus
		}
	}
	return NGReasonUserNotFound
}

// CanBlockUser checks if the user can block another user
func (rcv *User) CanBlockUser(targetUserID UserID) NGReason {
	// 自分自身はブロックできない
	if rcv.ID == targetUserID {
		return NGReasonSelfOperation
	}

	for _, ru := range rcv.RelatedUsers {
		if ru.ID == targetUserID {
			if ru.Status == RelatedUserStatusBlocked {
				return NGReasonAlreadyRequested // 既にブロック済み
			}
			return ""
		}
	}

	return ""
}
