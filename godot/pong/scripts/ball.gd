extends CharacterBody2D

var HAS_STARTED = false
@export var speed = 400

func _ready():
	var x = (randi() % 100 + 50) * (-1)**(randi()%2)
	var y = (randi() % 100) * (-1)**(randi()%2)

	velocity = Vector2(x, y).normalized() * speed

func _process(delta: float):
	if not HAS_STARTED and Input.is_action_just_pressed("space", true):
		HAS_STARTED = true

	if not HAS_STARTED:
		return

	var collision: KinematicCollision2D = move_and_collide(velocity * delta)
	if collision:
		var reflect = collision.get_remainder().bounce(collision.get_normal())
		velocity = velocity.bounce(collision.get_normal())
		move_and_collide(reflect)
