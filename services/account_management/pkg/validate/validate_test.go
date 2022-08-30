package validate

import (
	"fmt"
	"testing"
)

func TestPass(t *testing.T) {
	arr := [7]string{
		"z1La/",
		"i@9>fGDORQ|q]?U8i(S2:aFyp6F!|Re$<vBnV@t#+jL[Q0WmBbSVjH+kqk.>X/8*i^*WtKO5#rOPDd{J[_5XG04Pk+.qpUpBkS<DHXPXw7knVZ]j@0m6FH*:6!CcKlc:|$#L:[oD#EA7W20e1N@X[Czo8:uc;zch(hBGT5_b{dYVt#Wxv}k%joq)X.7x$zI53/>jwbpZVU&dYilvg;gj0zybxa{>]Pr?uEkr6yJO6#Mh2six,Z)!(sR/lx8YWHgg",
		"wle//!cwA",
		"qwd12NcqMA",
		"aljw123!#",
		"ALJ!@#561",
		"asdZXC@!#456",
	}
	for _, v := range arr {
		fmt.Println(v, ": ", ValidPassword(v))
	}
}

func TestEmail(t *testing.T) {
	arr := [6]string{
		"avvevm",
		"lvn/dv@psvd.12c",
		"pf.dsc.@123.42.43",
		".cq.vf@svwe.vw.co",
		"fine@mail..co",
		"fine@m.o",
	}
	for _, v := range arr {
		fmt.Println(v, ": ", ValidEmail(v))
	}
}
