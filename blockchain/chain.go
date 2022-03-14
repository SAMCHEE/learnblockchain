package blockchain

type Chain struct {
	Blocks []*Block
}

func (c *Chain) AddBlock(data string) {
	lastBlock := c.Blocks[len(c.Blocks)-1]
	prevHash := lastBlock.Hash
	c.Blocks = append(c.Blocks, CreateBlock(data, prevHash))
}

func InitChain() *Chain {
	c := new(Chain)
	c.Blocks = append(c.Blocks, GenesisBlock())
	return c
}
