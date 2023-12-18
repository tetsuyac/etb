package main

import (
	"etb/base"
	"fmt"
)

type testing string

var Tester testing

func (t testing) T(c func(int) int) {
	nvo := map[int]int{10: 23, 100: 2318, 1000: 233168}
	base.Suite("suite01", func() {
		base.Describe("case01", func() {
			m := base.Po(nvo)
			// https://gosamples.dev/interface-to-string/
			// https://stackoverflow.com/questions/27545270/how-to-get-a-value-from-map
			na := base.I2S[int](m["na"].([]any)) //// so far... cannot debug...
			va := base.I2S[int](m["va"].([]any))
			np := base.I2S[string](m["np"].([]any))
			vp := base.I2S[string](m["vp"].([]any))
			for i := 0; i < len(na); i++ {
				s := "sum of all the multiples of 3 or 5 below {0}{1}: {2}{3}"
				_i := i
				base.Test(fmt.Sprintf(s, fmt.Sprint(na[i]), np[i], vp[i], fmt.Sprint(va[i])), func() {
					base.T.Expect(c(na[_i])).ToBe(va[_i])
				}, vp[i])
			}
		})
	})
}
func main() {
	fmt.Println("T001.f()")
}
