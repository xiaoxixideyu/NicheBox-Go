package biz

import "strings"

func CheckIfCommentOrderValid(order string) bool {
	order = strings.ToLower(order)
	if order != OrderByCreateTimeAsc && order != OrderByCreateTimeDesc && order != OrderByLikeCount {
		return false
	}
	return true
}

func CheckIfBoxContentOrderValid(order string) bool {
	order = strings.ToLower(order)
	if order != OrderByCreateTimeAsc && order != OrderByCreateTimeDesc {
		return false
	}
	return true
}
