package crossChainUtils

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	BLSCertificateVerifier "github.com/Layr-Labs/teal/transporter/bindings/BLSCertificateVerifier"
)

func ToBN254G1Point(p *bls.G1Point) BLSCertificateVerifier.BN254G1Point {
	return BLSCertificateVerifier.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}

func ToBN254G2Point(p *bls.G2Point) BLSCertificateVerifier.BN254G2Point {
	return BLSCertificateVerifier.BN254G2Point{
		X: [2]*big.Int{p.X.A1.BigInt(new(big.Int)), p.X.A0.BigInt(new(big.Int))},
		Y: [2]*big.Int{p.Y.A1.BigInt(new(big.Int)), p.Y.A0.BigInt(new(big.Int))},
	}
}
