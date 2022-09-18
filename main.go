package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"crypto/ecdsa"
	"math/big"

	"github.com/EdsonAlcala/proof-id-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wabarc/ipfs-pinner/pkg/pinata"

	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Proofs struct {
	Pi_A     []int   `json:"pi_a"`
	Pi_B     [][]int `json:"pi_b"`
	Pi_C     []int   `json:"pi_c"`
	Protocol string  `json:"protocol"`
	Curve    string  `json:"curve"`
}

func main() {
	e := echo.New()

	e.GET("/generateProof", generateProof)
	e.GET("/verifyProof/:tokenid", verifyProof)
	e.Logger.Fatal(e.Start(":8080"))
}

func generateProof(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// binding body data
	price := c.QueryParam("price")
	times := c.QueryParam("times")
	salary := c.QueryParam("salary")

	fmt.Println("Generating a proof for paying the following amount: ", price)
	// Generate proof calling generate witness
	generateWitnessPath := os.Getenv("WITNESS_PATH")
	circomPath := os.Getenv("CIRCOM_PATH")

	proofs := &exec.Cmd{
		Path:   generateWitnessPath,
		Args:   []string{"sh " + generateWitnessPath, circomPath, price, times, salary},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fmt.Println("Proofs: ", proofs)

	// Pin data to IPFS
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	proofPath := os.Getenv("PROOF_PATH")
	fmt.Printf("Proof path: %s\n", proofPath)
	pnt := pinata.Pinata{Apikey: apiKey, Secret: secretKey}
	hash, err := pnt.PinFile(proofPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hash IPFS", hash)

	// var proofContract *contract.ProofContracts
	proofContractAddress := os.Getenv("PROOF_CONTRACT_ADDRESS")
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	client, _ := ethclient.Dial(rpcEndpoint)
	contractAddress := common.HexToAddress(proofContractAddress)
	proofContract, _ := contract.NewProofContract(contractAddress, client)

	privateKeyString := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	fmt.Printf("public key: %s\n", publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// // ***********************************************************************
	tokenUri := "https://api.ipfsbrowser.com/ipfs/get.php?hash=" + hash

	addressToAssignNft := os.Getenv("ADDRESS_TO_ASSIGN_NFT")
	addressToAssignTo := common.HexToAddress(addressToAssignNft)

	tx, err := proofContract.MintNFT(auth, addressToAssignTo, tokenUri)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return c.String(http.StatusOK, "hash: "+hash)
}

func verifyProof(c echo.Context) error {
	tokenId := c.Param("tokenId")
	fmt.Printf("tokenId: %s\n", tokenId)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	proofContractAddress := os.Getenv("PROOF_CONTRACT_ADDRESS")
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	client, _ := ethclient.Dial(rpcEndpoint)
	contractAddress := common.HexToAddress(proofContractAddress)
	proofContract, _ := contract.NewProofContract(contractAddress, client)

	n := new(big.Int)
	n, _ = n.SetString(tokenId, 10)
	nft, _ := proofContract.TokenURI(&bind.CallOpts{}, n)
	fmt.Printf("ntf uri: %v\n", nft)
	hash := "QmXnNHFAj3LocJvLiPScLirnxeiRjHctLkeoq7K32zYeWv"
	// DUMMY HASH
	nft_ipfs := "https://api.ipfsbrowser.com/ipfs/get.php?hash=" + hash
	// retrieve proof from IPFS
	resp, err := http.Get(nft_ipfs)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO prepare data
	sb := string(body)
	fmt.Println("sb: ", sb)

	A := [2]*big.Int{}
	B := [2][2]*big.Int{[2]*big.Int{}, [2]*big.Int{}}
	C := [2]*big.Int{}
	D := [1]*big.Int{}

	// Contract call
	verifierContractAddress := os.Getenv("VERIFIER_CONTRACT_ADDRESS")
	vContractAddress := common.HexToAddress(verifierContractAddress)
	verifierContract, _ := contract.NewRangeProofVerifierContract(vContractAddress, client)

	auth := &bind.CallOpts{
		Pending: true,
	}

	tx, err := verifierContract.VerifyProof(auth, A, B, C, D)
	fmt.Printf("tx sent: %v", tx)
	result := "true"
	return c.String(http.StatusOK, "proof: "+result)

}
