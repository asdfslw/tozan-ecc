package std_test

import (
	"encoding/hex"
	"fmt"
	"github.com/armfazh/tozan-ecc/curve/std"
	"math/big"
	"testing"
)

// 测试 点乘
func Test(t *testing.T) {
	tests := []struct {
		alg    std.ID
		wantPx string
		wantPy string
		wantQx string
		wantQy string
	}{
		{
			std.SECP256K1,
			"175e159f728b865a72f99cc6c6fc846de0b93833fd2222ed73fce5b551e5b739",
			"d3506e0d9e3c79eba4ef97a51ff71f5eacb5955add24345c6efa6ffee9fed695",
			"723cbaa6e5db996d6bf771c00bd548c7b700dbffa6c0e77bcb6115925232fcda",
			"96e867b5595cc498a921137488824d6e2660a0653779494801dc069d9eb39f5f",
		},
		{
			std.SECP256R1,
			"e716aed2cf069e4d997789672e6d6bd2508676f2f4fd0a64f077e8daa245573f",
			"353663e694fc72ab5912b06687b9a851d13d0df2fa07c9b3505fc26b469218d1",
			"f8f5dccf4c6a93d7a4a54daafaa3449aa87a8069875405d43725c5dce392d805",
			"e58176cf66d63054389d3e336461327351f3da64a52143ba026619516cda02fa",
		},
		{
			std.STARKNET,
			"029f26828c9f9616373509c0ab067a1bbe3e4ca7440e0a238bb3f1c7e396286e",
			"01ed0925b1705cbfc9c7ba9aa0b6950721e330b4ca6a06e08b3134f3c9ff31a2",
			"0314c8986b9a3fca1998ab928443205e5e1d387f373c55e67b18fef00f647e90",
			"06772826cd306607f55f25bea21e93fc8438fbe216887bbe89a12135a768b924",
		},
		{
			std.ED25519,
			"7d13c0248b891b47eb524f2692008e2f97b199bac426cb5902b9003a29ded6ea",
			"59a976ab2c01a81a91f1a56c75ccc77a9e1e9e878e9fe9c3952080a6805b20d5",
			"0af367956af630266b1cc760154256ed79da960dddca9d72a1e8cf27d8d43a77",
			"21108d900134d3b3708dd28ace96b0b23dda9100e4b6a62a8131bd2f2ba408c5",
		},
		{
			std.ZCASH_JUBJUB,
			"4b44cd10d4815a227d627b7120a2e5714e6ad9ad487fae8a32b5058a6399d20d",
			"1ff8fe69e5eac105b01d71ec953d0e574e0842bfc9b69c9a0a8c0735f62e71af",
			"23b8343831d2c81d5eb18dc8099ac3d3a4dcc09ead647293ea6232f5acb82224",
			"67798678674e6695b78266382521efb812230b220f4850c4c698787f0433f482",
		},
		{
			std.FOURQ,
			"55d47d46b8b0c7e2b450d3a4dafbad78 + 1d38c62c6fca0341f98ebff17b2f8276 * i",
			"6e88f539f648cee4f6a8775b6f8724e0 + 7508ec836aa7da83571c3c087bf147f6 * i",
			"65c57caf132541229e5e64c02965241a + 6f303fa0001f7711194a0e062a2d2d5c * i",
			"73ff9338ad602fad8d4bfd84fd7c4890 + 6bfa2793735c57c82845e035abae6309 * i",
		},
	}

	for _, test := range tests {
		fmt.Printf("alg = %s\n", test.alg)
		c, G, err := test.alg.New()
		if err != nil {
			t.Fatalf("new err: %v\n", err.Error())
		}
		sk, _ := new(big.Int).SetString("1000", 16)
		PK := c.ScalarMult(G, sk)

		fmt.Printf("P bitLen = %v\n", c.Field().P().BitLen())
		baseFieldBytes := (c.Field().P().BitLen() + 7) / 8

		// 点的序列化
		var strX, strY string
		x := PK.X().Polynomial()
		strX = cor2Str(x, baseFieldBytes)
		y := PK.Y().Polynomial()
		strY = cor2Str(y, baseFieldBytes)
		fmt.Printf("P.x = %s\n", strX)
		fmt.Printf("P.y = %s\n", strY)
		if strX != test.wantPx || strY != test.wantPy {
			panic("err")
		}

		// 点的反序列化
		F := c.Field()
		P := c.NewPoint(F.Elt(x), F.Elt(y))
		Q := c.ScalarMult(P, sk)

		x = Q.X().Polynomial()
		strX = cor2Str(x, baseFieldBytes)
		fmt.Printf("Q.x = %s\n", strX)
		y = Q.Y().Polynomial()
		strY = cor2Str(y, baseFieldBytes)
		fmt.Printf("Q.y = %s\n", strY)
		if strX != test.wantQx || strY != test.wantQy {
			panic("err")
		}
	}
}

func cor2Str(x []*big.Int, baseFieldBytes int) string {
	strX := hex.EncodeToString(padZero(x[0].Bytes(), baseFieldBytes))
	if len(x) > 1 {
		strX += " + " + hex.EncodeToString(padZero(x[1].Bytes(), baseFieldBytes)) + " * i"
	}
	return strX
}

func padZero(d []byte, lenNeed int) (pd []byte) {
	lenOri := len(d)
	if lenOri <= lenNeed {
		lenPad := lenNeed - lenOri
		pd = d
		for i := 0; i < lenPad; i++ {
			pd = append([]byte{0}, pd...)
		}
	} else {
		panic("padZero lenOri > lenNeed")
	}
	return
}
