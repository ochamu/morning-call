package domain

// NGReason定数定義 - バリデーションメッセージ用
const (
	// User関連のNGReason
	NGReasonNotFriend        NGReason = "フレンドではありません。"
	NGReasonAlreadyFriend    NGReason = "既にフレンドです。"
	NGReasonAlreadyRequested NGReason = "既にフレンド申請済みです。"
	NGReasonBlocked          NGReason = "ブロックされています。"
	NGReasonBlockedByUser    NGReason = "このユーザーをブロックしています。"
	NGReasonUserNotFound     NGReason = "ユーザーが見つかりません。"
	NGReasonSelfOperation    NGReason = "自分自身に対する操作はできません。"
	NGReasonNoPermission     NGReason = "権限がありません。"
	NGReasonInvalidStatus    NGReason = "無効なステータスです。"
	NGReasonPendingRequest   NGReason = "承認待ちのリクエストがあります。"

	// MorningCall関連のNGReason
	NGReasonInvalidTime         NGReason = "無効な時刻設定です。"
	NGReasonPastTime            NGReason = "過去の時刻は設定できません。"
	NGReasonTooFarInFuture      NGReason = "設定可能な期間を超えています。"
	NGReasonAlreadyCompleted    NGReason = "既に完了しています。"
	NGReasonAlreadyDeleted      NGReason = "既に削除されています。"
	NGReasonNotSender           NGReason = "送信者ではありません。"
	NGReasonNotReceiver         NGReason = "受信者ではありません。"
	NGReasonMorningCallNotFound NGReason = "モーニングコールが見つかりません。"
	NGReasonDuplicateSchedule   NGReason = "同じ時刻に既にモーニングコールが設定されています。"

	// バリデーション関連のNGReason
	NGReasonInvalidEmail     NGReason = "無効なメールアドレス形式です。"
	NGReasonInvalidUsername  NGReason = "無効なユーザー名です。"
	NGReasonUsernameTooShort NGReason = "ユーザー名が短すぎます。"
	NGReasonUsernameTooLong  NGReason = "ユーザー名が長すぎます。"
	NGReasonMessageTooLong   NGReason = "メッセージが長すぎます。"
	NGReasonEmptyMessage     NGReason = "メッセージが空です。"
	NGReasonInvalidParameter NGReason = "無効なパラメータです。"
)
