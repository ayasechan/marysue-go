package marysue

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const CHARACTERS = "薰璃安莹洁莉樱殇雪羽晗灵血娜丽魑魅塔利亚伤梦儿海蔷玫瑰泪邪凡多姆威恩夏影琉舞雅蕾玥瑷曦月瑟薇蓝岚紫蝶馨琦洛凤颜鸢希玖兮雨烟叶兰凝冰伊如落心语凌爱陌悠千艳优花晶墨阳云筱残莲沫渺琴依然丝可茉黎幽幻银韵倾乐慕文思蕊清碎音芊黛怡莎苏香城萌美迷离白嫩风霜萝妖百合珠喃之倩情恋弥绯芸茜魂澪琪欣呗缈娅吉拉斯基柔惠朵茹妙铃裳纱颖蕴燢浅萦璎糜凪莳娥寂翼巧哀俏涅盘辰芝艾柒曼妲眉御寇妮米菲奥格萨温蒂"
const Spliter = "·"

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min) + min)
}

func findString(s, sub string) int {
	byteIndex := strings.Index(s, sub)
	if byteIndex == -1 {
		return byteIndex
	}
	return len([]rune(s[:byteIndex]))
}

func UInt64ToString(n uint64) string {
	r := []rune{}
	runeCharacters := []rune(CHARACTERS)
	length := uint64(len(runeCharacters))

	for n > 0 {
		q := n % length
		r = append(r, runeCharacters[q])
		n = n / length
	}
	return string(r)
}

func ByteArrayToString(cipherText []byte) string {
	r := []string{}
	cursor := 0
	source := cipherText
	length := 8
	for cursor < len(cipherText) {
		l := randInt(1, length)

		if cursor+l > len(cipherText) {
			l = len(cipherText) - cursor
		}
		if source[cursor+l-1] == 0 {
			continue
		}
		buf := make([]byte, length)
		copy(buf, source[cursor:cursor+l])
		n := binary.LittleEndian.Uint64(buf)
		r = append(r, UInt64ToString(n))
		cursor += l
	}
	return strings.Join(r, Spliter)
}

func StringToUInt64(s string) uint64 {
	var n uint64
	runeS := []rune(s)
	length := len([]rune(CHARACTERS))

	i := len(runeS) - 1
	for i >= 0 {
		q := findString(CHARACTERS, string(runeS[i]))
		if q == -1 {
			panic(fmt.Sprintf("illegal characters: %v", runeS[i]))
		}
		n = n*uint64(length) + uint64(q)
		i -= 1
	}
	return n
}

func StringToByteArray(s string) []byte {
	r := new(bytes.Buffer)
	for _, frag := range strings.Split(s, Spliter) {
		n := StringToUInt64(frag)
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, n)
		trimIndex := 8 - 1
		for buf[trimIndex] == 0 {
			trimIndex -= 1
		}
		r.Write(buf[:trimIndex+1])
	}
	return r.Bytes()
}
