# GetJwtPublicKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** |  | 
**PublicKeyBase64** | **string** | Base64-encoded Ed25519 public key | 

## Methods

### NewGetJwtPublicKey200Response

`func NewGetJwtPublicKey200Response(alg string, publicKeyBase64 string, ) *GetJwtPublicKey200Response`

NewGetJwtPublicKey200Response instantiates a new GetJwtPublicKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJwtPublicKey200ResponseWithDefaults

`func NewGetJwtPublicKey200ResponseWithDefaults() *GetJwtPublicKey200Response`

NewGetJwtPublicKey200ResponseWithDefaults instantiates a new GetJwtPublicKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *GetJwtPublicKey200Response) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *GetJwtPublicKey200Response) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *GetJwtPublicKey200Response) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetPublicKeyBase64

`func (o *GetJwtPublicKey200Response) GetPublicKeyBase64() string`

GetPublicKeyBase64 returns the PublicKeyBase64 field if non-nil, zero value otherwise.

### GetPublicKeyBase64Ok

`func (o *GetJwtPublicKey200Response) GetPublicKeyBase64Ok() (*string, bool)`

GetPublicKeyBase64Ok returns a tuple with the PublicKeyBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyBase64

`func (o *GetJwtPublicKey200Response) SetPublicKeyBase64(v string)`

SetPublicKeyBase64 sets PublicKeyBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


