extends Node


var config = ConfigFile.new()
const SETTINGS_FILE_PATH = "user://settings.ini"

func _ready():
	if !FileAccess.file_exists(SETTINGS_FILE_PATH):
		config.set_value("audio", "master_volume", 50)
		config.set_value("audio", "muted", false)
		config.set_value("video", "fullscreen", true)
		config.save(SETTINGS_FILE_PATH)
	else:
		config.load(SETTINGS_FILE_PATH)


func save_video_settings(key: String, value):
	config.set_value("video", key, value)
	config.save(SETTINGS_FILE_PATH)


func load_video_settings() -> Dictionary:
	var video_settings = {}
	for key in config.get_section_keys("video"):
		video_settings[key] = config.get_value("video", key)
	return video_settings
	
	
func save_audio_settings(key: String, value):
	config.set_value("audio", key, value)
	config.save(SETTINGS_FILE_PATH)
	
	
func load_audio_settings() -> Dictionary:
	var audio_settings = {}
	for key in config.get_section_keys("audio"):
		audio_settings[key] = config.get_value("audio", key)
	return audio_settings