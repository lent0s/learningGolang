package service

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// проверка и приведение имени к нормальному виду
func nameCheck(n string) (string, error) {

	nameFull := strings.Fields(n)
	for partNum := 0; partNum < len(nameFull); partNum++ {
		if err := checkLegalChar(nameFull[partNum]); err != nil {
			return "", err
		}
		nameFull[partNum] = removeIncorrectDash(nameFull[partNum])
		if len(nameFull[partNum]) == 0 {
			nameFull = append(nameFull[:partNum], nameFull[partNum+1:]...)
			partNum--
			continue
		}
		nameFull[partNum] = strings.ToLower(nameFull[partNum])
		if checkNotPatronymic(nameFull[partNum]) {
			nameFull[partNum] = dropCaps(nameFull[partNum])
		}
	}
	n = strings.Join(nameFull, " ")
	if len(n) == 0 {
		return "", fmt.Errorf("user name is empty")
	}
	return n, nil
}

// проверка на допустимые кирилл. и латиницу
func checkLegalChar(s string) error {

	var legalChar = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0020, 0x0020, 1}, // space
			{0x002D, 0x002D, 1}, // -
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005A, 1}, // A-Z
			{0x0061, 0x007A, 1}, // a-z
			//{0x0401, 0x0401, 1}, // Ё
			//{0x0410, 0x044F, 1}, // А-Я а-я
			//{0x0451, 0x0451, 1}, // ё
		},
		R32: []unicode.Range32{},
	}

	allowed := func() string {
		allowed := ""
		for _, set := range legalChar.R16 {
			allowed += "["
			interval := strings.Split(strings.Trim(
				fmt.Sprint(set), "{}"), " ")
			lo, _ := strconv.Atoi(interval[0])
			hi, _ := strconv.Atoi(interval[1])
			for ; lo <= hi; lo++ {
				allowed += string(lo)
			}
			allowed += "] "
		}
		return allowed
	}

	for _, letter := range s {
		if !unicode.In(letter, legalChar) {
			return fmt.Errorf("user name contains illegal characters:"+
				" [%s] %U\n%v", string(letter), letter, allowed())
		}
	}
	return nil
}

// убрать тире в начале, в конце, дубли
func removeIncorrectDash(s string) string {

	for i := 0; i < len(s)-1; i++ {
		if s[i] == 45 && s[i+1] == 45 {
			s = s[:i] + s[i+1:]
			i--
		}
	}
	if strings.LastIndex(s, "-") == len(s)-1 {
		s = s[:len(s)-1]
	}
	if strings.Index(s, "-") == 0 {
		s = s[1:]
	}
	return s
}

// проверка на патронимы
func checkNotPatronymic(s string) bool {

	var patronymic = []string{
		//male:
		"ибн", "сын", "бен", "бар", "оглы", "оглу", "уулу", "улы",
		"ibn", "son", "ben", "bin", "ogly", "fils de", "ina", "ap", "ab",
		//female:
		"бинт", "дочь", "бат", "кызы", "гызы",
		"bint", "daughter", "fille de", "ferch", "verch", "merch", "qizi",
	}

	for _, prefix := range patronymic {
		if s == prefix {
			return false
		}
	}
	return true
}

// регистр "как в предложениях"
func dropCaps(s string) string {

	parts := strings.Split(s, "-")
	for _, part := range parts {
		if part == "" {
			return ""
		}
	}
	for i, _ := range parts {
		parts[i] = strings.ToTitle(string([]rune(parts[i])[0])) +
			string([]rune(parts[i])[1:])
	}
	s = strings.Join(parts, "-")
	return s
}

// проверка на отсутствие дружбы
func alreadyFriends(sourceIdFriends []int, targetId int) bool {

	for _, detected := range sourceIdFriends {
		if detected == targetId {
			return true
		}
	}
	return false
}

// удаление юзера из списка друзей
func deleteFriend(sourceId *[]int, targetId int) {

	for i, friend := range *sourceId {
		if friend == targetId {
			*sourceId = append((*sourceId)[:i], (*sourceId)[i+1:]...)
			return
		}
	}
}
