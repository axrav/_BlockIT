package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Pos       int
	Data      SampleData
	TimeStamp string
	Hash      string
	PrevHash  string
}

func (bl *Block) GenerateHash() {
	bytes, _ := json.Marshal(bl.Data)
	data := fmt.Sprint(bl.Pos) + bl.TimeStamp + string(bytes)
	hash := sha256.New()
	hash.Write([]byte(data))
	bl.Hash = hex.EncodeToString(hash.Sum(nil))
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(data *SampleData, lastBlock *Block) *Block {
	newBlock := &Block{
		Pos:       lastBlock.Pos + 1,
		PrevHash:  lastBlock.Hash,
		TimeStamp: time.Now().String(),
	}
	newBlock.GenerateHash()
	return newBlock

}
func (bl *Block) ValidateHash(hash string) bool {
	bl.GenerateHash()
	return bl.Hash == hash
}
func (b *Blockchain) AddBlock(s *SampleData) {
	lastBlock := b.Blocks[len(b.Blocks)-1]
	lastHash := lastBlock.Hash
	newBlock := NewBlock(s, lastBlock)
	if lastHash == newBlock.PrevHash {
		if newBlock.ValidateHash(newBlock.Hash) {
			b.Blocks = append(b.Blocks, newBlock)
		}

	}

}

func (b *Blockchain) Background() {
	for _, block := range b.Blocks {
		bytes, _ := json.MarshalIndent(block.Data, "", " ")
		fmt.Printf("PrevHash: %s\nData:%s\nHash: %s\nTimestamp: %s", block.PrevHash, string(bytes), block.Hash, block.TimeStamp)
	}
}

type SampleData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Amount    int    `json:"amount"`
	IsGenesis bool   `json:"is_genesis"`
}

func FirstBlock() *Block {
	return NewBlock(&SampleData{IsGenesis: true}, &Block{})

}

func CurrentBlock() *Blockchain {
	NewBlockchain := Blockchain{[]*Block{FirstBlock()}}
	return &NewBlockchain
}
