package usecases

func inRange(value string, start string, end string) bool {
	return value >= start && value <= end
}
