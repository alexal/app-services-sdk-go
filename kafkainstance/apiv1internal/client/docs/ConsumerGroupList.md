# ConsumerGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ConsumerGroup**](ConsumerGroup.md) | Consumer group list items | [optional] 
**Total** | Pointer to **float32** | The total number of consumer groups. | [optional] 
**Size** | Pointer to **float32** | The number of consumer groups per page. | [optional] 
**Page** | Pointer to **int32** | The page | [optional] 


## Methods

### NewConsumerGroupList

`func NewConsumerGroupList() *ConsumerGroupList`

NewConsumerGroupList instantiates a new ConsumerGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupListWithDefaults

`func NewConsumerGroupListWithDefaults() *ConsumerGroupList`

NewConsumerGroupListWithDefaults instantiates a new ConsumerGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetItems

`func (o *ConsumerGroupList) GetItems() []ConsumerGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConsumerGroupList) GetItemsOk() (*[]ConsumerGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConsumerGroupList) SetItems(v []ConsumerGroup)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConsumerGroupList) HasItems() bool`

HasItems returns a boolean if a field has been set.


### GetTotal

`func (o *ConsumerGroupList) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConsumerGroupList) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConsumerGroupList) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConsumerGroupList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


### GetSize

`func (o *ConsumerGroupList) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ConsumerGroupList) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ConsumerGroupList) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ConsumerGroupList) HasSize() bool`

HasSize returns a boolean if a field has been set.


### GetPage

`func (o *ConsumerGroupList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ConsumerGroupList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ConsumerGroupList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ConsumerGroupList) HasPage() bool`

HasPage returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

