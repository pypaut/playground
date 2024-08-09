package nested

type Editor struct {
	cursorPos int
	scrollPos int
	text      string
}

func (e *Editor) BuildMemento() *EditorMemento {
	return NewEditorMemento(e.cursorPos, e.scrollPos, e.text)
}

func (e *Editor) Restore(memento *EditorMemento) {
	e.cursorPos = memento.cursorPos
	e.scrollPos = memento.scrollPos
	e.text = memento.text
}
