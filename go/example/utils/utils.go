package utils

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	minimalCertificateVerifier "github.com/Layr-Labs/teal/go/example/contracts/bindings/MinimalCertificateVerifier"
)

func NewBlsKeyPairPanics(privateKey string) *bls.KeyPair {
	kp, err := bls.NewKeyPairFromString(privateKey)
	if err != nil {
		panic(err)
	}
	return kp
}

func ToBN254G1Points(ps []*bls.G1Point) []minimalCertificateVerifier.BN254G1Point {
	points := make([]minimalCertificateVerifier.BN254G1Point, len(ps))
	for i, p := range ps {
		points[i] = ToBN254G1Point(p)
	}
	return points
}

func ToBN254G1Point(p *bls.G1Point) minimalCertificateVerifier.BN254G1Point {
	return minimalCertificateVerifier.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}

func ToBN254G2Points(ps []*bls.G2Point) []minimalCertificateVerifier.BN254G2Point {
	points := make([]minimalCertificateVerifier.BN254G2Point, len(ps))
	for i, p := range ps {
		points[i] = ToBN254G2Point(p)
	}
	return points
}

func ToBN254G2Point(p *bls.G2Point) minimalCertificateVerifier.BN254G2Point {
	return minimalCertificateVerifier.BN254G2Point{
		X: [2]*big.Int{p.X.A1.BigInt(new(big.Int)), p.X.A0.BigInt(new(big.Int))},
		Y: [2]*big.Int{p.Y.A1.BigInt(new(big.Int)), p.Y.A0.BigInt(new(big.Int))},
	}
}
