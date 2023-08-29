package inviteCodeGen

import (
	"fmt"
	"testing"
)

func TestInviteCode(t *testing.T) {
	inviteCode := code{
		base:    "HVE8S2DZX9C7P5IK3MJUAR4WYLTN6BGQ",
		decimal: 32,
		pad:     "F",
		len:     6,
	}
	// 初始化检查
	if res, err := inviteCode.InitCheck(); !res {
		fmt.Println(err)
		return
	}
	id := int64(1)
	code := inviteCode.IdToCode(1)
	fmt.Printf("id=%v, code=%v\n", id, code)

	code = "VFXXXX"
	id = inviteCode.CodeToId(code)
	fmt.Printf("code=%v, id=%v\n", code, id)
}

func TestInviteCode2(t *testing.T) {
	InitBaseMap()

	fmt.Printf("len(baseStr):%d, len(base):%d\n", len(baseStr), len(base))

	res := Base34(1)

	fmt.Printf("=base:1544804416->%s, %d\n", string(res), len(res))

	str := "VIVZ4EH"

	num, err := Base34ToNum([]byte(str))

	if err == nil {
		fmt.Printf("=base:%s->%d\n", str, num)
	} else {
		fmt.Printf("===============err:%s\n", err.Error())
	}
}

func TestByteString(t *testing.T) {
	s := "a"
	b := []byte(s)

	fmt.Println(b)
}
