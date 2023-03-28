package service

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// check data in request
func convReqToService(req map[string]interface{}) (s Service, err error) {

	s.Name = fmt.Sprint(req["name"])

	if s.Age, err = strconv.Atoi(fmt.Sprint(req["age"])); (err != nil ||
		s.Age < 0) && req["age"] != nil {
		return s, fmt.Errorf("user age is incorrect")
	}

	if s.SourceId, err = strconv.Atoi(fmt.Sprint(req["source_id"])); err != nil &&
		req["source_id"] != nil {
		return s, fmt.Errorf("source_id is incorrect")
	}

	if s.TargetId, err = strconv.Atoi(fmt.Sprint(req["target_id"])); err != nil &&
		req["target_id"] != nil {
		return s, fmt.Errorf("target_id is incorrect")
	}

	if s.NewAge, err = strconv.Atoi(fmt.Sprint(req["new age"])); (err != nil ||
		s.NewAge < 0) && req["new age"] != nil {
		return s, fmt.Errorf("user new age is incorrect")
	}
	return s, nil
}

// name spelling check, normalize view of name
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

// check for legal chars: latin, cyrillic, numbers, space, dash
func checkLegalChar(s string) error {

	var legalChar = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0020, 0x0020, 1}, // space
			{0x002D, 0x002D, 1}, // - dash
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005A, 1}, // A-Z
			{0x0061, 0x007A, 1}, // a-z
			{0x0401, 0x0401, 1}, // Ё
			{0x0410, 0x044F, 1}, // А-Я а-я
			{0x0451, 0x0451, 1}, // ё
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

// trim excess dashes: doubles, prefix, postfix
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

// true if s is not patronymic
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

// case like sentences
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

// check for friendship
func alreadyFriends(sourceIdFriends []int, targetId int) bool {

	for _, friend := range sourceIdFriends {
		if friend == targetId {
			return true
		}
	}
	return false
}

// delete user from friends list
func deleteFriend(sourceId *[]int, targetId int) {

	for i, friend := range *sourceId {
		if friend == targetId {
			*sourceId = append((*sourceId)[:i], (*sourceId)[i+1:]...)
			return
		}
	}
}

// append friend in sorted list
func appendFriend(friends []int, friend int) (result []int) {

	if len(friends) == 0 {
		return append(result, friend)
	}
	min := 0
	max := len(friends) - 1
	current := 0

	for {
		if friend < friends[min] {
			result = append(result, friends[:min]...)
			result = append(result, friend)
			return append(result, friends[min:]...)
		}
		if friend > friends[max] {
			result = append(result, friends[:max]...)
			result = append(result, friends[max])
			result = append(result, friend)
			if max == len(friends)-1 {
				return
			}
			return append(result, friends[max+1:]...)
		}

		current = min + (max-min)/2
		if friends[current] > friend {
			max = current - 1
		} else {
			min = current + 1
		}
	}
}
