package collections_test

import (
	"math"
	"testing"

	"github.com/joshuatcasey/collections"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDerivations(t *testing.T, context spec.G, it spec.S) {
	Expect := NewWithT(t).Expect

	context("DeriveFunc", func() {
		it("can be used to calculate the max from an array of ints", func() {
			ints := []int{1, 2, 3, 4, 99, math.MinInt32, 1000}

			max := collections.DeriveFunc(ints, func(current, max int) int {
				if current > max {
					return current
				}
				return max
			})

			Expect(max).To(Equal(1000))
		})

		it("can be used to calculate the max from an array of ints, ignoring the int default value", func() {
			ints := []int{-1000000, -1, -10, -100}

			max := collections.DeriveFunc(ints, func(current, max int) int {
				if current > max {
					return current
				}
				return max
			})

			Expect(max).To(Equal(-1))
		})

		it("can be used to calculate the max from an array of structs", func() {
			type StructWithInt struct {
				Int int
			}

			structs := []StructWithInt{
				{Int: 1},
				{Int: 0},
				{Int: math.MinInt32},
				{Int: 99},
			}

			maxStruct := collections.DeriveFunc(structs, func(current, max StructWithInt) StructWithInt {
				if current.Int > max.Int {
					return current
				}
				return max
			})

			Expect(maxStruct.Int).To(Equal(99))
		})

		it("gracefully handles nil string array", func() {
			derived := collections.DeriveFunc(nil, func(current, max string) string {
				return "not empty"
			})

			Expect(derived).To(Equal(""))
		})

		it("gracefully handles nil struct array", func() {
			type StructWithInt struct {
				Int int
			}

			derived := collections.DeriveFunc(nil, func(current, max StructWithInt) StructWithInt {
				return StructWithInt{Int: 999}
			})

			Expect(derived).To(Equal(StructWithInt{Int: 0}))
		})

		it("can be used to sum an array of ints", func() {
			ints := []int{1, 2, 3, 4}

			sum := collections.DeriveFunc(ints, func(current, sum int) int {
				return current + sum
			})

			Expect(sum).To(Equal(10))
		})

		it("can be used to sum an array of strings", func() {
			strings := []string{"a", "b", "c", "d"}

			sum := collections.DeriveFunc(strings, func(current, sum string) string {
				return sum + current
			})

			Expect(sum).To(Equal("abcd"))
		})
	})

	context("Max", func() {
		it("can be used to calculate the max from an array of int", func() {
			ints := []int{1, 2, 3, 4, 99, math.MinInt32, 1000}
			max := collections.Max(ints)
			Expect(max).To(Equal(1000))
		})
	})

	context("MaxParams", func() {
		it("can be used to calculate the max from an array of ints", func() {
			max := collections.MaxParams(1, 2, 3, 4, 99, math.MinInt32, 1000)
			Expect(max).To(Equal(1000))
		})

		it("can be used to calculate the min from an array of strings", func() {
			max := collections.MaxParams("abc", "def", "xyz", "mno")
			Expect(max).To(Equal("xyz"))
		})
	})

	context("Min", func() {
		it("can be used to calculate the min from an array of ints", func() {
			ints := []int{1, 2, 3, 4, 99, -400, 1000}
			min := collections.Min(ints)
			Expect(min).To(Equal(-400))
		})
	})

	context("MinParams", func() {
		it("can be used to calculate the min from an array of ints", func() {
			min := collections.MinParams(1, 2, 3, 4, 99, math.MinInt32, 1000)
			Expect(min).To(Equal(math.MinInt32))
		})

		it("can be used to calculate the min from an array of strings", func() {
			min := collections.MinParams("abc", "def", "xyz", "mno")
			Expect(min).To(Equal("abc"))
		})
	})

	context("Sum", func() {
		it("can be used to calculate the sum from an array of ints", func() {
			ints := []int{-40, 60, 3}
			sum := collections.Sum(ints)
			Expect(sum).To(Equal(23))
		})
	})

	context("SumParams", func() {
		it("can be used to calculate the min from an array of ints", func() {
			sum := collections.SumParams(-40, 60, 3)
			Expect(sum).To(Equal(23))
		})

		it("can be used to calculate the sum from an array of strings", func() {
			sum := collections.SumParams("abc", "def", "xyz", "mno")
			Expect(sum).To(Equal("abcdefxyzmno"))
		})
	})
}
