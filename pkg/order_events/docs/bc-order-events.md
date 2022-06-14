# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [order_events.proto](#order_events.proto)
    - [BcTxIdentity](#order_events.BcTxIdentity)
    - [BcTxInfo](#order_events.BcTxInfo)
    - [Order](#order_events.Order)
    - [OrderEvent](#order_events.OrderEvent)
    - [OrderIdentity](#order_events.OrderIdentity)
  
    - [BcTxStatus](#order_events.BcTxStatus)
    - [OrderStatus](#order_events.OrderStatus)
    - [OrderType](#order_events.OrderType)
  
- [Scalar Value Types](#scalar-value-types)



<a name="order_events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order_events.proto



<a name="order_events.BcTxIdentity"></a>

### BcTxIdentity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TxID | [string](#string) |  |  |
| Vout | [uint32](#uint32) |  |  |






<a name="order_events.BcTxInfo"></a>

### BcTxInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| BcTxIdentifier | [BcTxIdentity](#order_events.BcTxIdentity) |  |  |
| Confirmations | [uint64](#uint64) |  |  |
| BlockNumber | [uint64](#uint64) |  |  |
| Status | [BcTxStatus](#order_events.BcTxStatus) |  |  |






<a name="order_events.Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderIdentifier | [OrderIdentity](#order_events.OrderIdentity) |  |  |
| Type | [OrderType](#order_events.OrderType) |  |  |
| Status | [OrderStatus](#order_events.OrderStatus) |  |  |
| NetworkIdentifier | [uint32](#uint32) |  |  |
| ProviderIdentifier | [uint32](#uint32) |  |  |
| MerchantIdentifier | [string](#string) |  |  |
| Amount | [uint64](#uint64) |  |  |
| RealFee | [uint64](#uint64) |  |  |
| EstimationFee | [uint64](#uint64) |  |  |
| EstimationFeeIdentifier | [string](#string) |  |  |
| WalletIdentifier | [string](#string) |  |  |
| AddressIdentifier | [string](#string) |  |  |
| Currency | [string](#string) |  |  |
| BcTx | [BcTxInfo](#order_events.BcTxInfo) | repeated |  |






<a name="order_events.OrderEvent"></a>

### OrderEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| EventIdentity | [string](#string) |  |  |
| NetworkIdentifier | [uint32](#uint32) |  |  |
| ProviderIdentifier | [uint32](#uint32) |  |  |
| MerchantIdentifier | [string](#string) |  |  |
| OrderIdentifier | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Producer | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| OrderInfo | [Order](#order_events.Order) |  |  |






<a name="order_events.OrderIdentity"></a>

### OrderIdentity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UUID | [string](#string) |  |  |
| ExternalOrderUUID | [string](#string) |  |  |
| ProviderOrderUUID | [string](#string) |  |  |





 


<a name="order_events.BcTxStatus"></a>

### BcTxStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PLACEHOLDER_BC_TX_STATUS | 0 |  |
| BC_TX_IN_MINING | 1 |  |
| BC_TX_SUCCESS | 2 |  |
| BC_TX_EXECUTION_FAILED | 3 |  |



<a name="order_events.OrderStatus"></a>

### OrderStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PLACEHOLDER_ORDER_STATUS | 0 |  |
| NEW | 1 |  |
| PROCESSING | 2 |  |
| CONFIRMED | 3 |  |
| FAILED | 4 |  |
| SUCCESS | 5 |  |
| CANCELED | 6 |  |



<a name="order_events.OrderType"></a>

### OrderType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PLACEHOLDER_ORDER_TYPE | 0 |  |
| WITHDRAW | 1 |  |
| DEPOSIT | 2 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

