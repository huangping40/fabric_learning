package counterfeiter_example

type PeerLedgerProvider interface {
	Create(genesisBlock *Block) (PeerLedger, error)
	// Open opens an already created ledger
	Open(ledgerID string) (PeerLedger, error)
	// Exists tells whether the ledger with given id exists
	Exists(ledgerID string) (bool, error)
	// List lists the ids of the existing ledgers
	List() ([]string, error)
	// Close closes the PeerLedgerProvider
	Close()
}

type PeerLedger interface {
}

type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
}

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

type MySpecialInterface interface {
	DoThings(string, uint64) (int, error)
}
