// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Chat undocumented
type Chat struct {
	// Entity is the base model of Chat
	Entity
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastUpdatedDateTime undocumented
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
	// Members undocumented
	Members []ConversationMember `json:"members,omitempty"`
	// Messages undocumented
	Messages []ChatMessage `json:"messages,omitempty"`
	// InstalledApps undocumented
	InstalledApps []TeamsAppInstallation `json:"installedApps,omitempty"`
}

// ChatActivityStatistics undocumented
type ChatActivityStatistics struct {
	// ActivityStatistics is the base model of ChatActivityStatistics
	ActivityStatistics
	// AfterHours undocumented
	AfterHours *Duration `json:"afterHours,omitempty"`
}

// ChatInfo undocumented
type ChatInfo struct {
	// Object is the base model of ChatInfo
	Object
	// ThreadID undocumented
	ThreadID *string `json:"threadId,omitempty"`
	// MessageID undocumented
	MessageID *string `json:"messageId,omitempty"`
	// ReplyChainMessageID undocumented
	ReplyChainMessageID *string `json:"replyChainMessageId,omitempty"`
}

// ChatMembersNotificationAudience undocumented
type ChatMembersNotificationAudience struct {
	// Object is the base model of ChatMembersNotificationAudience
	Object
}

// ChatMessage undocumented
type ChatMessage struct {
	// Entity is the base model of ChatMessage
	Entity
	// ReplyToID undocumented
	ReplyToID *string `json:"replyToId,omitempty"`
	// From undocumented
	From *IdentitySet `json:"from,omitempty"`
	// Etag undocumented
	Etag *string `json:"etag,omitempty"`
	// MessageType undocumented
	MessageType *ChatMessageType `json:"messageType,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// DeletedDateTime undocumented
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// Summary undocumented
	Summary *string `json:"summary,omitempty"`
	// Attachments undocumented
	Attachments []ChatMessageAttachment `json:"attachments,omitempty"`
	// Mentions undocumented
	Mentions []ChatMessageMention `json:"mentions,omitempty"`
	// Importance undocumented
	Importance *ChatMessageImportance `json:"importance,omitempty"`
	// PolicyViolation undocumented
	PolicyViolation *ChatMessagePolicyViolation `json:"policyViolation,omitempty"`
	// Reactions undocumented
	Reactions []ChatMessageReaction `json:"reactions,omitempty"`
	// Locale undocumented
	Locale *string `json:"locale,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// Replies undocumented
	Replies []ChatMessage `json:"replies,omitempty"`
	// HostedContents undocumented
	HostedContents []ChatMessageHostedContent `json:"hostedContents,omitempty"`
}

// ChatMessageAttachment undocumented
type ChatMessageAttachment struct {
	// Object is the base model of ChatMessageAttachment
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}

// ChatMessageHostedContent undocumented
type ChatMessageHostedContent struct {
	// Entity is the base model of ChatMessageHostedContent
	Entity
}

// ChatMessageMention undocumented
type ChatMessageMention struct {
	// Object is the base model of ChatMessageMention
	Object
	// ID undocumented
	ID *int `json:"id,omitempty"`
	// MentionText undocumented
	MentionText *string `json:"mentionText,omitempty"`
	// Mentioned undocumented
	Mentioned *IdentitySet `json:"mentioned,omitempty"`
}

// ChatMessagePolicyViolation undocumented
type ChatMessagePolicyViolation struct {
	// Object is the base model of ChatMessagePolicyViolation
	Object
	// DlpAction undocumented
	DlpAction *ChatMessagePolicyViolationDlpActionTypes `json:"dlpAction,omitempty"`
	// JustificationText undocumented
	JustificationText *string `json:"justificationText,omitempty"`
	// PolicyTip undocumented
	PolicyTip *ChatMessagePolicyViolationPolicyTip `json:"policyTip,omitempty"`
	// UserAction undocumented
	UserAction *ChatMessagePolicyViolationUserActionTypes `json:"userAction,omitempty"`
	// VerdictDetails undocumented
	VerdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes `json:"verdictDetails,omitempty"`
}

// ChatMessagePolicyViolationPolicyTip undocumented
type ChatMessagePolicyViolationPolicyTip struct {
	// Object is the base model of ChatMessagePolicyViolationPolicyTip
	Object
	// GeneralText undocumented
	GeneralText *string `json:"generalText,omitempty"`
	// ComplianceURL undocumented
	ComplianceURL *string `json:"complianceUrl,omitempty"`
	// MatchedConditionDescriptions undocumented
	MatchedConditionDescriptions []string `json:"matchedConditionDescriptions,omitempty"`
}

// ChatMessageReaction undocumented
type ChatMessageReaction struct {
	// Object is the base model of ChatMessageReaction
	Object
	// ReactionType undocumented
	ReactionType *string `json:"reactionType,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// User undocumented
	User *IdentitySet `json:"user,omitempty"`
}