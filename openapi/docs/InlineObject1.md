# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthConsumerKey** | **string** | The 25-character hexadecimal string that was obtained from Interactive Brokers during the OAuth consumer registration process. | [optional] 
**OauthToken** | **string** | The request token obtained from IB via /request_token. | [optional] 
**OauthSignatureMethod** | **string** | The signature method used to sign the request. Currently only &#39;RSA-SHA256&#39; is supported. | [optional] 
**OauthSignature** | **string** | The signature for the request generated using the method specified in the oauth_signature_method parameter. See section 9 of the OAuth v1.0a specification for more details on signing requests. | [optional] 
**OauthTimestamp** | **string** | Timestamp expressed in seconds since 1/1/1970 00:00:00 GMT. Must be a positive integer and greater than or equal to any timestamp used in previous requests. | [optional] 
**OauthNonce** | **string** | A random string uniquely generated for each request. | [optional] 
**OauthVerifier** | **string** | The verification code received from IB after the user has granted authorization. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


