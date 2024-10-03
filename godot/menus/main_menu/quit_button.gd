extends Button


func _on_pressed() -> void:
	get_tree().quit()


func _on_mouse_entered() -> void:
	self.grab_focus()
