package nested

type EditorMemento struct {
	cursorPos int
	scrollPos int
	text      string
}

func NewEditorMemento(cursorPos int, scrollPos int, text string) *EditorMemento {
	return &EditorMemento{
		cursorPos: cursorPos,
		scrollPos: scrollPos,
		text:      text,
	}
}
