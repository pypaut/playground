extends Button


func _on_mouse_entered() -> void:
	self.grab_focus()


func _on_pressed() -> void:
	var muted = ConfigFileHandler.load_audio_settings()["muted"]
	ConfigFileHandler.save_audio_settings("muted", !muted)


func _on_ready() -> void:
	var audio_settings = ConfigFileHandler.load_audio_settings()
	self.text = "Volume: {vol}".format({"vol": audio_settings["master_volume"]})
	if audio_settings["muted"]:
		self.text = "Volume: ({vol})".format({"vol": audio_settings["master_volume"]})


func _process(_delta: float) -> void:
	var audio_settings = ConfigFileHandler.load_audio_settings()
	self.text = "Volume: {vol}".format({"vol": audio_settings["master_volume"]})
	if audio_settings["muted"]:
		self.text = "Volume: ({vol})".format({"vol": audio_settings["master_volume"]})
