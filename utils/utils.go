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
func (b Blockchain) AddBlock(s *SampleData) {
	lastBlock := b.Blocks[len(b.Blocks)-1]
	lastHash := lastBlock.Hash
	newBlock := NewBlock(s, lastBlock)
	if lastHash == newBlock.PrevHash {
		if newBlock.ValidateHash(newBlock.Hash) {
			b.Blocks = append(b.Blocks, newBlock)
		}

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

var NewBlockchain Blockchain = Blockchain{[]*Block{FirstBlock()}}

func CurrentBlock() *Blockchain {
	return &NewBlockchain
}
