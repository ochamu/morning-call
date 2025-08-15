package domain

// フレンド申請のステータス管理
type RelatedUser struct {
	ID       UserID
	Username string
	Email    string
	Status   RelatedUserStatus
}