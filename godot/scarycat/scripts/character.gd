extends CharacterBody3D


const SPEED = 8.0
const JUMP_VELOCITY = 10.0
const CAMERA_SENSITIVITY = 10
const MAX_CAMERA_ANGLE = 1.5


func _ready() -> void:
	Input.set_mouse_mode(Input.MOUSE_MODE_CAPTURED)


func _unhandled_input(event: InputEvent) -> void:
	if event is InputEventMouseMotion:
		# Rotate the whole character left and right
		rotation.y = rotation.y - event.relative.x * 1/(CAMERA_SENSITIVITY*10)
		# Rotate the camera only up and down
		$Camera3D.rotation.x = $Camera3D.rotation.x - event.relative.y * 1/(CAMERA_SENSITIVITY*10)
		print($Camera3D.rotation.x)
		$Camera3D.rotation.x = clamp($Camera3D.rotation.x, -MAX_CAMERA_ANGLE, MAX_CAMERA_ANGLE)


func _physics_process(delta: float) -> void:
	# Add the gravity.
	if not is_on_floor():
		velocity += get_gravity() * delta * 3.0

	# Handle jump.
	if Input.is_action_just_pressed("space") and is_on_floor():
		velocity.y = JUMP_VELOCITY

	# Get the input direction and han dle the movement/deceleration.
	# As good practice, you should replace UI actions with custom gameplay actions.
	var input_dir := Input.get_vector("left", "right", "up", "down")
	var direction := (transform.basis * Vector3(input_dir.x, 0, input_dir.y)).normalized()
	if direction:
		velocity.x = direction.x * SPEED
		velocity.z = direction.z * SPEED
	else:
		velocity.x = move_toward(velocity.x, 0, SPEED)
		velocity.z = move_toward(velocity.z, 0, SPEED)

	move_and_slide()
