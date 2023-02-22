package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The bigchain driver source couldn't reasonably reach bigchaindb.
// As far as go is concerned,
// so the MongoDB base is the same, so it was implemented directly.

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
}

type BlockChain struct {
	Blocks     []*Block
	Collection *mongo.Collection
}

func (chain *BlockChain) NewBlock(data string) *Block {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := &Block{
		Index:     len(chain.Blocks),
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Hash:      "",
	}
	_, err := chain.Collection.InsertOne(nil, bson.M{
		"index":     newBlock.Index,
		"timestamp": newBlock.Timestamp,
		"data":      newBlock.Data,
		"prevHash":  newBlock.PrevHash,
		"hash":      newBlock.Hash,
	})
	if err != nil {
		log.Fatal(err)
	}
	return newBlock
}

func HashBlock(block *Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp.String(), block.Data, block.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func NewBlockChain() *BlockChain {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("blockdb").Collection("blocks")
	collection.Drop(nil)

	genesisBlock := &Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}

	genesisBlock.Hash = HashBlock(genesisBlock)
	_, err = collection.InsertOne(nil, bson.M{
		"index":     genesisBlock.Index,
		"timestamp": genesisBlock.Timestamp,
		"data":      genesisBlock.Data,
		"prevHash":  genesisBlock.PrevHash,
		"hash":      genesisBlock.Hash,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &BlockChain{
		Blocks:     []*Block{genesisBlock},
		Collection: collection,
	}
}

func main() {
	chain := NewBlockChain()
	chain.NewBlock("Transaction Data 1")
	chain.NewBlock("Transaction Data 2")
}
