extends CharacterBody2D

var PLAYER_SPEED = 700

@onready var ANIMATION_PLAYER = $AnimationPlayer
@onready var SPRITE = $Sprite2D

func _ready() -> void:
	ANIMATION_PLAYER.play("idle")

func _process(delta: float) -> void:
	var direction = 0
	if Input.is_action_pressed("left"):
		direction -= 1
	if Input.is_action_pressed("right"):
		direction += 1
		
	if abs(direction) > 0:
		ANIMATION_PLAYER.play("run")
	else:
		ANIMATION_PLAYER.play("idle")
		
	if direction < 0:
		SPRITE.flip_h = true
	elif direction > 0:
		SPRITE.flip_h = false
		
	position.x += direction * PLAYER_SPEED * delta
