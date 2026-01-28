extends Node3D

var KIRA_IN_FOV = false
var KIRA_SPEED = 1000


############
# BUILTINS #
############

func _process(delta: float) -> void:
	var character_pos = $Character.global_position
	character_pos.y += 2.5 # so that they look at my head

	# Sprites look at player
	$Nala.look_at(character_pos)
	$Jack.look_at(character_pos)

	handle_kira(delta, character_pos)


##########
# CUSTOM #
##########

func handle_kira(delta: float, character_pos: Vector3):
	# Conditions: close enough + kira in fov
	var distance_with_player = $Kira.global_position.distance_to(character_pos)
	if distance_with_player > 50 or not KIRA_IN_FOV:
		return

	$Kira.look_at(character_pos)
	var kira_direction = $Kira.global_position.direction_to(character_pos)
	$Kira.velocity = kira_direction * KIRA_SPEED * delta
	$Kira.move_and_slide()
	

###########
# SIGNALS #
###########

func _on_visible_on_screen_notifier_3d_screen_entered() -> void:
	KIRA_IN_FOV = true

func _on_visible_on_screen_notifier_3d_screen_exited() -> void:
	KIRA_IN_FOV = false
