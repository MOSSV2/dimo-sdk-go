package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	mrand "math/rand"
	"net"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shirou/gopsutil/v3/disk"
	"golang.org/x/crypto/sha3"
)

func ECDSAToAddr(sk *ecdsa.PrivateKey) common.Address {
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func HexToAddr(sk string) common.Address {
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return common.Address{}
	}

	return ECDSAToAddr(privateKey)
}

func ToEthAddress(pubkey []byte) []byte {
	if len(pubkey) == 65 {
		d := sha3.NewLegacyKeccak256()
		d.Write(pubkey[1:])
		payload := d.Sum(nil)
		return payload[12:]
	}

	return pubkey
}

func GetDiskStatus(path string) (types.DiskStats, error) {
	m := types.DiskStats{
		Path: path,
	}
	dus, err := disk.Usage(path)
	if err != nil {
		return m, err
	}

	m.Total = dus.Total
	m.Free = dus.Free

	m.Used = m.Total - m.Free
	return m, nil
}

func DataToName(d []byte) string {
	dsum := sha256.Sum256(d)
	return hex.EncodeToString(dsum[:])
}

func HexToAddress(addr string) common.Address {
	return common.HexToAddress(addr)
}

func RandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)

	return randomBytes
}

func ShuffleString(array []string) {
	var temp string
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num := r.Intn(i + 1)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}
}

func LocalIp() string {
	address, _ := net.InterfaceAddrs()
	var ip = "localhost"
	for _, address := range address {
		if ipAddress, ok := address.(*net.IPNet); ok && !ipAddress.IP.IsLoopback() {
			if ipAddress.IP.To4() != nil {
				ip = ipAddress.IP.String()
			}
		}
	}
	return ip
}

func Disorder(array []interface{}) {
	var temp interface{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num := r.Intn(i + 1)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}
}
