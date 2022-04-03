package helper

var Array = array{}

type array struct {
}

// 检查元素是否在集合中
func (h *array) In(value string, arr []string) bool {

	for _, item := range arr {
		if item == value {
			return true
		}
	}

	return false
}
