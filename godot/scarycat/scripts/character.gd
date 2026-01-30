extends CharacterBody3D

const SPEED = 8.0
const JUMP_VELOCITY = 10.0
const CAMERA_SENSITIVITY = 10
const MAX_CAMERA_ANGLE = 1.5

@onready var raycast: RayCast3D = $Camera3D/RayCast3D

var interactable = null

################
### BUILTINS ###
################


func _ready() -> void:
	Input.set_mouse_mode(Input.MOUSE_MODE_CAPTURED)


func _process(_delta: float):
	check_interaction()


func _input(event):
	"""
	Handle "interact" input to start interaction
	"""
	if interactable and event.is_action_pressed("interact"):
		interactable.start_interact()


func _unhandled_input(event: InputEvent) -> void:
	"""
	Mouse movements
	"""
	if event is InputEventMouseMotion:
		# Rotate the whole character left and right
		rotation.y = rotation.y - event.relative.x * 1/(CAMERA_SENSITIVITY*10)

		# Rotate the camera only up and down
		$Camera3D.rotation.x = $Camera3D.rotation.x - event.relative.y * 1/(CAMERA_SENSITIVITY*10)
		$Camera3D.rotation.x = clamp($Camera3D.rotation.x, -MAX_CAMERA_ANGLE, MAX_CAMERA_ANGLE)


func _physics_process(delta: float) -> void:
	"""
	Gravity, keyboard movements
	"""
	# Add the gravity
	if not is_on_floor():
		velocity += get_gravity() * delta * 3.0

	# Handle jump
	if Input.is_action_just_pressed("space") and is_on_floor():
		velocity.y = JUMP_VELOCITY

	var input_dir := Input.get_vector("left", "right", "up", "down")
	var direction := (transform.basis * Vector3(input_dir.x, 0, input_dir.y)).normalized()
	if direction:
		velocity.x = direction.x * SPEED
		velocity.z = direction.z * SPEED
	else:
		velocity.x = move_toward(velocity.x, 0, SPEED)
		velocity.z = move_toward(velocity.z, 0, SPEED)

	move_and_slide()


##############
### CUSTOM ###
##############


func check_interaction():
	"""
	Use the raycast to check for interaction
	"""
	var is_colliding = raycast.is_colliding()

	# No collisions at all, reset interaction
	if not is_colliding:
		reset_interaction()
		return

	if is_colliding:
		interactable = raycast.get_collider()

		# is null or is random object
		if not interactable or not interactable.has_method("start_interact"):
			reset_interaction()
			return

		$HUD/InteractionText.text = "E - " + interactable.get_name()
		return


func reset_interaction():
	"""
	"""
	$HUD/InteractionText.text = ""
	if interactable != null and interactable.has_method("stop_interact"):
		interactable.stop_interact()
	interactable = null
