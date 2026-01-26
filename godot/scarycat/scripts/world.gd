extends Node3D

var KIRA_VISIBLE = false
var KIRA_SPEED = 1000


func _ready() -> void:
	# $MenuCamera.make_current()
	pass


func _process(delta: float) -> void:
	var character_pos = $Character.global_position
	
	# Sprites look at player
	$Nala.look_at(character_pos)
	$Kira.look_at(character_pos)
	$Jack.look_at(character_pos)
	
	# Move Kira
	if KIRA_VISIBLE:
		var kira_direction = $Kira.global_position.direction_to(character_pos)
		$Kira.velocity = kira_direction * KIRA_SPEED * delta
		$Kira.move_and_slide()


func _on_visible_on_screen_notifier_3d_screen_entered() -> void:
	KIRA_VISIBLE = true


func _on_visible_on_screen_notifier_3d_screen_exited() -> void:
	KIRA_VISIBLE = false
