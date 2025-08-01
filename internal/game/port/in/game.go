package in

type GameUseCase interface {
	StartGame()
	MakeMove(move string) error
}
