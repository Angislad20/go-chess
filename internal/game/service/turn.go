package service

func TogglePlayer(current string) string {
	if current == "w" {
		return "b"
	}
	return "w"
}
