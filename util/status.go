package util

const (
	unknownStatus = "unknown"
)

var (
	statusMap = map[string]string{
		"0": "offline",
		"1": "effected",
		"2": "pending",
		"3": "success",
		"4": "failed",
	}
)

// GetStatus filter thr status from client input
func GetStatus(s string) string {
	if val, ok := statusMap[s]; ok {
		return val
	}
	return unknownStatus
}
