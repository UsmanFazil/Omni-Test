# omni
**omni** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Before running the project, set the Ethereum RPC endpoint environment variable:
export ETHEREUM_RPC_ENDPOINT="https://your-ethereum-rpc-url-here"

### Usage

To fetch data from the Ethereum chain and store it:

Omnid tx omni fetch-eth-data "queryid" --from bob

Replace "queryid" with the unique identifier for your query. This ID will be used to associate the fetched Ethereum data in the Omni chain.

For example, if you want to fetch data and associate it with the ID 12345, you would run:

Omnid tx omni fetch-eth-data "12345" --from bob

To retrieve the saved Ethereum data based on a specific ID:

Omnid q omni show-eth-data <query_id>

Replace <query_id> with the unique identifier associated with the data you want to retrieve.

For example, to retrieve data associated with the ID "12345", you would run:

Omnid q omni show-eth-data 12345


### Future improvments

The Omni module provides essential functionality for fetching and storing Ethereum data. 
However, to enhance its utility, I would do following things.

1. Dynamic Contract Address and Storage Slot
Current Limitation:
The contract address and storage slot are currently hardcoded, limiting the module's flexibility to query different Ethereum contracts and their respective storage slots.

Proposed Improvement:
Allow users to pass the contract address and storage slot as parameters in the query. This change will make the module more versatile, enabling users to fetch data from any Ethereum contract without needing to modify the code.


Omnid tx omni fetch-eth-data <query_id> --contract=<contract_address> --slot=<storage_slot> --from=<user>

2. Enhanced Error Handling
Current Limitation:
The current error handling mechanism might not provide detailed feedback on specific issues, making it challenging for users to understand and rectify problems.

Proposed Improvement:
Implement comprehensive error handling to provide clear and actionable feedback. This includes:

Descriptive Error Messages: Instead of generic error messages, provide detailed descriptions pinpointing the exact issue. For instance, instead of "Failed to connect," use "Failed to connect to Ethereum RPC at [endpoint]."

Retry Mechanisms: For transient issues, like temporary network glitches, implement a retry mechanism with exponential backoff to automatically attempt the operation again.

Error Logging: Log errors with timestamps, error codes, and other relevant details. This will aid in debugging and issue resolution.

User Feedback: Ensure that all errors are communicated back to the user with clarity, suggesting potential solutions or steps to rectify them.