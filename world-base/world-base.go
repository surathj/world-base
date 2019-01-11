package main

import (
	"log"
	"net/http"
	"encoding/json"
	bc "github.com/surathj/world-base/blockchain"
)

func main(){
	blockchain := bc.NewBlockChain()

	http.HandleFunc("/world_base/mine_block", func(w http.ResponseWriter, r *http.Request) {
		previousBlock := blockchain.GetPreviousBlock()
		previousProof := previousBlock["proof"].(int)
		proof := blockchain.ProofOfWork(previousProof)
		previousHash := blockchain.Hash(previousBlock)
		block := blockchain.CreateBlock(proof, previousHash)
		jsonBlock, _ := json.Marshal(block)
		w.Write(jsonBlock)
	})

	http.HandleFunc("/world_base/get_chain", func(w http.ResponseWriter, r *http.Request){
		jsonChain, _ := json.Marshal(blockchain)
		w.Write(jsonChain)
	})

	log.Println("HTTP SERVER RUNNING ON PORT 3000")
	http.ListenAndServe("localhost:3000", nil)
	log.Println("HTTP SERVER STOPPED")
}
