extends Node

###############
### SIGNALS ###
###############

func _on_continue_button_pressed() -> void:
	$PauseMenu.hide()


func _on_leave_button_pressed() -> void:
	get_tree().change_scene_to_file("res://scenes/main_menu.tscn")

########################
### STANDARD METHODS ###
########################

func _process(_delta: float) -> void:
	if Input.is_action_just_pressed("escape"):
		$PauseMenu.show()
