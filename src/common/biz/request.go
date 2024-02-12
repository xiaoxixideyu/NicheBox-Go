package biz

func CheckIfCommentOrderValid(order string) bool {
	if order != OrderByCreateTimeAsc && order != OrderByCreateTimeDesc && order != OrderByLikeCount {
		return false
	}
	return true
}

func CheckIfBoxContentOrderValid(order string) bool {
	if order != OrderByCreateTimeAsc && order != OrderByCreateTimeDesc {
		return false
	}
	return true
}
