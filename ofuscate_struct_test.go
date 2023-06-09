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
		Fifth []string
		DBs    []DB
	}

	D struct {
		DAs []DA
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
				Expect(Ofuscate(input, "CA.First")["CA"]).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(Ofuscate(input, "CA.Second")["CA"]).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(Ofuscate(input, "CA.Third")["CA"]).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(Ofuscate(input, "CA.Fourth")["CA"]).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CB", func() {
				Expect(Ofuscate(input, "CA.CB")["CA"]).
					To(HaveKeyWithValue("CB", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB on third level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(Ofuscate(input, "CA.CB.First")["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(Ofuscate(input, "CA.CB.Second")["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(Ofuscate(input, "CA.CB.Third")["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(Ofuscate(input, "CA.CB.Fourth")["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CC", func() {
				Expect(Ofuscate(input, "CA.CB.CC")["CA"].(map[string]interface{})["CB"]).
					To(HaveKeyWithValue("CC", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB.CC on fourth level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(Ofuscate(input, "CA.CB.CC.First")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"]).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(Ofuscate(input, "CA.CB.CC.Second")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"]).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(Ofuscate(input, "CA.CB.CC.Third")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"]).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(Ofuscate(input, "CA.CB.CC.Fourth")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"]).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CD", func() {
				Expect(Ofuscate(input, "CA.CB.CC.CD")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"]).
					To(HaveKeyWithValue("CD", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB.CC.CD on fifth level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.First")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.Second")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.Third")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.Fourth")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})

			It("Ofuscating CE", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"],
				).
					To(HaveKeyWithValue("CE", "XXX"))
			})
		})

		When("Ofuscating C.CA.CB.CC.CD.CE on sixth level of nesting", func() {
			It("Ofuscating First", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.First")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("First", "XXX"))
			})

			It("Ofuscating Second", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.Second")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("Second", "XXX"))
			})

			It("Ofuscating Third", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.Third")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("Third", "XXX"))
			})

			It("Ofuscating Fourth", func() {
				Expect(
					Ofuscate(input, "CA.CB.CC.CD.CE.Fourth")["CA"].(map[string]interface{})["CB"].(map[string]interface{})["CC"].(map[string]interface{})["CD"].(map[string]interface{})["CE"],
				).
					To(HaveKeyWithValue("Fourth", "XXX"))
			})
		})
	})

	Context("", func() {
		var input D

		BeforeEach(func() {
			input = D{
				DAs: []DA{
					{
						First:  "DA",
						Second: 1,
						Third:  2,
						Fourth: "another",
						Fifth: []string{"DA", "another"},
						DBs:    []DB{
							{
								First:  "DB",
								Second: 2,
								Third:  3,
								Fourth: nil,
								DCs:    []DC{
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
								DCs:    []DC{
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
						Fifth: []string{"DA", "another"},
						DBs:    []DB{
							{
								First:  "DB",
								Second: 2,
								Third:  3,
								Fourth: nil,
								DCs:    []DC{
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
								DCs:    []DC{
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

			When("Ofuscating first element", func () {
				It("Ofuscating the entire array", func() {
					Expect(Ofuscate(input, "DAs[0].DBs")["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the entire array with brackets", func() {
					Expect(Ofuscate(input, "DAs[0].DBs[]")["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the first element of the arr", func() {
					Expect(Ofuscate(input, "DAs[0].DBs[0]")["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{"XXX", input.DAs[0].DBs[0]}))
				})

				It("Ofuscating the second element of the arr", func() {
					Expect(Ofuscate(input, "DAs[0].DBs[1]")["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{input.DAs[0].DBs[1], "XXX"}))
				})

				It("Ofuscating with index out of range is not a problem", func() {
					Expect(Ofuscate(input, "DAs[0].DBs[9]")["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[0].DBs))
				})

				It("Ofuscating with index under 0 is not a problem either", func() {
					Expect(Ofuscate(input, "DAs[0].DBs[-1]")["DAs"].([]any)[0].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[0].DBs))
				})
			})

			When("Ofuscating second element", func () {
				It("Ofuscating the entire array", func() {
					Expect(Ofuscate(input, "DAs[1].DBs")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the entire array with brackets", func() {
					Expect(Ofuscate(input, "DAs[1].DBs[]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the first element of the arr", func() {
					Expect(Ofuscate(input, "DAs[1].DBs[0]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{"XXX", input.DAs[1].DBs[0]}))
				})

				It("Ofuscating the second element of the arr", func() {
					Expect(Ofuscate(input, "DAs[1].DBs[1]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{input.DAs[1].DBs[1], "XXX"}))
				})

				It("Ofuscating with index out of range is not a problem", func() {
					Expect(Ofuscate(input, "DAs[1].DBs[9]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})

				It("Ofuscating with index under 0 is not a problem either", func() {
					Expect(Ofuscate(input, "DAs[1].DBs[-1]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})
			})

			When("Ofuscating every element", func () {
				It("Ofuscating the entire array", func() {
					Expect(Ofuscate(input, "DAs[].DBs")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the entire array with brackets", func() {
					Expect(Ofuscate(input, "DAs[].DBs[]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", "XXX"))
				})

				It("Ofuscating the first element of the arr", func() {
					Expect(Ofuscate(input, "DAs[].DBs[0]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{"XXX", input.DAs[1].DBs[0]}))
				})

				It("Ofuscating the second element of the arr", func() {
					Expect(Ofuscate(input, "DAs[].DBs[1]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", []any{input.DAs[1].DBs[1], "XXX"}))
				})

				It("Ofuscating with index out of range is not a problem", func() {
					Expect(Ofuscate(input, "DAs[].DBs[9]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})

				It("Ofuscating with index under 0 is not a problem either", func() {
					Expect(Ofuscate(input, "DAs[].DBs[-1]")["DAs"].([]any)[1].(map[string]interface{})).
						To(HaveKeyWithValue("DBs", input.DAs[1].DBs))
				})
			})
		})
	})

})
