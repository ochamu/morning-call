package domain

// システムの利用者
type User struct {
	ID           UserID
	Username     string
	Email        string
	MorningCalls []MorningCall
	RelatedUsers []RelatedUser
}
