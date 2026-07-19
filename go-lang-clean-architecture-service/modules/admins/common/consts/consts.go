package consts

const (
	STATUS_ACTIVE   = 1
	STATUS_INACTIVE = 2
)

var (
	StatusUserToViewCreate = []int{
		STATUS_ACTIVE,
		STATUS_INACTIVE}
	MapStatus = map[int]string{
		STATUS_ACTIVE:   "Hien",
		STATUS_INACTIVE: "An",
	}
)
