package ofuscatestruct_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/imartinezalberte/ofuscate-struct"
)

type (
	i8  = int8
	u8  = uint8
	i16 = int16
	u16 = uint16
	i32 = int32
	u32 = uint32
	i64 = int64
	u64 = uint64
	f32 = float32
	f64 = float64

	A struct {
		First, Second string
		Third         int
		Fourth        i8
		Fifth         u8
		Sixth         i16
		Seventh       u16
		Eighth        i32
		Ninth         u32
		Tenth         i64
		Eleventh      u64
		Twelfth       uint
		Thirteenth    f32
		Fourteenth    f64
		Fifteenth     any
	}

	B struct {
		First      []string
		Second     []int
		Third      []A
		Fourth     []i8
		Fifth      []u8
		Sixth      []i16
		Seventh    []u16
		Eighth     []i32
		Ninth      []u32
		Tenth      []i64
		Eleventh   []u64
		Twelfth    []uint
		Thirteenth []f32
		Fourteenth []f64
		Fifteenth  []any
	}

	CE struct {
		First  string
		Second int
		Third  f32
		Fourth any
	}

	// These structures belongs to the same test context, the mega nesting structure
	CD struct {
		First  string
		Second int
		Third  f32
		Fourth any
		CE     CE
	}

	CC struct {
		First  string
		Second int
		Third  f32
		Fourth any
		CD     CD
	}

	CB struct {
		First  string
		Second int
		Third  f32
		Fourth any
		CC     CC
	}

	CA struct {
		First  string
		Second int
		Third  f32
		Fourth any
		CB     CB
	}

	C struct {
		First  string
		Second int
		Third  f32
		Fourth any
		CA     CA
	}

	// These structures belongs to the same test context, the mega nesting structure using arrays
	DC struct {
		First  string
		Second int
		Third  f32
		Fourth any
	}

	DB struct {
		First  string
		Second int
		Third  f32
		Fourth any
		DCs    []DC
	}

	DA struct {
		First  string
		Second int
		Third  f32
		Fourth any
		Fifth  []string
		DBs    []DB
	}

	D struct {
		DAs []DA
	}

	// Simple struct to ofuscate using unidimensional and bidimensional arrays/slices
	E struct {
		First  string
		Second string
		Third  string
	}
)

var _ = Describe("OfuscateStruct", func() {
	Context("case with A being only standard strings and one int", func() {
		var input A

		BeforeEach(func() {
			input = A{
				"hello",
				"world",
				1,
				i8(2),
				u8(3),
				i16(4),
				u16(5),
				i32(6),
				u32(7),
				i64(8),
				u64(9),
				uint(10),
				f32(11.5),
				f64(12.5),
				"other stuff",
			}
		})

		It("ofuscating First attribute which is string", func() {
			Expect(Ofuscate(input, "First")).To(HaveKeyWithValue("First", "XXX"))
		})

		It("ofuscating Second attribute which is string", func() {
			Expect(Ofuscate(input, "Second")).To(HaveKeyWithValue("Second", "XXX"))
		})

		It("ofuscasting Third attribute which is int", func() {
			Expect(Ofuscate(input, "Third")).To(HaveKeyWithValue("Third", "XXX"))
		})

		It("ofuscasting Fourth attribute which is i8", func() {
			Expect(Ofuscate(input, "Fourth")).To(HaveKeyWithValue("Fourth", "XXX"))
		})

		It("ofuscasting Fifth attribute which is u8", func() {
			Expect(Ofuscate(input, "Fifth")).To(HaveKeyWithValue("Fifth", "XXX"))
		})

		It("ofuscasting Sixth attribute which is i16", func() {
			Expect(Ofuscate(input, "Sixth")).To(HaveKeyWithValue("Sixth", "XXX"))
		})

		It("ofuscasting Seventh attribute which is u16", func() {
			Expect(Ofuscate(input, "Seventh")).To(HaveKeyWithValue("Seventh", "XXX"))
		})

		It("ofuscasting Eighth attribute which is i32", func() {
			Expect(Ofuscate(input, "Eighth")).To(HaveKeyWithValue("Eighth", "XXX"))
		})

		It("ofuscasting Ninth attribute which is u32", func() {
			Expect(Ofuscate(input, "Ninth")).To(HaveKeyWithValue("Ninth", "XXX"))
		})

		It("ofuscasting Tenth attribute which is i64", func() {
			Expect(Ofuscate(input, "Tenth")).To(HaveKeyWithValue("Tenth", "XXX"))
		})

		It("ofuscasting Eleventh attribute which is u64", func() {
			Expect(Ofuscate(input, "Eleventh")).To(HaveKeyWithValue("Eleventh", "XXX"))
		})

		It("ofuscasting Twelfth attribute which is uint", func() {
			Expect(Ofuscate(input, "Twelfth")).To(HaveKeyWithValue("Twelfth", "XXX"))
		})

		It("ofuscasting Thirteenth attribute which is f32", func() {
			Expect(Ofuscate(input, "Thirteenth")).To(HaveKeyWithValue("Thirteenth", "XXX"))
		})

		It("ofuscasting Fourteenth attribute which is f64", func() {
			Expect(Ofuscate(input, "Fourteenth")).To(HaveKeyWithValue("Fourteenth", "XXX"))
		})

		It("ofuscasting Fifth attribute which is any", func() {
			Expect(Ofuscate(input, "Fifth")).To(HaveKeyWithValue("Fifth", "XXX"))
		})
	})

	Context("Testing arrays", func() {
		var input B

		BeforeEach(func() {
			input = B{
				[]string{"hello", "world"},
				[]int{1, 2, 4},
				[]A{{}, {}},
				[]i8{2, 3},
				[]u8{3, 4},
				[]i16{4, 5},
				[]u16{5, 6},
				[]i32{6, 7},
				[]u32{7, 8},
				[]i64{8, 9},
				[]u64{9, 10},
				[]uint{10, 11},
				[]f32{11, 12},
				[]f64{12, 13},
				[]any{13, "hello"},
			}
		})

		When("array of strings is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "First")).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "First[]")).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "First[0]")).
					To(HaveKeyWithValue("First", []string{"XXX", input.First[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "First[1]")).
					To(HaveKeyWithValue("First", []string{input.First[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "First[9]")).
					To(HaveKeyWithValue("First", input.First))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "First[-1]")).
					To(HaveKeyWithValue("First", input.First))
			})
		})

		When("array of ints is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Second")).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Second[]")).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Second[0]")).
					To(HaveKeyWithValue("Second", []any{"XXX", input.Second[1], input.Second[2]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Second[1]")).
					To(HaveKeyWithValue("Second", []any{input.Second[0], "XXX", input.Second[2]}))
			})

			It("Ofuscating the third element of the arr", func() {
				Expect(Ofuscate(input, "Second[2]")).
					To(HaveKeyWithValue("Second", []any{input.Second[0], input.Second[1], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Second[9]")).
					To(HaveKeyWithValue("Second", input.Second))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Second[-1]")).
					To(HaveKeyWithValue("Second", input.Second))
			})
		})

		When("array of i8s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Fourth")).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Fourth[]")).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Fourth[0]")).
					To(HaveKeyWithValue("Fourth", []any{"XXX", input.Fourth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Fourth[1]")).
					To(HaveKeyWithValue("Fourth", []any{input.Fourth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Fourth[9]")).
					To(HaveKeyWithValue("Fourth", input.Fourth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Fourth[-1]")).
					To(HaveKeyWithValue("Fourth", input.Fourth))
			})
		})

		When("array of u8s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Fifth")).
					To(HaveKeyWithValue("Fifth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Fifth[]")).
					To(HaveKeyWithValue("Fifth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Fifth[0]")).
					To(HaveKeyWithValue("Fifth", []any{"XXX", input.Fifth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Fifth[1]")).
					To(HaveKeyWithValue("Fifth", []any{input.Fifth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Fifth[9]")).
					To(HaveKeyWithValue("Fifth", input.Fifth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Fifth[-1]")).
					To(HaveKeyWithValue("Fifth", input.Fifth))
			})
		})

		When("array of i16s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Sixth")).
					To(HaveKeyWithValue("Sixth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Sixth[]")).
					To(HaveKeyWithValue("Sixth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Sixth[0]")).
					To(HaveKeyWithValue("Sixth", []any{"XXX", input.Sixth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Sixth[1]")).
					To(HaveKeyWithValue("Sixth", []any{input.Sixth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Sixth[9]")).
					To(HaveKeyWithValue("Sixth", input.Sixth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Sixth[-1]")).
					To(HaveKeyWithValue("Sixth", input.Sixth))
			})
		})

		When("array of u16s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Seventh")).
					To(HaveKeyWithValue("Seventh", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Seventh[]")).
					To(HaveKeyWithValue("Seventh", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Seventh[0]")).
					To(HaveKeyWithValue("Seventh", []any{"XXX", input.Seventh[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Seventh[1]")).
					To(HaveKeyWithValue("Seventh", []any{input.Seventh[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Seventh[9]")).
					To(HaveKeyWithValue("Seventh", input.Seventh))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Seventh[-1]")).
					To(HaveKeyWithValue("Seventh", input.Seventh))
			})
		})

		When("array of i32s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Eighth")).
					To(HaveKeyWithValue("Eighth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Eighth[]")).
					To(HaveKeyWithValue("Eighth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Eighth[0]")).
					To(HaveKeyWithValue("Eighth", []any{"XXX", input.Eighth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Eighth[1]")).
					To(HaveKeyWithValue("Eighth", []any{input.Eighth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Eighth[9]")).
					To(HaveKeyWithValue("Eighth", input.Eighth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Eighth[-1]")).
					To(HaveKeyWithValue("Eighth", input.Eighth))
			})
		})

		When("array of u32s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Ninth")).
					To(HaveKeyWithValue("Ninth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Ninth[]")).
					To(HaveKeyWithValue("Ninth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Ninth[0]")).
					To(HaveKeyWithValue("Ninth", []any{"XXX", input.Ninth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Ninth[1]")).
					To(HaveKeyWithValue("Ninth", []any{input.Ninth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Ninth[9]")).
					To(HaveKeyWithValue("Ninth", input.Ninth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Ninth[-1]")).
					To(HaveKeyWithValue("Ninth", input.Ninth))
			})
		})

		When("array of i64s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Tenth")).
					To(HaveKeyWithValue("Tenth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Tenth[]")).
					To(HaveKeyWithValue("Tenth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Tenth[0]")).
					To(HaveKeyWithValue("Tenth", []any{"XXX", input.Tenth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Tenth[1]")).
					To(HaveKeyWithValue("Tenth", []any{input.Tenth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Tenth[9]")).
					To(HaveKeyWithValue("Tenth", input.Tenth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Tenth[-1]")).
					To(HaveKeyWithValue("Tenth", input.Tenth))
			})
		})

		When("array of u64s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Eleventh")).
					To(HaveKeyWithValue("Eleventh", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Eleventh[]")).
					To(HaveKeyWithValue("Eleventh", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Eleventh[0]")).
					To(HaveKeyWithValue("Eleventh", []any{"XXX", input.Eleventh[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Eleventh[1]")).
					To(HaveKeyWithValue("Eleventh", []any{input.Eleventh[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Eleventh[9]")).
					To(HaveKeyWithValue("Eleventh", input.Eleventh))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Eleventh[-1]")).
					To(HaveKeyWithValue("Eleventh", input.Eleventh))
			})
		})

		When("array of uints is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Twelfth")).
					To(HaveKeyWithValue("Twelfth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Twelfth[]")).
					To(HaveKeyWithValue("Twelfth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Twelfth[0]")).
					To(HaveKeyWithValue("Twelfth", []any{"XXX", input.Twelfth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Twelfth[1]")).
					To(HaveKeyWithValue("Twelfth", []any{input.Twelfth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Twelfth[9]")).
					To(HaveKeyWithValue("Twelfth", input.Twelfth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Twelfth[-1]")).
					To(HaveKeyWithValue("Twelfth", input.Twelfth))
			})
		})

		When("array of f32s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Thirteenth")).
					To(HaveKeyWithValue("Thirteenth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Thirteenth[]")).
					To(HaveKeyWithValue("Thirteenth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Thirteenth[0]")).
					To(HaveKeyWithValue("Thirteenth", []any{"XXX", input.Thirteenth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Thirteenth[1]")).
					To(HaveKeyWithValue("Thirteenth", []any{input.Thirteenth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Thirteenth[9]")).
					To(HaveKeyWithValue("Thirteenth", input.Thirteenth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Thirteenth[-1]")).
					To(HaveKeyWithValue("Thirteenth", input.Thirteenth))
			})
		})

		When("array of f64s is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Fourteenth")).
					To(HaveKeyWithValue("Fourteenth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Fourteenth[]")).
					To(HaveKeyWithValue("Fourteenth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Fourteenth[0]")).
					To(HaveKeyWithValue("Fourteenth", []any{"XXX", input.Fourteenth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Fourteenth[1]")).
					To(HaveKeyWithValue("Fourteenth", []any{input.Fourteenth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Fourteenth[9]")).
					To(HaveKeyWithValue("Fourteenth", input.Fourteenth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Fourteenth[-1]")).
					To(HaveKeyWithValue("Fourteenth", input.Fourteenth))
			})
		})

		When("array of anys is ofuscated", func() {
			It("Ofuscating the entire array", func() {
				Expect(Ofuscate(input, "Fifteenth")).
					To(HaveKeyWithValue("Fifteenth", "XXX"))
			})

			It("Ofuscating the entire array with brackets", func() {
				Expect(Ofuscate(input, "Fifteenth[]")).
					To(HaveKeyWithValue("Fifteenth", "XXX"))
			})

			It("Ofuscating the first element of the arr", func() {
				Expect(Ofuscate(input, "Fifteenth[0]")).
					To(HaveKeyWithValue("Fifteenth", []any{"XXX", input.Fifteenth[1]}))
			})

			It("Ofuscating the second element of the arr", func() {
				Expect(Ofuscate(input, "Fifteenth[1]")).
					To(HaveKeyWithValue("Fifteenth", []any{input.Fifteenth[0], "XXX"}))
			})

			It("Ofuscating with index out of range is not a problem", func() {
				Expect(Ofuscate(input, "Fifteenth[9]")).
					To(HaveKeyWithValue("Fifteenth", input.Fifteenth))
			})

			It("Ofuscating with index under 0 is not a problem either", func() {
				Expect(Ofuscate(input, "Fifteenth[-1]")).
					To(HaveKeyWithValue("Fifteenth", input.Fifteenth))
			})
		})
	})

	Context("Down the nesting hell without using arrays", func() {
		var input C

		BeforeEach(func() {
			input = C{
				First:  "C",
				Second: 1,
				Third:  2,
				Fourth: "hello",
				CA: CA{
					First:  "CA",
					Second: 3,
					Third:  4,
					Fourth: 2.3,
					CB: CB{
						First:  "CB",
						Second: 5,
						Third:  6,
						Fourth: 1,
						CC: CC{
							First:  "CC",
							Second: 7,
							Third:  8,
							Fourth: struct {
								A string
							}{A: "hello"},
							CD: CD{
								First:  "CD",
								Second: 9,
								Third:  10,
								Fourth: map[string]string{"First": "Hello", "Second": "Bye"},
								CE: CE{
									First:  "CE",
									Second: 11,
									Third:  12,
									Fourth: "Bye",
								},
							},
						},
					},
				},
			}
		})

		When("Ofuscating C on the very beginning", func() {
			It("Ofuscating First", func() {
				Expect(Ofuscate(input, "First")).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(Ofuscate(input, "Second")).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(Ofuscate(input, "Third")).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(Ofuscate(input, "Fourth")).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CA", func() {
				Expect(Ofuscate(input, "CA")).
					To(HaveKeyWithValue("CA", "XXX"))
			})
		})

		When("Ofuscating C.CA on second level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(Ofuscate(input, "CA.First").(map[string]interface{})["CA"]).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(Ofuscate(input, "CA.Second").(map[string]interface{})["CA"]).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(Ofuscate(input, "CA.Third").(map[string]interface{})["CA"]).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(Ofuscate(input, "CA.Fourth").(map[string]interface{})["CA"]).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CB", func() {
				Expect(Ofuscate(input, "CA.CB").(map[string]interface{})["CA"]).
					To(HaveKeyWithValue("CB", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB on third level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(Ofuscate(input, "CA.CB.First").(map[string]interface{})["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(Ofuscate(input, "CA.CB.Second").(map[string]interface{})["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(Ofuscate(input, "CA.CB.Third").(map[string]interface{})["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(Ofuscate(input, "CA.CB.Fourth").(map[string]interface{})["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CC", func() {
				Expect(Ofuscate(input, "CA.CB.CC").(map[string]interface{})["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("CC", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB.CC on fourth level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.First").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"],
				).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.Second").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"],
				).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.Third").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"],
				).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.Fourth").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"],
				).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CD", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"],
				).
					To(HaveKeyWithValue("CD", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB.CC.CD on fifth level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.First").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.Second").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.Third").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.Fourth").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CE", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("CE", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB.CC.CD.CE on sixth level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.First").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.Second").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.Third").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.Fourth").(map[string]interface{})["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})
		})
	})

	Context("playing with arrays in depth level", func() {
		var input D

		BeforeEach(func() {
			input = D{
				DAs: []DA{
					{
						First:  "DA",
						Second: 1,
						Third:  2,
						Fourth: "another",
						Fifth:  []string{"DA", "another"},
						DBs: []DB{
							{
								First:  "DB",
								Second: 2,
								Third:  3,
								Fourth: nil,
								DCs: []DC{
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
								},
							},
							{
								First:  "DB",
								Second: 2,
								Third:  3,
								Fourth: nil,
								DCs: []DC{
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
								},
							},
						},
					},
					{
						First:  "DA",
						Second: 1,
						Third:  2,
						Fourth: nil,
						Fifth:  []string{"DA", "another"},
						DBs: []DB{
							{
								First:  "DB",
								Second: 2,
								Third:  3,
								Fourth: nil,
								DCs: []DC{
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
								},
							},
							{
								First:  "DB",
								Second: 2,
								Third:  3,
								Fourth: nil,
								DCs: []DC{
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
									{
										First:  "DC",
										Second: 3,
										Third:  4,
										Fourth: nil,
									},
								},
							},
						},
					},
				},
			}
		})

		When("Ofuscating DA[*].DB", func() {
			When("Ofuscating first element", func() {
				It("Ofuscating the entire array", func() {
					Expect(Ofuscate(input, "DAs[0].DBs").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the entire array with brackets", func() {
					Expect(Ofuscate(input, "DAs[0].DBs[]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the first element of the arr", func() {
					Expect(
						Ofuscate(input, "DAs[0].DBs[0]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", []any{"XXX", input.DAs[0].DBs[0]}))
				})

				It("Ofuscating the second element of the arr", func() {
					Expect(
						Ofuscate(input, "DAs[0].DBs[1]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", []any{input.DAs[0].DBs[1], "XXX"}))
				})

				It("Ofuscating another array of basic types", func() {
					Expect(
						Ofuscate(input, "DAs[0].Fifth[]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{}),
					).
						To(HaveKeyWithValue("Fifth", "XXX"))
				})

				It("Ofuscating an index of an array of basic types", func() {
					Expect(
						Ofuscate(input, "DAs[0].Fifth[0]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{}),
					).
						To(HaveKeyWithValue("Fifth", []string{"XXX", input.DAs[0].Fifth[1]}))
				})

				It("Ofuscating with index out of range is not a problem", func() {
					Expect(
						Ofuscate(input, "DAs[0].DBs[9]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", input.DAs[0].DBs))
				})

				It("Ofuscating with index under 0 is not a problem either", func() {
					Expect(
						Ofuscate(input, "DAs[0].DBs[-1]").(map[string]interface{})["DAs"].([]any)[0].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", input.DAs[0].DBs))
				})
			})

			When("Ofuscating second element", func() {
				It("Ofuscating the entire array", func() {
					Expect(Ofuscate(input, "DAs[1].DBs").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the entire array with brackets", func() {
					Expect(Ofuscate(input, "DAs[1].DBs[]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the first element of the arr", func() {
					Expect(
						Ofuscate(input, "DAs[1].DBs[0]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", []any{"XXX", input.DAs[1].DBs[0]}))
				})

				It("Ofuscating the second element of the arr", func() {
					Expect(
						Ofuscate(input, "DAs[1].DBs[1]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", []any{input.DAs[1].DBs[1], "XXX"}))
				})

				It("Ofuscating with index out of range is not a problem", func() {
					Expect(
						Ofuscate(input, "DAs[1].DBs[9]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})

				It("Ofuscating with index under 0 is not a problem either", func() {
					Expect(
						Ofuscate(input, "DAs[1].DBs[-1]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})
			})

			When("Ofuscating every element", func() {
				It("Ofuscating the entire array", func() {
					Expect(Ofuscate(input, "DAs[].DBs").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the entire array with brackets", func() {
					Expect(Ofuscate(input, "DAs[].DBs[]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the first element of the arr", func() {
					Expect(Ofuscate(input, "DAs[].DBs[0]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{"XXX", input.DAs[1].DBs[0]}))
				})

				It("Ofuscating the second element of the arr", func() {
					Expect(Ofuscate(input, "DAs[].DBs[1]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{input.DAs[1].DBs[1], "XXX"}))
				})

				It("Ofuscating with index out of range is not a problem", func() {
					Expect(Ofuscate(input, "DAs[].DBs[9]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})

				It("Ofuscating with index under 0 is not a problem either", func() {
					Expect(
						Ofuscate(input, "DAs[].DBs[-1]").(map[string]interface{})["DAs"].([]any)[1].(map[string]interface{}),
					).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})
			})
		})
	})

	Context("Ofucasting unidimensional slices", func() {
		var (
			i8s     []i8
			u8s     []u8
			bytes   []byte
			i16s    []i16
			u16s    []u16
			i32s    []i32
			u32s    []u32
			i64s    []i64
			u64s    []u64
			ints    []int
			uints   []uint
			f32s    []f32
			f64s    []f64
			strings []string
			bools   []bool
			structs []E
		)

		BeforeEach(func() {
			i8s = []i8{1, 2, 3, 4}
			u8s = []u8{1, 2, 3, 4}
			bytes = []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f} // hello
			i16s = []i16{1, 2, 3, 4}
			u16s = []u16{1, 2, 3, 4}
			i32s = []i32{1, 2, 3, 4}
			u32s = []u32{1, 2, 3, 4}
			i64s = []i64{1, 2, 3, 4}
			u64s = []u64{1, 2, 3, 4}
			ints = []int{1, 2, 3, 4}
			uints = []uint{1, 2, 3, 4}
			f32s = []f32{1, 2, 3, 4}
			f64s = []f64{1, 2, 3, 4}
			strings = []string{"hello", "world", "bye", "goodbye"}
			bools = []bool{true, false, true, false}
			structs = []E{
				{
					First:  "first in first index",
					Second: "second in first index",
					Third:  "thid in first index",
				},
				{
					First:  "first in second index",
					Second: "second in second index",
					Third:  "thid in second index",
				},
			}
		})

		When("ofuscating i8s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(i8s, "[0]")).To(Equal([]any{"XXX", i8(2), i8(3), i8(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(i8s, "[1]")).To(Equal([]any{i8(1), "XXX", i8(3), i8(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(i8s, "[2]")).To(Equal([]any{i8(1), i8(2), "XXX", i8(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(i8s, "[3]")).To(Equal([]any{i8(1), i8(2), i8(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(i8s, "[4]")).To(Equal([]any{i8(1), i8(2), i8(3), i8(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(i8s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating u8s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(u8s, "[0]")).To(Equal([]any{"XXX", u8(2), u8(3), u8(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(u8s, "[1]")).To(Equal([]any{u8(1), "XXX", u8(3), u8(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(u8s, "[2]")).To(Equal([]any{u8(1), u8(2), "XXX", u8(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(u8s, "[3]")).To(Equal([]any{u8(1), u8(2), u8(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(u8s, "[4]")).To(Equal([]any{u8(1), u8(2), u8(3), u8(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(u8s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating bytes", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(bytes, "[0]")).To(Equal([]any{"XXX", byte(0x65), byte(0x6c), byte(0x6c), byte(0x6f)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(bytes, "[1]")).To(Equal([]any{byte(0x68), "XXX", byte(0x6c), byte(0x6c), byte(0x6f)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(bytes, "[2]")).To(Equal([]any{byte(0x68), byte(0x65), "XXX", byte(0x6c), byte(0x6f)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(bytes, "[3]")).To(Equal([]any{byte(0x68), byte(0x65), byte(0x6c), "XXX", byte(0x6f)}))
			})

			It("Ofuscating Fifth element", func() {
				Expect(Ofuscate(bytes, "[4]")).To(Equal([]any{byte(0x68), byte(0x65), byte(0x6c), byte(0x6c), "XXX"}))
			})

			It("Ofuscating Sixth element, unexitent", func() {
				Expect(Ofuscate(bytes, "[5]")).To(Equal([]any{byte(0x68), byte(0x65), byte(0x6c), byte(0x6c), byte(0x6f)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(bytes, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating i16s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(i16s, "[0]")).To(Equal([]any{"XXX", i16(2), i16(3), i16(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(i16s, "[1]")).To(Equal([]any{i16(1), "XXX", i16(3), i16(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(i16s, "[2]")).To(Equal([]any{i16(1), i16(2), "XXX", i16(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(i16s, "[3]")).To(Equal([]any{i16(1), i16(2), i16(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(i16s, "[4]")).To(Equal([]any{i16(1), i16(2), i16(3), i16(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(i16s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating u16s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(u16s, "[0]")).To(Equal([]any{"XXX", u16(2), u16(3), u16(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(u16s, "[1]")).To(Equal([]any{u16(1), "XXX", u16(3), u16(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(u16s, "[2]")).To(Equal([]any{u16(1), u16(2), "XXX", u16(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(u16s, "[3]")).To(Equal([]any{u16(1), u16(2), u16(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(u16s, "[4]")).To(Equal([]any{u16(1), u16(2), u16(3), u16(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(u16s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating i32s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(i32s, "[0]")).To(Equal([]any{"XXX", i32(2), i32(3), i32(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(i32s, "[1]")).To(Equal([]any{i32(1), "XXX", i32(3), i32(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(i32s, "[2]")).To(Equal([]any{i32(1), i32(2), "XXX", i32(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(i32s, "[3]")).To(Equal([]any{i32(1), i32(2), i32(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(i32s, "[4]")).To(Equal([]any{i32(1), i32(2), i32(3), i32(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(i32s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating u32s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(u32s, "[0]")).To(Equal([]any{"XXX", u32(2), u32(3), u32(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(u32s, "[1]")).To(Equal([]any{u32(1), "XXX", u32(3), u32(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(u32s, "[2]")).To(Equal([]any{u32(1), u32(2), "XXX", u32(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(u32s, "[3]")).To(Equal([]any{u32(1), u32(2), u32(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(u32s, "[4]")).To(Equal([]any{u32(1), u32(2), u32(3), u32(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(u32s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating i64s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(i64s, "[0]")).To(Equal([]any{"XXX", i64(2), i64(3), i64(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(i64s, "[1]")).To(Equal([]any{i64(1), "XXX", i64(3), i64(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(i64s, "[2]")).To(Equal([]any{i64(1), i64(2), "XXX", i64(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(i64s, "[3]")).To(Equal([]any{i64(1), i64(2), i64(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(i64s, "[4]")).To(Equal([]any{i64(1), i64(2), i64(3), i64(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(i64s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating u64s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(u64s, "[0]")).To(Equal([]any{"XXX", u64(2), u64(3), u64(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(u64s, "[1]")).To(Equal([]any{u64(1), "XXX", u64(3), u64(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(u64s, "[2]")).To(Equal([]any{u64(1), u64(2), "XXX", u64(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(u64s, "[3]")).To(Equal([]any{u64(1), u64(2), u64(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(u64s, "[4]")).To(Equal([]any{u64(1), u64(2), u64(3), u64(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(u64s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating ints", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(ints, "[0]")).To(Equal([]any{"XXX", int(2), int(3), int(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(ints, "[1]")).To(Equal([]any{int(1), "XXX", int(3), int(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(ints, "[2]")).To(Equal([]any{int(1), int(2), "XXX", int(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(ints, "[3]")).To(Equal([]any{int(1), int(2), int(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(ints, "[4]")).To(Equal([]any{int(1), int(2), int(3), int(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(ints, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating uints", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(uints, "[0]")).To(Equal([]any{"XXX", uint(2), uint(3), uint(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(uints, "[1]")).To(Equal([]any{uint(1), "XXX", uint(3), uint(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(uints, "[2]")).To(Equal([]any{uint(1), uint(2), "XXX", uint(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(uints, "[3]")).To(Equal([]any{uint(1), uint(2), uint(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(uints, "[4]")).To(Equal([]any{uint(1), uint(2), uint(3), uint(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(uints, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating f32s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(f32s, "[0]")).To(Equal([]any{"XXX", f32(2), f32(3), f32(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(f32s, "[1]")).To(Equal([]any{f32(1), "XXX", f32(3), f32(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(f32s, "[2]")).To(Equal([]any{f32(1), f32(2), "XXX", f32(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(f32s, "[3]")).To(Equal([]any{f32(1), f32(2), f32(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(f32s, "[4]")).To(Equal([]any{f32(1), f32(2), f32(3), f32(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(f32s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating f64s", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(f64s, "[0]")).To(Equal([]any{"XXX", f64(2), f64(3), f64(4)}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(f64s, "[1]")).To(Equal([]any{f64(1), "XXX", f64(3), f64(4)}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(f64s, "[2]")).To(Equal([]any{f64(1), f64(2), "XXX", f64(4)}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(f64s, "[3]")).To(Equal([]any{f64(1), f64(2), f64(3), "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(f64s, "[4]")).To(Equal([]any{f64(1), f64(2), f64(3), f64(4)}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(f64s, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating strings", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(strings, "[0]")).To(Equal([]any{"XXX", "world", "bye", "goodbye"}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(strings, "[1]")).To(Equal([]any{"hello", "XXX", "bye", "goodbye"}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(strings, "[2]")).To(Equal([]any{"hello", "world", "XXX", "goodbye"}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(strings, "[3]")).To(Equal([]any{"hello", "world", "bye", "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(strings, "[4]")).To(Equal([]any{"hello", "world", "bye", "goodbye"}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(strings, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating bools", func() {
			It("Ofuscating first element", func() {
				Expect(Ofuscate(bools, "[0]")).To(Equal([]any{"XXX", false, true, false}))
			})

			It("Ofuscating second element", func() {
				Expect(Ofuscate(bools, "[1]")).To(Equal([]any{true, "XXX", true, false}))
			})

			It("Ofuscating third element", func() {
				Expect(Ofuscate(bools, "[2]")).To(Equal([]any{true, false, "XXX", false}))
			})

			It("Ofuscating fourth element", func() {
				Expect(Ofuscate(bools, "[3]")).To(Equal([]any{true, false, true, "XXX"}))
			})

			It("Ofuscating Fifth element, unexistent", func() {
				Expect(Ofuscate(bools, "[4]")).To(Equal([]any{true, false, true, false}))
			})

			It("Ofuscating everything", func() {
				Expect(Ofuscate(bools, "[]")).To(Equal([]any{"XXX", "XXX", "XXX", "XXX"}))
			})
		})

		When("ofuscating structs", func() {
			When("ofuscating first element", func() {
				It("entirely", func() {
					Expect(Ofuscate(structs, "[0]")).To(Equal([]any{"XXX", structs[1]}))
				})

				It("First attribute", func() {
					Expect(Ofuscate(structs, "[0].First")).To(Equal([]any{map[string]any{
						"First":  "XXX",
						"Second": structs[0].Second,
						"Third":  structs[0].Third,
					}, structs[1]}))
				})

				It("Second attribute", func() {
					Expect(Ofuscate(structs, "[0].Second")).To(Equal([]any{map[string]any{
						"First":  structs[0].First,
						"Second": "XXX",
						"Third":  structs[0].Third,
					}, structs[1]}))
				})

				It("Third attribute", func() {
					Expect(Ofuscate(structs, "[0].Third")).To(Equal([]any{map[string]any{
						"First":  structs[0].First,
						"Second": structs[0].Second,
						"Third":  "XXX",
					}, structs[1]}))
				})
			})

			When("ofuscating seocnd element", func() {
				It("entirely", func() {
					Expect(Ofuscate(structs, "[1]")).To(Equal([]any{structs[0], "XXX"}))
				})

				It("First attribute", func() {
					Expect(Ofuscate(structs, "[1].First")).To(Equal([]any{structs[0], map[string]any{
						"First":  "XXX",
						"Second": structs[1].Second,
						"Third":  structs[1].Third,
					}}))
				})

				It("Second attribute", func() {
					Expect(Ofuscate(structs, "[1].Second")).To(Equal([]any{structs[0], map[string]any{
						"First":  structs[1].First,
						"Second": "XXX",
						"Third":  structs[1].Third,
					}}))
				})

				It("Third attribute", func() {
					Expect(Ofuscate(structs, "[1].Third")).To(Equal([]any{structs[0], map[string]any{
						"First":  structs[1].First,
						"Second": structs[1].Second,
						"Third":  "XXX",
					}}))
				})
			})
		})
	})

	Context("Ofuscasting bidimensional slices", func() {
		var (
			i8s     [][]i8
			u8s     [][]u8
			bytes   [][]byte
			i16s    [][]i16
			u16s    [][]u16
			i32s    [][]i32
			u32s    [][]u32
			i64s    [][]i64
			u64s    [][]u64
			ints    [][]int
			uints   [][]uint
			f32s    [][]f32
			f64s    [][]f64
			strings [][]string
			bools   [][]bool
		)

		BeforeEach(func() {
			i8s = [][]i8{{1, 2, 3, 4}, {1, 2, 3, 4}}
			u8s = [][]u8{{1, 2, 3, 4}, {1, 2, 3, 4}}
			bytes = [][]byte{
				{0x68, 0x65, 0x6c, 0x6c, 0x6f},
				{0x68, 0x65, 0x6c, 0x6c, 0x6f},
			}
			i16s = [][]i16{{1, 2, 3, 4}, {1, 2, 3, 4}}
			u16s = [][]u16{{1, 2, 3, 4}, {1, 2, 3, 4}}
			i32s = [][]i32{{1, 2, 3, 4}, {1, 2, 3, 4}}
			u32s = [][]u32{{1, 2, 3, 4}, {1, 2, 3, 4}}
			i64s = [][]i64{{1, 2, 3, 4}, {1, 2, 3, 4}}
			u64s = [][]u64{{1, 2, 3, 4}, {1, 2, 3, 4}}
			ints = [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}}
			uints = [][]uint{{1, 2, 3, 4}, {1, 2, 3, 4}}
			f32s = [][]f32{{1, 2, 3, 4}, {1, 2, 3, 4}}
			f64s = [][]f64{{1, 2, 3, 4}, {1, 2, 3, 4}}
			strings = [][]string{{"hello", "world", "bye", "goodbye"}, {"hello", "world", "bye", "goodbye"}}
			bools = [][]bool{{true, false, true, false}, {true, false, true, false}}

			_ = strings
			_ = bytes
			_ = bools
		})

		When("ofuscating i8s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", i8(2), i8(3), i8(4)},
						[]i8{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{i8(1), "XXX", i8(3), i8(4)},
						[]i8{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{i8(1), i8(2), "XXX", i8(4)},
						[]i8{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{i8(1), i8(2), i8(3), "XXX"},
						[]i8{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{i8(1), i8(2), i8(3), i8(4)},
						[]i8{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]i8{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]i8{i8(1), i8(2), i8(3), i8(4)},
						[]any{"XXX", i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]i8{i8(1), i8(2), i8(3), i8(4)},
						[]any{i8(1), "XXX", i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]i8{i8(1), i8(2), i8(3), i8(4)},
						[]any{i8(1), i8(2), "XXX", i8(4)},
					}
					Expect(Ofuscate(i8s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]i8{i8(1), i8(2), i8(3), i8(4)},
						[]any{i8(1), i8(2), i8(3), "XXX"},
					}
					Expect(Ofuscate(i8s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]i8{i8(1), i8(2), i8(3), i8(4)},
						[]any{i8(1), i8(2), i8(3), i8(4)},
					}
					Expect(Ofuscate(i8s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]i8{i8(1), i8(2), i8(3), i8(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(i8s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating u8s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", u8(2), u8(3), u8(4)},
						[]u8{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{u8(1), "XXX", u8(3), u8(4)},
						[]u8{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{u8(1), u8(2), "XXX", u8(4)},
						[]u8{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{u8(1), u8(2), u8(3), "XXX"},
						[]u8{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{u8(1), u8(2), u8(3), u8(4)},
						[]u8{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]u8{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]u8{u8(1), u8(2), u8(3), u8(4)},
						[]any{"XXX", u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]u8{u8(1), u8(2), u8(3), u8(4)},
						[]any{u8(1), "XXX", u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]u8{u8(1), u8(2), u8(3), u8(4)},
						[]any{u8(1), u8(2), "XXX", u8(4)},
					}
					Expect(Ofuscate(u8s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]u8{u8(1), u8(2), u8(3), u8(4)},
						[]any{u8(1), u8(2), u8(3), "XXX"},
					}
					Expect(Ofuscate(u8s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]u8{u8(1), u8(2), u8(3), u8(4)},
						[]any{u8(1), u8(2), u8(3), u8(4)},
					}
					Expect(Ofuscate(u8s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]u8{u8(1), u8(2), u8(3), u8(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(u8s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating i16s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", i16(2), i16(3), i16(4)},
						[]i16{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{i16(1), "XXX", i16(3), i16(4)},
						[]i16{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{i16(1), i16(2), "XXX", i16(4)},
						[]i16{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{i16(1), i16(2), i16(3), "XXX"},
						[]i16{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{i16(1), i16(2), i16(3), i16(4)},
						[]i16{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]i16{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]i16{i16(1), i16(2), i16(3), i16(4)},
						[]any{"XXX", i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]i16{i16(1), i16(2), i16(3), i16(4)},
						[]any{i16(1), "XXX", i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]i16{i16(1), i16(2), i16(3), i16(4)},
						[]any{i16(1), i16(2), "XXX", i16(4)},
					}
					Expect(Ofuscate(i16s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]i16{i16(1), i16(2), i16(3), i16(4)},
						[]any{i16(1), i16(2), i16(3), "XXX"},
					}
					Expect(Ofuscate(i16s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]i16{i16(1), i16(2), i16(3), i16(4)},
						[]any{i16(1), i16(2), i16(3), i16(4)},
					}
					Expect(Ofuscate(i16s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]i16{i16(1), i16(2), i16(3), i16(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(i16s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating u16s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", u16(2), u16(3), u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{u16(1), "XXX", u16(3), u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{u16(1), u16(2), "XXX", u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{u16(1), u16(2), u16(3), "XXX"},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{u16(1), u16(2), u16(3), u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{"XXX", u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), "XXX", u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), u16(2), "XXX", u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), u16(2), u16(3), "XXX"},
					}
					Expect(Ofuscate(u16s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(u16s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating u16s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", u16(2), u16(3), u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{u16(1), "XXX", u16(3), u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{u16(1), u16(2), "XXX", u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{u16(1), u16(2), u16(3), "XXX"},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{u16(1), u16(2), u16(3), u16(4)},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]u16{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{"XXX", u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), "XXX", u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), u16(2), "XXX", u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), u16(2), u16(3), "XXX"},
					}
					Expect(Ofuscate(u16s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{u16(1), u16(2), u16(3), u16(4)},
					}
					Expect(Ofuscate(u16s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]u16{u16(1), u16(2), u16(3), u16(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(u16s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating i32s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", i32(2), i32(3), i32(4)},
						[]i32{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{i32(1), "XXX", i32(3), i32(4)},
						[]i32{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{i32(1), i32(2), "XXX", i32(4)},
						[]i32{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{i32(1), i32(2), i32(3), "XXX"},
						[]i32{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{i32(1), i32(2), i32(3), i32(4)},
						[]i32{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]i32{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]i32{i32(1), i32(2), i32(3), i32(4)},
						[]any{"XXX", i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]i32{i32(1), i32(2), i32(3), i32(4)},
						[]any{i32(1), "XXX", i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]i32{i32(1), i32(2), i32(3), i32(4)},
						[]any{i32(1), i32(2), "XXX", i32(4)},
					}
					Expect(Ofuscate(i32s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]i32{i32(1), i32(2), i32(3), i32(4)},
						[]any{i32(1), i32(2), i32(3), "XXX"},
					}
					Expect(Ofuscate(i32s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]i32{i32(1), i32(2), i32(3), i32(4)},
						[]any{i32(1), i32(2), i32(3), i32(4)},
					}
					Expect(Ofuscate(i32s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]i32{i32(1), i32(2), i32(3), i32(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(i32s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating u32s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", u32(2), u32(3), u32(4)},
						[]u32{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{u32(1), "XXX", u32(3), u32(4)},
						[]u32{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{u32(1), u32(2), "XXX", u32(4)},
						[]u32{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{u32(1), u32(2), u32(3), "XXX"},
						[]u32{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{u32(1), u32(2), u32(3), u32(4)},
						[]u32{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]u32{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]u32{u32(1), u32(2), u32(3), u32(4)},
						[]any{"XXX", u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]u32{u32(1), u32(2), u32(3), u32(4)},
						[]any{u32(1), "XXX", u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]u32{u32(1), u32(2), u32(3), u32(4)},
						[]any{u32(1), u32(2), "XXX", u32(4)},
					}
					Expect(Ofuscate(u32s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]u32{u32(1), u32(2), u32(3), u32(4)},
						[]any{u32(1), u32(2), u32(3), "XXX"},
					}
					Expect(Ofuscate(u32s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]u32{u32(1), u32(2), u32(3), u32(4)},
						[]any{u32(1), u32(2), u32(3), u32(4)},
					}
					Expect(Ofuscate(u32s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]u32{u32(1), u32(2), u32(3), u32(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(u32s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating i64s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", i64(2), i64(3), i64(4)},
						[]i64{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{i64(1), "XXX", i64(3), i64(4)},
						[]i64{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{i64(1), i64(2), "XXX", i64(4)},
						[]i64{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{i64(1), i64(2), i64(3), "XXX"},
						[]i64{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{i64(1), i64(2), i64(3), i64(4)},
						[]i64{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]i64{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]i64{i64(1), i64(2), i64(3), i64(4)},
						[]any{"XXX", i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]i64{i64(1), i64(2), i64(3), i64(4)},
						[]any{i64(1), "XXX", i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]i64{i64(1), i64(2), i64(3), i64(4)},
						[]any{i64(1), i64(2), "XXX", i64(4)},
					}
					Expect(Ofuscate(i64s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]i64{i64(1), i64(2), i64(3), i64(4)},
						[]any{i64(1), i64(2), i64(3), "XXX"},
					}
					Expect(Ofuscate(i64s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]i64{i64(1), i64(2), i64(3), i64(4)},
						[]any{i64(1), i64(2), i64(3), i64(4)},
					}
					Expect(Ofuscate(i64s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]i64{i64(1), i64(2), i64(3), i64(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(i64s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating u64s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", u64(2), u64(3), u64(4)},
						[]u64{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{u64(1), "XXX", u64(3), u64(4)},
						[]u64{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{u64(1), u64(2), "XXX", u64(4)},
						[]u64{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{u64(1), u64(2), u64(3), "XXX"},
						[]u64{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{u64(1), u64(2), u64(3), u64(4)},
						[]u64{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]u64{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]u64{u64(1), u64(2), u64(3), u64(4)},
						[]any{"XXX", u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]u64{u64(1), u64(2), u64(3), u64(4)},
						[]any{u64(1), "XXX", u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]u64{u64(1), u64(2), u64(3), u64(4)},
						[]any{u64(1), u64(2), "XXX", u64(4)},
					}
					Expect(Ofuscate(u64s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]u64{u64(1), u64(2), u64(3), u64(4)},
						[]any{u64(1), u64(2), u64(3), "XXX"},
					}
					Expect(Ofuscate(u64s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]u64{u64(1), u64(2), u64(3), u64(4)},
						[]any{u64(1), u64(2), u64(3), u64(4)},
					}
					Expect(Ofuscate(u64s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]u64{u64(1), u64(2), u64(3), u64(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(u64s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating ints", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", int(2), int(3), int(4)},
						[]int{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{int(1), "XXX", int(3), int(4)},
						[]int{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{int(1), int(2), "XXX", int(4)},
						[]int{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{int(1), int(2), int(3), "XXX"},
						[]int{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{int(1), int(2), int(3), int(4)},
						[]int{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]int{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]int{int(1), int(2), int(3), int(4)},
						[]any{"XXX", int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]int{int(1), int(2), int(3), int(4)},
						[]any{int(1), "XXX", int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]int{int(1), int(2), int(3), int(4)},
						[]any{int(1), int(2), "XXX", int(4)},
					}
					Expect(Ofuscate(ints, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]int{int(1), int(2), int(3), int(4)},
						[]any{int(1), int(2), int(3), "XXX"},
					}
					Expect(Ofuscate(ints, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]int{int(1), int(2), int(3), int(4)},
						[]any{int(1), int(2), int(3), int(4)},
					}
					Expect(Ofuscate(ints, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]int{int(1), int(2), int(3), int(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(ints, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating uints", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", uint(2), uint(3), uint(4)},
						[]uint{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{uint(1), "XXX", uint(3), uint(4)},
						[]uint{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{uint(1), uint(2), "XXX", uint(4)},
						[]uint{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{uint(1), uint(2), uint(3), "XXX"},
						[]uint{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{uint(1), uint(2), uint(3), uint(4)},
						[]uint{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]uint{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]uint{uint(1), uint(2), uint(3), uint(4)},
						[]any{"XXX", uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]uint{uint(1), uint(2), uint(3), uint(4)},
						[]any{uint(1), "XXX", uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]uint{uint(1), uint(2), uint(3), uint(4)},
						[]any{uint(1), uint(2), "XXX", uint(4)},
					}
					Expect(Ofuscate(uints, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]uint{uint(1), uint(2), uint(3), uint(4)},
						[]any{uint(1), uint(2), uint(3), "XXX"},
					}
					Expect(Ofuscate(uints, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]uint{uint(1), uint(2), uint(3), uint(4)},
						[]any{uint(1), uint(2), uint(3), uint(4)},
					}
					Expect(Ofuscate(uints, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]uint{uint(1), uint(2), uint(3), uint(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(uints, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating f32s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", f32(2), f32(3), f32(4)},
						[]f32{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{f32(1), "XXX", f32(3), f32(4)},
						[]f32{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{f32(1), f32(2), "XXX", f32(4)},
						[]f32{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{f32(1), f32(2), f32(3), "XXX"},
						[]f32{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{f32(1), f32(2), f32(3), f32(4)},
						[]f32{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]f32{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]f32{f32(1), f32(2), f32(3), f32(4)},
						[]any{"XXX", f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]f32{f32(1), f32(2), f32(3), f32(4)},
						[]any{f32(1), "XXX", f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]f32{f32(1), f32(2), f32(3), f32(4)},
						[]any{f32(1), f32(2), "XXX", f32(4)},
					}
					Expect(Ofuscate(f32s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]f32{f32(1), f32(2), f32(3), f32(4)},
						[]any{f32(1), f32(2), f32(3), "XXX"},
					}
					Expect(Ofuscate(f32s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]f32{f32(1), f32(2), f32(3), f32(4)},
						[]any{f32(1), f32(2), f32(3), f32(4)},
					}
					Expect(Ofuscate(f32s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]f32{f32(1), f32(2), f32(3), f32(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(f32s, "[1][]")).To(Equal(result))
				})
			})
		})

		When("ofuscating f64s", func() {
			When("Ofuscating first parent element", func() {
				It("Ofuscating first child element", func() {
					result := []any{
						[]any{"XXX", f64(2), f64(3), f64(4)},
						[]f64{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[0][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]any{f64(1), "XXX", f64(3), f64(4)},
						[]f64{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[0][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]any{f64(1), f64(2), "XXX", f64(4)},
						[]f64{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[0][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]any{f64(1), f64(2), f64(3), "XXX"},
						[]f64{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[0][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]any{f64(1), f64(2), f64(3), f64(4)},
						[]f64{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[0][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]any{"XXX", "XXX", "XXX", "XXX"},
						[]f64{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[0][]")).To(Equal(result))
				})
			})

			When("Ofuscating second parent element", func() {
				It("Ofuscating second child element", func() {
					result := []any{
						[]f64{f64(1), f64(2), f64(3), f64(4)},
						[]any{"XXX", f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[1][0]")).To(Equal(result))
				})

				It("Ofuscating second child element", func() {
					result := []any{
						[]f64{f64(1), f64(2), f64(3), f64(4)},
						[]any{f64(1), "XXX", f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[1][1]")).To(Equal(result))
				})

				It("Ofuscating third child element", func() {
					result := []any{
						[]f64{f64(1), f64(2), f64(3), f64(4)},
						[]any{f64(1), f64(2), "XXX", f64(4)},
					}
					Expect(Ofuscate(f64s, "[1][2]")).To(Equal(result))
				})

				It("Ofuscating fourth child element", func() {
					result := []any{
						[]f64{f64(1), f64(2), f64(3), f64(4)},
						[]any{f64(1), f64(2), f64(3), "XXX"},
					}
					Expect(Ofuscate(f64s, "[1][3]")).To(Equal(result))
				})

				It("Ofuscating Fifth child element, unexistent", func() {
					result := []any{
						[]f64{f64(1), f64(2), f64(3), f64(4)},
						[]any{f64(1), f64(2), f64(3), f64(4)},
					}
					Expect(Ofuscate(f64s, "[1][4]")).To(Equal(result))
				})

				It("Ofuscating everything", func() {
					result := []any{
						[]f64{f64(1), f64(2), f64(3), f64(4)},
						[]any{"XXX", "XXX", "XXX", "XXX"},
					}
					Expect(Ofuscate(f64s, "[1][]")).To(Equal(result))
				})
			})
		})
	})
})
