package common

import "nichebox/common/biz"

func CheckIfOrderValid(order string) bool {
	if order != biz.OrderByTimeDesc && order != biz.OrderByTimeAsc && order != biz.OrderByLikeCount {
		return false
	}
	return true
}
