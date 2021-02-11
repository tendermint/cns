# Chain Name Service Draft

The CNS module is a new Cosmos SDK module that allows Cosmos addresses to purchase, sell, and administer chain names. The development of CNS is planned to be divided into 3 separate phases in order to ensure there's enough time for existing chains and prevent front running.

## Phase 0

The goal of phase 0 is to launch a directory of chain names for Cosmos SDK chains that currently exist. For expediency, we plan on setting up a multisig that initially has the authority to modify information about blockchains on the CNS.  After all existing major blockchains are recorded on the CNS and phase 0 is over the multisig will hand off these chain names one by one to trusted parties from the various sovereign Cosmos chain communities.

## Phase 1

The goal of phase 1 is to allow users of Cosmos Hub to purchase chain names and assign them to the blockchain networks they have affiliation with. Proceeds from the sale of chain names will go to the community pool. For example, a newly launched chain using its governance mechanism can elect an account that the chain will trust to register a chain name on CNS.

## Phase 2

The goal of phase 2 is name transfers.  If by the time CNS reaches phase 2 an NFT module becomes available on the Hub, chain names will conform to the NFT standard. Otherwise, CNS will include functionality to transfer ownership of names through a proposal system where a “buyer” sends a proposal with an atom value (bid) and the current chain name owner can either accept or reject the proposal. If a proposal is accepted, the bid value is transferred to the seller and the chain name ownership is transferred to the buyer.

# State

Goal 1 (end-users): let users determine whether the tokens they got through IBC came from a chain they claim to come from. This will be possible through the mapping of chain names to IBC clients.

Goal 2 (validators): let nodes (validators, full-nodes, etc.) join networks.

```go
type ChainInfo struct {
	ChainName     string
	Owner         OwnerProps
	Expiration    int64
	Metadata      [][2]string

	// Goal 1
	CanonicalIBCClientId string

	// Goal 2
	Seed          []seed
	SourceCodeURL string //how likely is the possibility of code url changing
	Version       VersionInfo
}
```

```go
type VersionInfo struct {
	 Version          int64
	 SourceCommitHash string
	 GenesisHash      string
}
```

```go
type OwnerProps interface {
	String()  string
}
```

```go
type Claim struct {
  ID        int64
  ChainName string
  Owner     OwnerProps
  Proof     string
}
```

# Messages

### Phase - 0

Ideally, we would want the chains to approve via a gov proposal the owner (likely address) of the chain name. This makes it easy for approvers and also prevent collision for the same namespace.

### MsgClaimChainName

Claim is about mapping chain name to owner. Name, owner, IBC client, meta.

```go
type MsgClaimChainName struct {
	ChainName     string
	CanonicalIBCClientId string
	Proof         string
	Owner         OwnerProps
}
```

### MsgApproveClaim

```go
type MsgApproveClaim struct {
	ClaimID  int32
	Approver sdk.AccAddress
}
```

### MsgRejectClaim

```go
type MsgRejectClaim struct{
	ClaimID int32
	Approver sdk.AccAddress
}
```

### MsgUpdateInfo

```go
type MsgUpdateInfo struct{
	ChainName     string
	Owner         OwnerProps
	Seed          []seed
	CanonicalIBCClientId string
	Version       VersionInfo
	Metadata      [][2]string
}
```