package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println("math.Pi =>", math.Pi)
	// 黄金比
	fmt.Println("math.Phi =>", math.Phi)
	// ネイピア数
	fmt.Println("math.E =>", math.E)
	// 2の平方根
	fmt.Println("math.Sqrt2 =>", math.Sqrt2) // 2しか無理
	fmt.Println("\n==================================================================\n")

	// Float の最小・最大
	fmt.Println("math.MaxFloat32 =>", math.MaxFloat32)
	fmt.Println("math.SmallestNonzeroFloat32 =>", math.SmallestNonzeroFloat32) // 0ではない最小値。`MinFloat32` は存在しない
	fmt.Println("math.MaxFloat64 =>", math.MaxFloat64)
	fmt.Println("math.SmallestNonzeroFloat64 =>", math.SmallestNonzeroFloat64) // 0ではない最小値。`MinFloat64` は存在しない
	fmt.Println("\n==================================================================\n")

	// Intの最大・最小
	fmt.Println("math.MaxInt8 =>", math.MaxInt8)
	fmt.Println("math.MinInt8 =>", math.MinInt8)
	fmt.Println("math.MaxInt16 =>", math.MaxInt16)
	fmt.Println("math.MinInt16 =>", math.MinInt16)
	fmt.Println("math.MaxInt32 =>", math.MaxInt32)
	fmt.Println("math.MinInt32 =>", math.MinInt32)
	fmt.Println("math.MaxInt64 =>", math.MaxInt64)
	fmt.Println("math.MinInt64 =>", math.MinInt64)
	fmt.Println("\n==================================================================\n")

	// Uint の最小・最大
	fmt.Println("math.MaxUint8 =>", math.MaxUint8)
	fmt.Println("math.MaxUint16 =>", math.MaxUint16)
	fmt.Println("math.MaxUint32 =>", math.MaxUint32)

	// `math` に定義されている `Max〇〇` とかは constant であり、型が存在しない。
	// `Println` する際に、文脈に合わせて型が推測される。数値の型推論は `Int` になる。`Uint` に型推論されることはない。
	// `math.MaxUint64` の値は `Int` の最大値を超えているから `overflow` のエラーになる。
	// なので、`math.MaxUint64` をそのまま使うことはできず、`unit` にキャストする必要がある。
	fmt.Println("uint64(math.MaxUint64) =>", uint64(math.MaxUint64))
	fmt.Println("\n==================================================================\n")

	// 累乗（Power）
	fmt.Println("math.Pow(0, 10) =>", math.Pow(0, 10))   // 0の10乗
	fmt.Println("math.Pow(1, 100) =>", math.Pow(1, 100)) // 1の100乗
	fmt.Println("math.Pow(2, 8) =>", math.Pow(2, 8))     // 2の8乗
	fmt.Println("\n==================================================================\n")

	// 絶対値（0からの距離なのでマイナスもプラスも関係ない）
	fmt.Println("math.Abs(3.14) =>", math.Abs(3.14))
	fmt.Println("math.Abs(-3.14) =>", math.Abs(-3.14))
	fmt.Println("\n==================================================================\n")

	// 平方根・立方根
	fmt.Println("math.Sqrt(2) =>", math.Sqrt(2))
	fmt.Println("math.Sqrt(10_000) =>", math.Sqrt(10_000))
	fmt.Println("math.Cbrt(8) =>", math.Cbrt(8))
	fmt.Println("\n==================================================================\n")

	// 最小・最大を取得
	fmt.Println("math.Max(100, 101) =>", math.Max(100, 101))
	fmt.Println("math.Min(-100, -101) =>", math.Min(-100, -101))
	fmt.Println("\n==================================================================\n")

	// Trunc（小数点の切り捨て）
	fmt.Println("math.Trunc(3.14) =>", math.Trunc(3.14))
	fmt.Println("math.Trunc(-3.14) =>", math.Trunc(-3.14))
	fmt.Println("\n==================================================================\n")

	// Floor, Ceil
	fmt.Println("math.Floor(1.5) =>", math.Floor(1.5))
	fmt.Println("math.Floor(-1.5) =>", math.Floor(-1.5))
	fmt.Println("math.Ceil(1.5) =>", math.Ceil(1.5))
	fmt.Println("math.Ceil(-1.5) =>", math.Ceil(-1.5))
	fmt.Println("\n==================================================================\n")

	// NaN
	n := math.Sqrt(-1)
	fmt.Println("n =>", n)
	fmt.Println("math.IsNaN(n) =>", math.IsNaN(n))
	fmt.Println("\n==================================================================\n")

	// Infinity
	fmt.Println("math.Inf(0) =>", math.Inf(0))
	fmt.Println("math.Inf(-1) =>", math.Inf(-1))
	fmt.Println("math.NaN() =>", math.NaN())
	fmt.Println("\n==================================================================\n")
}
