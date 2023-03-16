// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// The case-insensitive input to indicate standard MIME type that describes the
// format of the file that will be uploaded.
type AttachmentItem struct {

	// A unique identifier for the attachment.
	AttachmentId *string

	// A case-sensitive name of the attachment being uploaded.
	AttachmentName *string

	// Describes the MIME file type of the attachment. For a list of supported file
	// types, see Feature specifications
	// (https://docs.aws.amazon.com/connect/latest/adminguide/feature-limits.html) in
	// the Amazon Connect Administrator Guide.
	ContentType *string

	// Status of the attachment.
	Status ArtifactStatus

	noSmithyDocumentSerde
}

// Connection credentials.
type ConnectionCredentials struct {

	// The connection token.
	ConnectionToken *string

	// The expiration of the token. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	Expiry *string

	noSmithyDocumentSerde
}

// An item - message or event - that has been sent.
type Item struct {

	// The time when the message or event was sent. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// Provides information about the attachments.
	Attachments []AttachmentItem

	// The contactId on which the transcript item was originally sent. This field is
	// populated only when the transcript item is from the current chat session.
	ContactId *string

	// The content of the message or event.
	Content *string

	// The type of content of the item.
	ContentType *string

	// The chat display name of the sender.
	DisplayName *string

	// The ID of the item.
	Id *string

	// The metadata related to the message. Currently this supports only information
	// related to message receipts.
	MessageMetadata *MessageMetadata

	// The ID of the sender in the session.
	ParticipantId *string

	// The role of the sender. For example, is it a customer, agent, or system.
	ParticipantRole ParticipantRole

	// The contactId on which the transcript item was originally sent. This field is
	// only populated for persistent chats when the transcript item is from the past
	// chat session. For more information, see Enable persistent chat
	// (https://docs.aws.amazon.com/connect/latest/adminguide/chat-persistence.html).
	RelatedContactId *string

	// Type of the item: message or event.
	Type ChatItemType

	noSmithyDocumentSerde
}

// Contains metadata related to a message.
type MessageMetadata struct {

	// The identifier of the message that contains the metadata information.
	MessageId *string

	// The list of receipt information for a message for different recipients.
	Receipts []Receipt

	noSmithyDocumentSerde
}

// The receipt for the message delivered to the recipient.
type Receipt struct {

	// The time when the message was delivered to the recipient.
	DeliveredTimestamp *string

	// The time when the message was read by the recipient.
	ReadTimestamp *string

	// The identifier of the recipient of the message.
	RecipientParticipantId *string

	noSmithyDocumentSerde
}

// A filtering option for where to start. For example, if you sent 100 messages,
// start with message 50.
type StartPosition struct {

	// The time in ISO format where to start. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// The ID of the message or event where to start.
	Id *string

	// The start position of the most recent message where you want to start.
	MostRecent int32

	noSmithyDocumentSerde
}

// Fields to be used while uploading the attachment.
type UploadMetadata struct {

	// The headers to be provided while uploading the file to the URL.
	HeadersToInclude map[string]string

	// This is the pre-signed URL that can be used for uploading the file to Amazon S3
	// when used in response to StartAttachmentUpload
	// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_StartAttachmentUpload.html).
	Url *string

	// The expiration time of the URL in ISO timestamp. It's specified in ISO 8601
	// format: yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	UrlExpiry *string

	noSmithyDocumentSerde
}

// The websocket for the participant's connection.
type Websocket struct {

	// The URL expiration timestamp in ISO date format. It's specified in ISO 8601
	// format: yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	ConnectionExpiry *string

	// The URL of the websocket.
	Url *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
