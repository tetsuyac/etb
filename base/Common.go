package base

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MO = " "
	SU = MO + " "
	DE = SU + " "
	TE = DE + " "
)

var _dr, _tr bool
var _tv, _cv int

func Suite(s string, f func()) bool {
	fmt.Printf("%v%v", SU, s)
	var _sr bool = true
	f()
	fmt.Printf("%v%v%v", SU, s, _sr)
	return _sr
}
func Describe(s string, f func()) bool {
	fmt.Printf("%v%v", DE, s)
	_dr = true
	f()
	fmt.Printf("%v%v:%v", DE, s, _dr)
	return _dr
}
func Test(s string, f func(), vp string) bool {
	_tr = true
	f()
	_dr = _dr && _tr
	fmt.Printf("%v%v = %v%v", TE, s, vp, strconv.Itoa(_cv))
	return _tr
}

type ToBear struct {
	cv int
	//toBe func(int)
}

var T ToBear

func (t ToBear) Expect(cv int) ToBear {
	t.cv = cv // TODO better
	_cv = cv  //
	return t
}

func (t ToBear) ToBe(tv int) {
	_tv = tv
	_tr = !!(_tr && tv == t.cv)
}
func Po(nvo map[int]int) map[string]any {
	var na, va []int
	for k, v := range nvo {
		na = append(na, k)
		va = append(na, v)
	}
	var nx, vx = px(na), px(va)
	var np, vp []string
	for n := range na {
		np = append(np, pd(nx-len(ps(n))))
		//     String np[] = Arrays.stream(na).mapToObj(n -> pd(nx - ps(n).length())).toArray(String[]::new);
	}
	for v := range va {
		vp = append(vp, pd(vx-len(ps(v))))
		//     String vp[] = IntStream.range(0, va.length).mapToObj(v -> pd(vx - ps(va[v]).length())).toArray(String[]::new);
	}
	o := map[string]any{"na": na, "va": va, "np": np, "vp": vp}
	return o
}
func pd(n int) string {
	//  return 0 < n ? " ".repeat(n) : "";
	var r string
	if 0 < n {
		r = strings.Repeat(" ", n)
	} else {
		r = ""
	}
	return r
}
func px(a []int) int {
	//     String sa[] = Arrays.stream(a).mapToObj(String::valueOf).toArray(String[]::new);
	//    Integer la[] = Arrays.stream(sa).map(String::length).toArray(Integer[]::new);
	//    return Collections.max(Arrays.asList(la)).intValue();
	var sa []string
	var la []int
	for _, s := range a {
		sa = append(sa, strconv.Itoa(s))
	}
	for _, l := range sa {
		la = append(la, len(l))
	}
	return max(la)
}
func max(a []int) int {
	var m int = 0
	for _, n := range a {
		if m < n {
			m = n
		}
	}
	return m
}
func ps(n int) string {
	return strconv.Itoa(n)
	//return String.valueOf(n);
}

// https://stackoverflow.com/questions/24453420/how-to-convert-interface-to-int
func I2S[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, v.(E))
	}
	return
}
