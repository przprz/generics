package functions_test

import (
	"generics/functions"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("min() with ordinary types", func() {
	Context("for arguments of the same type", func() {
		It("works", func() {
			Expect(functions.Min(1, 2)).To(Equal(1))
		})
	})

	Context("for signed and unsigned number", func() {
		It("works", func() {
			Expect(functions.Min(1, -1)).To(Equal(-1))
		})
	})

	Context("when different type passed as one of arguments", func() {
		Context("when unexpected type is a untyped constant", func() {
			It("works", func() {
				Expect(functions.Min(1, 2.0)).To(Equal(1))
			})
		})

		Context("when unexpected type is a variable", func() {
			It("❌ won't compile", func() {
				//b := 2.0
				//Expect(functions.Min(1, b)).To(Equal(1))
				//
				// Error message:
				// ./functions_test.go:31:29: cannot use 2.1 (untyped float constant) as int value in argument to functions.Min (truncated)
			})
		})

		Context("when unexpected type is cast", func() {
			Context("when argument is a constant", func() {
				It("❌ won't compile", func() {
					//Expect(functions.Min(1, int(2.1))).To(Equal(1))
					//
					// Error message:
					// ./functions_test.go:39:33: cannot convert 2.1 (untyped float constant) to type int
				})
			})

			Context("when argument is a variable", func() {
				It("works, but 👀it truncates the other variable", func() {
					b := 2.1
					Expect(functions.Min(1, int(b))).To(Equal(1))
				})
			})
		})
	})

	Context("when need to operate on other types", func() {
		It("requires a separate function, operating on these types", func() {
			Expect(functions.FloatMin(1.0, 2.0)).To(Equal(1.0))
		})
	})
})

var _ = Describe("min() with generic type parameters", func() {
	Context("for arguments of the same type", func() {
		Context("when passed as constants", func() {
			It("types are inferred", func() {
				Expect(functions.MinGeneric(1, 2)).To(Equal(1))
			})
		})

		Context("when arguments are passed as variables", func() {
			It("types are inferred", func() {
				a := 1
				b := 2
				Expect(functions.MinGeneric(a, b)).To(Equal(1))
			})
		})

		Context("when argument types are declared", func() {
			It("works", func() {
				var a, b float64 = 1, 2
				Expect(functions.MinGeneric(a, b)).To(Equal(1.0))
			})
		})
	})

	Context("when different types passed", func() {
		Context("one variable declared, the other passed as constant", func() {
			It("types are inferred", func() {
				var a = 3.4
				Expect(functions.MinGeneric(a, 2)).To(Equal(2.0))
			})
		})

		Context("when types are not declared", func() {
			It("❌ won't compile", func() {
				//Expect(functions.MinGeneric(1, 2.0)).To(Equal(1.0))
				//
				// Error message:
				// ./functions_test.go:72:35: default type float64 of 2.0 does not match inferred type int for T
			})
		})

		Context("when type is declared", func() {
			Context("when arguments are passed as constants", func() {
				It("types are inferred", func() {
					Expect(functions.MinGeneric[float64](1, 2.0)).To(Equal(1.0))
					Expect(functions.MinGeneric[int](1, 2.0)).To(Equal(1))
				})
			})

			Context("when arguments are passed as variables", func() {
				It("❌ won't compile", func() {
					//t := 1
					//Expect(functions.MinGeneric[float64](t, 2.0)).To(Equal(1.0))
					//
					// Error message:
					// ./functions_test.go:96:43: cannot use t (variable of type int) as float64 value in argument to functions.MinGeneric[float64]
				})
			})
		})

		Context("when declared type doesn't match the returned type", func() {
			It("❌ won't compile", func() {
				//Expect(functions.MinGeneric[int](1, 2.0)).To(Equal(1.0))
				//
				// Error message:
				//   [FAILED] Expected
				//      <int>: 1
				//  to equal
				//      <float64>: 1
				//  In [It] at: /Users/pawel.przeniczny/dev/go/generics/functions/functions_test.go:82
			})
		})

		Context("when different types are declared", func() {
			It("❌ won't compile", func() {
				//Expect(functions.MinGeneric[int, float64](1, 2.0)).To(Equal(1.0))
				//
				// Error message:
				// ./functions_test.go:72:35: default type float64 of 2.0 does not match inferred type int for T
			})
		})

		Context("when types are passed explicitly", func() {
			Context("when arguments are of different types", func() {
				It("❌ won't compile", func() {
					//var a float64 = 1
					//var b float32 = 2
					//Expect(functions.MinGeneric(a, b)).To(Equal(1.0))
					//
					// Error message:
					//	// ./functions_test.go:60:37: type float32 of b does not match inferred type float64 for T
				})

				Context("when argument types are declared", func() {
					It("❌ won't compile", func() {
						//var a float64 = 1
						//var b float32 = 2
						//Expect(functions.MinGeneric[float64, float32](a, b)).To(Equal(1.0))
						//
						// Error message:
						// ./functions_test.go:155:43: got 2 type arguments but want 1
					})
				})

				Context("when unexpected type is cast", func() {
					It("works", func() {
						var a float64 = 1
						var b float32 = 2
						Expect(functions.MinGeneric(a, float64(b))).To(Equal(1.0))
					})
				})
			})
		})

	})
})
