extends OptionButton


func _on_mouse_entered() -> void:
	self.grab_focus()


func _on_item_selected(index: int) -> void:
	if index == 0: # Fullscreen
		ConfigFileHandler.save_video_settings("fullscreen", true)
	elif index == 1: # Windowed
		ConfigFileHandler.save_video_settings("fullscreen", false)
	else:
		pass
