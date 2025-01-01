package pwdx

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	log.Println(Pwdx(14).Dir())
}

func Test2(t *testing.T) {
	log.Println(ManyPwdx(1, 2, 3, 14).GetMap())
}

func Test3(t *testing.T) {
	m := ManyPwdx(1, 2, 3, 14)
	for k, v := range m {
		log.Println(k, v.Dir())
	}
}
