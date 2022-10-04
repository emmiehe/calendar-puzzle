package puzzle

type Board struct {
	Positions [][]Slot
}

type Slot struct {
	Label     string
	Available bool // some slots might be unavailable even before game starts

	PieceID int // 0 indicates no piece
}

type Piece struct {
	ID    int
	Shape []Position
	Block [][]int
}

type Position struct {
	Row int
	Col int
}

// Normalize normalize a given piece's shape?
func (p *Piece) Normalize() {
}

// Rotate rotates the piece to the right
func (p *Piece) Rotate() {
}

// Flip flips the piece 180 along y-axis
func (p *Piece) Flip() {
}

// Put puts a piece on the board, if successful, returns true
// pos always means the top left corner of the piece's box position
func (board *Board) Put(piece *Piece, pos *Position) bool {
	var toBeSet []*Slot
	for _, shapePosition := range piece.Shape {
		r := shapePosition.Row + pos.Row
		c := shapePosition.Col + pos.Col
		if r >= len(board.Positions) || c >= len(board.Positions[0]) {
			return false
		}

		slot := &board.Positions[r][c]
		if !slot.Available || slot.PieceID != 0 {
			return false
		}
		toBeSet = append(toBeSet, slot)
	}
	for _, slot := range toBeSet {
		slot.PieceID = piece.ID
	}
	return true
}
