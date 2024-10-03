extends Button


func _on_mouse_entered() -> void:
	self.grab_focus()


func _on_pressed() -> void:
	get_tree().change_scene_to_file("res://options_menu/options_menu.tscn")
