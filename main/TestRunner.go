package main

import (
	"etb/base"
	"fmt"
	"os"
	"plugin"
)

func main() {
	TestRunner()
}

func TestRunner() {
	var ta = []int{1 /*1,2,3 , 51*/}
	var sa = []string{"C", "T"}
	var spec = map[string][]string{"euler-go": sa}

	//static Instant start = null;

	for _, t := range ta {
		tn := fmt.Sprintf("%03d", t)
		//tn = strings.Replace(tn, "0", " ", -1)
		fmt.Printf("test%v\n", tn)
		for k, v := range spec {
			tcp := []string{
				fmt.Sprintf("./code/%s%s.so", v[0], tn),
				fmt.Sprintf("./test/%s%s.so", v[1], tn)}
			fmt.Printf("%s%s%s\n", base.MO, k, tn)
			cto := So(t, tcp)
			co := cto[IC].(Coder)
			to := cto[IT].(Tester)
			//start = Instant.now();
			to.T(co.C)
			//Instant finish = Instant.now();
			fmt.Printf("{0}{1}{2}: {3,number,#.##} ms\n", base.MO, k, tn, 10 /*Duration.between(start, finish).toNanos() / 1000000.0)*/)
		}
		fmt.Printf("test{0}\n", tn)
	}
}

type Coder interface {
	C(int) int
}

type Tester interface {
	T(func(int) int)
}

var (
	IC = 0
	IT = 1
) // code, test index

func So(n int, tcp []string) []interface{} {
	na := []string{"Coder", "Tester"} // name array
	//var ma []string                   // module array
	var ra []interface{} // func array
	tid := fmt.Sprintf("%03d", n)
	switch tid {
	case "001":
		//ma = []string{tcp[IC] + ".so", tcp[IT] + ".so"}
	default:
		fmt.Printf("no test set found.\n")
		os.Exit(1)
	}
	for i, m := range tcp {
		pn, err := plugin.Open(m) // plugin name
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sn, err := pn.Lookup(na[i]) // symbol name
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch i {
		case IC:
			var fn Coder
			fn, ok := sn.(Coder)
			if !ok {
				fmt.Printf("unexpected type from module symbol\n")
				os.Exit(1)
			}
			ra = append(ra, fn)
		case IT:
			var fn Tester
			fn, ok := sn.(Tester)
			if !ok {
				fmt.Printf("unexpected type from module symbol\n")
				os.Exit(1)
			}
			ra = append(ra, fn)
		}
	}
	return ra
}
