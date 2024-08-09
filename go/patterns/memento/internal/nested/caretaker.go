package nested

type Caretaker struct {
	editor  *Editor
	history []*EditorMemento
}

func (c *Caretaker) DoSomething() {
	// Business code here...

	memento := editor.BuildMemento()
	c.history = append(c.history, memento)
}

func (c *Caretaker) Undo() {
	c.editor.Restore(c.history[len(c.history)-1])
	c.history = c.history[:len(c.history)-1]
}
