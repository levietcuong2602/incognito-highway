package p2p

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multihash"
)

func ParseListenner(s, defaultIP string, defaultPort int) (string, int) {
	if s == "" {
		return defaultIP, defaultPort
	}
	splitStr := strings.Split(s, ":")
	if len(splitStr) > 1 {
		p, e := strconv.Atoi(splitStr[1])
		if e != nil {
			panic(e)
		}
		return splitStr[0], p
	}
	return splitStr[0], 0
}

func IDFromPublicKey(pk crypto.PubKey) (peer.ID, error) {
	b, err := pk.Bytes()
	if err != nil {
		return "", err
	}
	var alg uint64 = multihash.SHA2_256
	hash, _ := multihash.Sum(b, alg, -1)
	return peer.ID(hash), nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func generateRand() []byte {
	res := make([]byte, 40)
	// fmt.Println(time.Now().UnixNano())
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 40; i++ {
		rand := byte(rand.Intn(256))
		res[i] = rand
	}
	return res
}
