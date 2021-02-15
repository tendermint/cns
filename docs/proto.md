# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [cns/cns.proto](#cns/cns.proto)
    - [ChainInfo](#tendermint.cns.cns.ChainInfo)
    - [VersionInfo](#tendermint.cns.cns.VersionInfo)
  
- [cns/query.proto](#cns/query.proto)
    - [Query](#tendermint.cns.cns.Query)
  
- [cns/tx.proto](#cns/tx.proto)
    - [MsgRegisterChainNameRequest](#tendermint.cns.cns.MsgRegisterChainNameRequest)
    - [MsgRegisterChainNameResponse](#tendermint.cns.cns.MsgRegisterChainNameResponse)
    - [MsgUpdateChainInfoRequest](#tendermint.cns.cns.MsgUpdateChainInfoRequest)
    - [MsgUpdateChainInfoResponse](#tendermint.cns.cns.MsgUpdateChainInfoResponse)
  
    - [Msg](#tendermint.cns.cns.Msg)
  
- [cns/genesis.proto](#cns/genesis.proto)
    - [GenesisState](#tendermint.cns.cns.GenesisState)
  
- [Scalar Value Types](#scalar-value-types)



<a name="cns/cns.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cns/cns.proto



<a name="tendermint.cns.cns.ChainInfo"></a>

### ChainInfo
TODO(sahith): Add json and yaml flags


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chain_name | [string](#string) |  |  |
| expiration | [int64](#int64) |  |  |
| metadata | [string](#string) | repeated |  |
| owner | [string](#string) |  |  |
| canonical_ibc_client | [string](#string) |  |  |
| seed | [string](#string) | repeated |  |
| source_code_url | [string](#string) |  |  |
| version | [VersionInfo](#tendermint.cns.cns.VersionInfo) |  |  |






<a name="tendermint.cns.cns.VersionInfo"></a>

### VersionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [int64](#int64) |  |  |
| source_commit_hash | [string](#string) |  |  |
| genesis_hash | [string](#string) |  |  |





 

 

 

 



<a name="cns/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cns/query.proto


 

 

 


<a name="tendermint.cns.cns.Query"></a>

### Query
Query defines the gRPC querier service.

this line is used by starport scaffolding # 2

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="cns/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cns/tx.proto



<a name="tendermint.cns.cns.MsgRegisterChainNameRequest"></a>

### MsgRegisterChainNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chain_name | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| seed | [string](#string) | repeated |  |
| source_code_url | [string](#string) |  |  |
| canonical_ibc_client | [string](#string) |  |  |
| version | [VersionInfo](#tendermint.cns.cns.VersionInfo) |  |  |






<a name="tendermint.cns.cns.MsgRegisterChainNameResponse"></a>

### MsgRegisterChainNameResponse







<a name="tendermint.cns.cns.MsgUpdateChainInfoRequest"></a>

### MsgUpdateChainInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chain_name | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| seed | [string](#string) | repeated |  |
| source_code_url | [string](#string) |  |  |
| canonical_ibc_client | [string](#string) |  |  |
| version | [VersionInfo](#tendermint.cns.cns.VersionInfo) |  |  |






<a name="tendermint.cns.cns.MsgUpdateChainInfoResponse"></a>

### MsgUpdateChainInfoResponse






 

 

 


<a name="tendermint.cns.cns.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterChainName | [MsgRegisterChainNameRequest](#tendermint.cns.cns.MsgRegisterChainNameRequest) | [MsgRegisterChainNameResponse](#tendermint.cns.cns.MsgRegisterChainNameResponse) |  |
| UpdateChainInfo | [MsgUpdateChainInfoRequest](#tendermint.cns.cns.MsgUpdateChainInfoRequest) | [MsgUpdateChainInfoResponse](#tendermint.cns.cns.MsgUpdateChainInfoResponse) |  |

 



<a name="cns/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cns/genesis.proto



<a name="tendermint.cns.cns.GenesisState"></a>

### GenesisState
GenesisState defines the capability module&#39;s genesis state.

this line is used by starport scaffolding # genesis/proto/state





 

 

 

 



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

