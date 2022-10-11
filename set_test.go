package collections_test

import (
	"testing"

	"github.com/joshuatcasey/collections"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testSet(t *testing.T, context spec.G, it spec.S) {
	Expect := NewWithT(t).Expect

	context("without a type", func() {
		it("uses string as the default", func() {
			set := collections.NewSet("a")
			Expect(set.Contains("a")).To(BeTrue())
		})
	})

	context("type string", func() {
		var (
			zero *collections.Set[string]
			one  *collections.Set[string]
			many *collections.Set[string]
		)

		it.Before(func() {
			zero = collections.NewSetOf[string]()
			one = collections.NewSetOf[string]("a")
			many = collections.NewSetOf[string]("a", "b", "c")
		})

		it("IsEmpty", func() {
			Expect(zero.IsEmpty()).To(BeTrue())
			Expect(one.IsEmpty()).To(BeFalse())
			Expect(many.IsEmpty()).To(BeFalse())
		})

		it("Size", func() {
			Expect(zero.Size()).To(Equal(0))
			Expect(one.Size()).To(Equal(1))
			Expect(many.Size() > 1).To(BeTrue())
		})

		it("Contains", func() {
			Expect(zero.Contains("a")).To(BeFalse())
			Expect(zero.Contains("")).To(BeFalse())

			Expect(one.Contains("a")).To(BeTrue())
			Expect(one.Contains("b")).To(BeFalse())
			Expect(one.Contains("a", "b")).To(BeFalse())
			Expect(one.Contains("")).To(BeFalse())

			Expect(many.Contains("a")).To(BeTrue())
			Expect(many.Contains("b")).To(BeTrue())
			Expect(many.Contains("c")).To(BeTrue())
			Expect(many.Contains("a", "b", "c")).To(BeTrue())
			Expect(many.Contains("d")).To(BeFalse())
			Expect(many.Contains("a", "b", "c", "d")).To(BeFalse())
			Expect(many.Contains("")).To(BeFalse())
		})

		it("Remove", func() {
			set := collections.NewSetOf[string]("a", "b", "c", "d", "e")
			set.Remove("b", "d")

			Expect(set.Elements()).To(ConsistOf("a", "c", "e"))
		})

		it("enforces uniqueness", func() {
			set := collections.NewSetOf[string]()

			set.Add("a")
			set.Add("a")
			set.Add("a")

			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Size()).To(Equal(1))
		})

		it("AddAll", func() {
			setA := collections.NewSetOf[string]("a", "b", "c")
			setB := collections.NewSetOf[string]("c", "d", "e")

			setA.AddAll(setB)
			Expect(setA.Elements()).To(ConsistOf("a", "b", "c", "d", "e"))
		})

		it("KeepOnly", func() {
			setA := collections.NewSetOf[string]("a", "b", "c")
			setB := collections.NewSetOf[string]("c", "d", "e")

			setA.KeepOnly(setB)
			Expect(setA.Elements()).To(ConsistOf("c"))
		})

		it("Elements", func() {
			Expect(many.Elements()).To(ConsistOf("a", "b", "c"))
		})
	})

	context("type int", func() {
		var (
			zero *collections.Set[int]
			one  *collections.Set[int]
			many *collections.Set[int]
		)

		it.Before(func() {
			zero = collections.NewSetOf[int]()
			one = collections.NewSetOf[int]()
			many = collections.NewSetOf[int]()

			one.Add(1)

			many.Add(1)
			many.Add(2)
			many.Add(3)
		})

		it("IsEmpty", func() {
			Expect(zero.IsEmpty()).To(BeTrue())
			Expect(one.IsEmpty()).To(BeFalse())
			Expect(many.IsEmpty()).To(BeFalse())
		})

		it("Size", func() {
			Expect(zero.Size()).To(Equal(0))
			Expect(one.Size()).To(Equal(1))
			Expect(many.Size() > 1).To(BeTrue())
		})

		it("Contains", func() {
			Expect(zero.Contains(1)).To(BeFalse())
			Expect(zero.Contains(0)).To(BeFalse())

			Expect(one.Contains(1)).To(BeTrue())
			Expect(one.Contains(2)).To(BeFalse())
			Expect(one.Contains(1, 22)).To(BeFalse())
			Expect(one.Contains(0)).To(BeFalse())

			Expect(many.Contains(1)).To(BeTrue())
			Expect(many.Contains(2)).To(BeTrue())
			Expect(many.Contains(3)).To(BeTrue())
			Expect(many.Contains(1, 2, 3)).To(BeTrue())
			Expect(many.Contains(4)).To(BeFalse())
			Expect(many.Contains(1, 2, 3, 4)).To(BeFalse())
			Expect(many.Contains(0)).To(BeFalse())
		})

		it("Remove", func() {
			set := collections.NewSetOf[int](1, 2, 3, 4, 5)
			set.Remove(2, 4)

			Expect(set.Elements()).To(ConsistOf(1, 3, 5))
		})

		it("enforces uniqueness", func() {
			set := collections.NewSetOf[int]()

			set.Add(1)
			set.Add(1)
			set.Add(1)

			Expect(set.Elements()).To(ConsistOf(1))
		})

		it("AddAll", func() {
			setA := collections.NewSetOf[int](1, 2, 3)
			setB := collections.NewSetOf[int](3, 4, 5)

			setA.AddAll(setB)
			Expect(setA.Elements()).To(ConsistOf(1, 2, 3, 4, 5))
		})

		it("KeepOnly", func() {
			setA := collections.NewSetOf[int]()
			setA.Add(1, 2, 3)

			setB := collections.NewSetOf[int]()
			setB.Add(3, 4, 5)

			setA.KeepOnly(setB)
			Expect(setA.Elements()).To(ConsistOf(3))
		})

		it("Elements", func() {
			Expect(many.Elements()).To(ConsistOf(1, 2, 3))
		})

		it("ForEach", func() {
			accumulator := 0
			many.ForEach(func(i int) {
				accumulator += i
			})

			Expect(accumulator).To(Equal(6))
		})
	})
}
