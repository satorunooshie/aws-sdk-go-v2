// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Translates input text from the source language to the target language. For a
// list of available languages and language codes, see what-is-languages.
func (c *Client) TranslateText(ctx context.Context, params *TranslateTextInput, optFns ...func(*Options)) (*TranslateTextOutput, error) {
	if params == nil {
		params = &TranslateTextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TranslateText", params, optFns, c.addOperationTranslateTextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TranslateTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TranslateTextInput struct {

	// The language code for the language of the source text. The language must be a
	// language supported by Amazon Translate. For a list of language codes, see
	// what-is-languages. To have Amazon Translate determine the source language of
	// your text, you can specify auto in the SourceLanguageCode field. If you specify
	// auto, Amazon Translate will call Amazon Comprehend
	// (https://docs.aws.amazon.com/comprehend/latest/dg/comprehend-general.html) to
	// determine the source language.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code requested for the language of the target text. The language
	// must be a language supported by Amazon Translate.
	//
	// This member is required.
	TargetLanguageCode *string

	// The text to translate. The text string can be a maximum of 5,000 bytes long.
	// Depending on your character set, this may be fewer than 5,000 characters.
	//
	// This member is required.
	Text *string

	// Settings to configure your translation output, including the option to mask
	// profane words and phrases.
	Settings *types.TranslationSettings

	// The name of the terminology list file to be used in the TranslateText request.
	// You can use 1 terminology list at most in a TranslateText request. Terminology
	// lists can contain a maximum of 256 terms.
	TerminologyNames []string

	noSmithyDocumentSerde
}

type TranslateTextOutput struct {

	// The language code for the language of the source text.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code for the language of the target text.
	//
	// This member is required.
	TargetLanguageCode *string

	// The translated text.
	//
	// This member is required.
	TranslatedText *string

	// Settings that configure the translation output.
	AppliedSettings *types.TranslationSettings

	// The names of the custom terminologies applied to the input text by Amazon
	// Translate for the translated text response.
	AppliedTerminologies []types.AppliedTerminology

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTranslateTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTranslateText{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTranslateText{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpTranslateTextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTranslateText(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opTranslateText(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "TranslateText",
	}
}
