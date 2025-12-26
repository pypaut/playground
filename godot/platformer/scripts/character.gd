extends CharacterBody2D

var SPEED = 400
var MAX_FALL_FORCE = 1500
var JUMP_FORCE = -700
var FALL_INCREASE = 25

var direction = Vector2.ZERO
var in_air = false


func _ready() -> void:
	direction.y = MAX_FALL_FORCE


func _physics_process(delta):
	# Up and down
	if not in_air and Input.is_action_just_pressed("space"):
		direction.y = JUMP_FORCE
	
	direction.y += FALL_INCREASE
	direction.y = clamp(direction.y, JUMP_FORCE, MAX_FALL_FORCE)
	
	# Left and right
	direction.x = Input.get_axis("left", "right") * SPEED
	
	# Apply collision and remove friction
	var collision = move_and_collide(direction * delta)
	if collision:
		var normal = collision.get_normal()
		if normal.y < 0:
			# Collision with floor, slow fall
			in_air = false
			direction.y = 0
		if normal.y > 0:
			# Collision with ceiling, slow fall
			direction.y = 0

		var remainder = collision.get_remainder().slide(normal)
		move_and_collide(remainder)
	else:
		in_air = true
