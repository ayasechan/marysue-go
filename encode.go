package marysue

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	CHARACTERS  = "薰璃安莹洁莉樱殇雪羽晗灵血娜丽魑魅塔利亚伤梦儿海蔷玫瑰泪邪凡多姆威恩夏影琉舞雅蕾玥瑷曦月瑟薇蓝岚紫蝶馨琦洛凤颜鸢希玖兮雨烟叶兰凝冰伊如落心语凌爱陌悠千艳优花晶墨阳云筱残莲沫渺琴依然丝可茉黎幽幻银韵倾乐慕文思蕊清碎音芊黛怡莎苏香城萌美迷离白嫩风霜萝妖百合珠喃之倩情恋弥绯芸茜魂澪琪欣呗缈娅吉拉斯基柔惠朵茹妙铃裳纱颖蕴燢浅萦璎糜凪莳娥寂翼巧哀俏涅盘辰芝艾柒曼妲眉御寇妮米菲奥格萨温蒂"
	SPLITER     = "·"
	BUFFER_SIZE = 64 / 8
)

var (
	CHARACTERS_LENGTH = uint64(utf8.RuneCountInString(CHARACTERS))
)

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min) + min)
}

func findString(s, sub string) int {
	byteIndex := strings.Index(s, sub)
	if byteIndex == -1 {
		return byteIndex
	}
	return utf8.RuneCountInString(s[:byteIndex])
}

func UInt64ToString(n uint64) string {
	r := []rune{}
	runeCharacters := []rune(CHARACTERS)

	for n > 0 {
		q := n % CHARACTERS_LENGTH
		r = append(r, runeCharacters[q])
		n = n / CHARACTERS_LENGTH
	}
	return string(r)
}

func ByteArrayToString(cipherText []byte) string {
	r := []string{}
	cursor := 0
	for cursor < len(cipherText) {
		l := randInt(1, BUFFER_SIZE)

		if cursor+l > len(cipherText) {
			l = len(cipherText) - cursor
		}
		if cipherText[cursor+l-1] == 0 {
			continue
		}
		buf := make([]byte, BUFFER_SIZE)
		copy(buf, cipherText[cursor:cursor+l])
		n := binary.LittleEndian.Uint64(buf)
		r = append(r, UInt64ToString(n))
		cursor += l
	}
	return strings.Join(r, SPLITER)
}

func StringToUInt64(s string) uint64 {
	var n uint64
	runeS := []rune(s)

	i := len(runeS) - 1
	for i >= 0 {
		q := findString(CHARACTERS, string(runeS[i]))
		if q == -1 {
			panic(fmt.Sprintf("illegal characters: %v", runeS[i]))
		}
		n = n*CHARACTERS_LENGTH + uint64(q)
		i -= 1
	}
	return n
}

func StringToByteArray(s string) []byte {
	r := new(bytes.Buffer)
	for _, frag := range strings.Split(s, SPLITER) {
		n := StringToUInt64(frag)
		buf := make([]byte, BUFFER_SIZE)
		binary.LittleEndian.PutUint64(buf, n)
		trimIndex := BUFFER_SIZE - 1
		for buf[trimIndex] == 0 {
			trimIndex -= 1
		}
		r.Write(buf[:trimIndex+1])
	}
	return r.Bytes()
}
