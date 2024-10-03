extends HSlider


func _on_ready() -> void:
	self.grab_focus()
	var master_volume = ConfigFileHandler.load_audio_settings()["master_volume"]
	self.value = master_volume


func _on_mouse_entered() -> void:
	self.grab_focus()


func _process(_delta: float) -> void:
	%VolumeButton.text = "Volume: {vol}".format({"vol": self.value})
	var muted = ConfigFileHandler.load_audio_settings()["muted"]
	if muted:
		%VolumeButton.text = "Volume: ({vol})".format({"vol": self.value})


func _on_drag_ended(value_changed: bool) -> void:
	if value_changed:
		ConfigFileHandler.save_audio_settings("master_volume", self.value)
