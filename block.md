
https://blockchain-fabric.blogspot.com/2017/04/hyperledger-fabric-v10-block-structure.html 描述了block的结构

![block picture]( https://i.imgur.com/TkLG3W7.png)


message Block {
    BlockHeader header = 1;
    BlockData data = 2;
    BlockMetadata metadata = 3;
}
message BlockHeader {
    uint64 number = 1; // The position in the blockchain  从0开始， genesis block
    bytes previous_hash = 2; // The hash of the previous block header, genesis block 对应的值是nil， 对应的是前面一个blockheader的sha256，而不是block的sha256
    bytes data_hash = 3; // The hash of the BlockData, by MerkleTree
}
 SizeOf(BlockHeader) = 8 bytes for Number + 32 bytes for PreviousHash + 32 bytes for  DataHash = 52 bytes

message BlockData {
    repeated bytes data = 1;  --》 Envelope
}



**********************************************
一共有4个。
message BlockMetadata {
    repeated bytes metadata = 1;
}

The four metadata stored in a block are listed below along with its index.
【orderer增加的】SIGNATURES: signature on the block creation.  (index: 0)
【orderer增加的】LAST_CONFIG: reference to the last configuration block.  (index: 1)
https://github.com/hyperledger/fabric/blob/release/orderer/multichain/chainsupport.go#L249  其中253 cs.lastConfig = block.Header.Number还没有明白。
【peer增加的】TRANSACTIONS_FILTER: valid and invalid transactions in a block.  (index: 2)
【orderer增加的】ORDERER: last offset persisted (kafka ordering metadata).  (index: 3)
对象是：KafkaMetadata。

结构上就是：
// This enum enlists indexes of the block metadata array
enum BlockMetadataIndex {
    SIGNATURES = 0;             // Block metadata array position for block signatures
    LAST_CONFIG = 1;            // Block metadata array position to store last configuration block sequence number
    TRANSACTIONS_FILTER = 2;    // Block metadata array position to store serialized bit array filter of invalid polcis
    ORDERER = 3;                // Block metadata array position to store operational metadata for orderers
                                // e.g. For Kafka, this is where we store the last offset written to the local ledger.
}


// Metadata is a common structure to be used to encode block metadata，目前仅仅包括last_config,orderer 情况：
message Metadata {
    bytes value = 1;
    repeated MetadataSignature signatures = 2;
}

message MetadataSignature {
    bytes signature_header = 1; // An encoded SignatureHeader
    bytes signature = 2;       // The signature over the concatenation of the Metadata value bytes, signatureHeader, and block header
}


说明：The Creator field holds x.509 certificate, publick key and membership service provided (MSP) who issued these identity to the client. The Nounce field contains some random bytes.

message SignatureHeader {
    // Creator of the message, specified as a certificate chain
    bytes creator = 1;  --》需要核实下。

    // Arbitrary number that may only be used once. Can be used to detect replay attacks.
    bytes nonce = 2;
}
代码：https://github.com/hyperledger/fabric/blob/release/common/localmsp/signer.go#L38

metadata：SIGNATURES也是采用message Metadata结构,不过value=nil, 代码在：
https://github.com/hyperledger/fabric/blob/release/orderer/multichain/chainsupport.go#L227
**********************************************



https://blockchain-fabric.blogspot.com/2017/04/under-construction-hyperledger-fabric.html?showComment=1520231379453  
关于blockData的描述：  以后继续

message Envelope {
    // A marshaled Payload
    bytes payload = 1;

    // A signature by the creator specified in the Payload header
    bytes signature = 2;
}


// Payload is the message contents (and header to allow for signing)
message Payload {

    // Header is included to provide identity and prevent replay
    Header header = 1;

    // Data, the encoding of which is defined by the type in the header
    bytes data = 2;
}

message Header {
    bytes channel_header = 1;
    bytes signature_header = 2;
}


// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
    Type      int32 
    Version   int32 
    Timestamp *google_protobuf.Timestamp 
    ChannelId string 
    TxId      string 
    Epoch     uint64 
    Extension []byte 
}

message ChaincodeHeaderExtension {
	bytes payload_visibility = 1; 目前 为nil，payload全部可见
	ChaincodeID chaincode_id = 2;
}

message SignatureHeader {
    // Creator of the message, specified as a certificate chain
    bytes creator = 1;

    // Arbitrary number that may only be used once. Can be used to detect replay attacks.
    bytes nonce = 2;
}

// These status codes are intended to resemble selected HTTP status codes
enum Status {
    UNKNOWN = 0;
    SUCCESS = 200;
    BAD_REQUEST = 400;
    FORBIDDEN = 403;
    NOT_FOUND = 404;
    REQUEST_ENTITY_TOO_LARGE = 413;
    INTERNAL_SERVER_ERROR = 500;
    SERVICE_UNAVAILABLE = 503;
}

定义了7中事务类型，其中二种是过渡的。
enum HeaderType {
    MESSAGE = 0;                   // Used for messages which are signed but opaque
    CONFIG = 1;                    // Used for messages which express the channel config
    CONFIG_UPDATE = 2;             // Used for transactions which update the channel config  过渡的
    ENDORSER_TRANSACTION = 3;      // Used by the SDK to submit endorser based transactions
    ORDERER_TRANSACTION = 4;       // Used internally by the orderer for management。 是orderer内部增加的。过渡的
    DELIVER_SEEK_INFO = 5;         // Used as the type for Envelope messages submitted to instruct the Deliver API to seek
    CHAINCODE_PACKAGE = 6;         // Used for packaging chaincode artifacts for install
}


