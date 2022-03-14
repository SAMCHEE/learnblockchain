package blockchain

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	b := new(Block)
	b.PrevHash = prevHash
	b.Data = []byte(data)
	pow := NewProof(b)
	hash, nonce := pow.Run()
	b.Hash = hash
	b.Nonce = nonce
	return b
}

func GenesisBlock() *Block { return CreateBlock("Genesis", []byte{}) }
