// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delivers a message to the specified queue.
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed:
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// Any characters not included in this list will be rejected. For more
// information, see the [W3C specification for characters].
//
// [W3C specification for characters]: http://www.w3.org/TR/REC-xml/#charsets
func (c *Client) SendMessage(ctx context.Context, params *SendMessageInput, optFns ...func(*Options)) (*SendMessageOutput, error) {
	if params == nil {
		params = &SendMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendMessage", params, optFns, c.addOperationSendMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendMessageInput struct {

	// The message to send. The minimum size is one character. The maximum size is 256
	// KiB.
	//
	// A message can include only XML, JSON, and unformatted text. The following
	// Unicode characters are allowed:
	//
	// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
	//
	// Any characters not included in this list will be rejected. For more
	// information, see the [W3C specification for characters].
	//
	// [W3C specification for characters]: http://www.w3.org/TR/REC-xml/#charsets
	//
	// This member is required.
	MessageBody *string

	// The URL of the Amazon SQS queue to which a message is sent.
	//
	// Queue URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	//  The length of time, in seconds, for which to delay a specific message. Valid
	// values: 0 to 900. Maximum: 15 minutes. Messages with a positive DelaySeconds
	// value become available for processing after the delay period is finished. If you
	// don't specify a value, the default value for the queue applies.
	//
	// When you set FifoQueue , you can't set DelaySeconds per message. You can set
	// this parameter only on a queue level.
	DelaySeconds int32

	// Each message attribute consists of a Name , Type , and Value . For more
	// information, see [Amazon SQS message attributes]in the Amazon SQS Developer Guide.
	//
	// [Amazon SQS message attributes]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-metadata.html#sqs-message-attributes
	MessageAttributes map[string]types.MessageAttributeValue

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of sent messages. If a message with a
	// particular MessageDeduplicationId is sent successfully, any messages sent with
	// the same MessageDeduplicationId are accepted successfully but aren't delivered
	// during the 5-minute deduplication interval. For more information, see [Exactly-once processing]in the
	// Amazon SQS Developer Guide.
	//
	//   - Every message must have a unique MessageDeduplicationId ,
	//
	//   - You may provide a MessageDeduplicationId explicitly.
	//
	//   - If you aren't able to provide a MessageDeduplicationId and you enable
	//   ContentBasedDeduplication for your queue, Amazon SQS uses a SHA-256 hash to
	//   generate the MessageDeduplicationId using the body of the message (but not the
	//   attributes of the message).
	//
	//   - If you don't provide a MessageDeduplicationId and the queue doesn't have
	//   ContentBasedDeduplication set, the action fails with an error.
	//
	//   - If the queue has ContentBasedDeduplication set, your MessageDeduplicationId
	//   overrides the generated one.
	//
	//   - When ContentBasedDeduplication is in effect, messages with identical content
	//   sent within the deduplication interval are treated as duplicates and only one
	//   copy of the message is delivered.
	//
	//   - If you send one message with ContentBasedDeduplication enabled and then
	//   another message with a MessageDeduplicationId that is the same as the one
	//   generated for the first MessageDeduplicationId , the two messages are treated
	//   as duplicates and only one copy of the message is delivered.
	//
	// The MessageDeduplicationId is available to the consumer of the message (this
	// can be useful for troubleshooting delivery issues).
	//
	// If a message is sent successfully but the acknowledgement is lost and the
	// message is resent with the same MessageDeduplicationId after the deduplication
	// interval, Amazon SQS can't detect duplicate messages.
	//
	// Amazon SQS continues to keep track of the message deduplication ID even after
	// the message is received and deleted.
	//
	// The maximum length of MessageDeduplicationId is 128 characters.
	// MessageDeduplicationId can contain alphanumeric characters ( a-z , A-Z , 0-9 )
	// and punctuation ( !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ ).
	//
	// For best practices of using MessageDeduplicationId , see [Using the MessageDeduplicationId Property] in the Amazon SQS
	// Developer Guide.
	//
	// [Using the MessageDeduplicationId Property]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html
	// [Exactly-once processing]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-exactly-once-processing.html
	MessageDeduplicationId *string

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The tag that specifies that a message belongs to a specific message group.
	// Messages that belong to the same message group are processed in a FIFO manner
	// (however, messages in different message groups might be processed out of order).
	// To interleave multiple ordered streams within a single queue, use MessageGroupId
	// values (for example, session data for multiple users). In this scenario,
	// multiple consumers can process the queue, but the session data of each user is
	// processed in a FIFO fashion.
	//
	//   - You must associate a non-empty MessageGroupId with a message. If you don't
	//   provide a MessageGroupId , the action fails.
	//
	//   - ReceiveMessage might return messages with multiple MessageGroupId values.
	//   For each MessageGroupId , the messages are sorted by time sent. The caller
	//   can't specify a MessageGroupId .
	//
	// The maximum length of MessageGroupId is 128 characters. Valid values:
	// alphanumeric characters and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~) .
	//
	// For best practices of using MessageGroupId , see [Using the MessageGroupId Property] in the Amazon SQS Developer
	// Guide.
	//
	// MessageGroupId is required for FIFO queues. You can't use it for Standard
	// queues.
	//
	// [Using the MessageGroupId Property]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagegroupid-property.html
	MessageGroupId *string

	// The message system attribute to send. Each message system attribute consists of
	// a Name , Type , and Value .
	//
	//   - Currently, the only supported message system attribute is AWSTraceHeader .
	//   Its type must be String and its value must be a correctly formatted X-Ray
	//   trace header string.
	//
	//   - The size of a message system attribute doesn't count towards the total size
	//   of a message.
	MessageSystemAttributes map[string]types.MessageSystemAttributeValue

	noSmithyDocumentSerde
}

// The MD5OfMessageBody and MessageId elements.
type SendMessageOutput struct {

	// An MD5 digest of the non-URL-encoded message attribute string. You can use this
	// attribute to verify that Amazon SQS received the message correctly. Amazon SQS
	// URL-decodes the message before creating the MD5 digest. For information about
	// MD5, see [RFC1321].
	//
	// [RFC1321]: https://www.ietf.org/rfc/rfc1321.txt
	MD5OfMessageAttributes *string

	// An MD5 digest of the non-URL-encoded message body string. You can use this
	// attribute to verify that Amazon SQS received the message correctly. Amazon SQS
	// URL-decodes the message before creating the MD5 digest. For information about
	// MD5, see [RFC1321].
	//
	// [RFC1321]: https://www.ietf.org/rfc/rfc1321.txt
	MD5OfMessageBody *string

	// An MD5 digest of the non-URL-encoded message system attribute string. You can
	// use this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest.
	MD5OfMessageSystemAttributes *string

	// An attribute containing the MessageId of the message sent to the queue. For
	// more information, see [Queue and Message Identifiers]in the Amazon SQS Developer Guide.
	//
	// [Queue and Message Identifiers]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html
	MessageId *string

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The large, non-consecutive number that Amazon SQS assigns to each message.
	//
	// The length of SequenceNumber is 128 bits. SequenceNumber continues to increase
	// for a particular MessageGroupId .
	SequenceNumber *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendMessage"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateSendMessageChecksum(stack, options); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSendMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendMessage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSendMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendMessage",
	}
}
