package repositories

// Dependencies:
// go install -trimpath go.uber.org/mock/mockgen@latest

//go:generate mockgen -package mock -destination mock/mock.go . Client,WordOfWisdom
