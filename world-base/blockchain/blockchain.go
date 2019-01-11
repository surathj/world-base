package blockchain

import (
	"time"
	//"log"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"encoding/json"
)

type IBlockchain interface {
	GetLength() int
	GetPreviousBlock() map[string]interface{}
	CreateBlock(proof int, prevHash string) interface{}
	ProofOfWork(prevProof int) int
	Hash(block map[string]interface{}) string
}

var _BlockChain *BlockChain

type BlockChain struct {
	Chain []map[string]interface{}
}

func NewBlockChain() *BlockChain {
	if _BlockChain != nil {
		return _BlockChain
	}
	_BlockChain = new(BlockChain)
	node := make(map[string]interface{})
	node["index"] = len(_BlockChain.Chain) + 1
	node["timestamp"] = time.Now().String()
	node["proof"] = 1
	node["previous_hash"] = 0
	_BlockChain.Chain = append(_BlockChain.Chain, node)
	return _BlockChain
}

func (blockChain *BlockChain) GetLength() int {
	return len(blockChain.Chain)
}

func (blockChain *BlockChain) GetPreviousBlock() map[string]interface{} {
	return blockChain.Chain[blockChain.GetLength()-1]
}

func (blockchain *BlockChain) CreateBlock(proof int, prevHash string) interface{} {
	block := make(map[string]interface{})
	block["index"] = blockchain.GetLength()+1
	block["timestamp"] = time.Now().String()
	block["proof"] = proof
	block["previous_hash"] = prevHash

	blockchain.Chain = append(blockchain.Chain, block)

	return block
}

func (blockChain *BlockChain) ProofOfWork(prevProof int) int {
	newProof := 1 // new nonce
	checkProof := false

	for checkProof == false {
		hasher := sha256.New()
		hasher.Write([]byte(strconv.Itoa(newProof - prevProof)))
		hashOperation := hex.EncodeToString(hasher.Sum(nil))
		if hashOperation[:4] == `0000` {
			checkProof = true
		} else {
			newProof += 1
		}
	}
	return newProof
}

func (blockChain *BlockChain) Hash(block map[string]interface{}) string {
	encodedBlock, _ := json.Marshal(block)
	hasher := sha256.New()
	hasher.Write(encodedBlock)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (blockChain *BlockChain) IsChainValid(chain []map[string]interface{}) bool {
	previousBlock := chain[0]
	blockIndex := 1

	for blockIndex < len(chain) {
		block := chain[blockIndex]
		if block["previous_hash"] != blockChain.Hash(previousBlock) {
			return false
		}
		previousProof := previousBlock["proof"].(int)
		proof := block["proof"].(int)
		hasher := sha256.New()
		hasher.Write([]byte(strconv.Itoa(proof - previousProof)))
		hashOperation := hex.EncodeToString(hasher.Sum(nil))
		if hashOperation[:4] != `0000` {
			return false
		}
		previousBlock = block
		blockIndex += 1
	}
	return true
}