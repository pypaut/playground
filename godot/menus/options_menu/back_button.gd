extends Button


func _on_pressed() -> void:
	get_tree().change_scene_to_file("res://main_menu/main_menu.tscn")


func _on_mouse_entered() -> void:
	self.grab_focus()
