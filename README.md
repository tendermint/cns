# Chain Name Service (CNS)

The CNS module is a new Cosmos SDK module that allows Cosmos addresses to obtain and administer chain names.

## Why Cosmos Hub needs a CNS

For the Cosmos Hub to function as an interchain router, it needs to access a canonical list of chain names. For example, token holders on the Hub need a trusted way to determine that, e.g., IRIS tokens on the Hub are actually from the IRISnet mainnet. If the Hub decides to add services like DEXes onto the Hub, trusted chain names allow token holders to more quickly verify ticker symbols and provide liquidity to the right pools.

A CNS on the Hub will provide security to interchain transfers by registering chain names, this will allow proper token tracking for every token transferred onto cosmos hub and off the cosmos hub. The cosmos hub will then become the defacto hub with all other zones relying on the trusted setup on the hub to define valid tokens. This value proposition will contribute to the growth of the ATOM.

While the CNS can exist on another chain and function correctly, we think that the CNS has the most aligned incentives with the Cosmos Hub. By adopting this module, ATOM stakers will be securing a crucial component of the interchain and will move us further into the direction of a Hub-centric model of Cosmos.
