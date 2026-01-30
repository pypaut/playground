extends Area3D

@onready var dialogue_label: RichTextLabel = $DialogueSprite/SubViewport/RichTextLabel

func _ready() -> void:
	dialogue_label.text = ""

func start_interact():
	"""
	Is character.gd, we check for the presence of this method before calling it
	"""
	dialogue_label.text = "Hello there."

func stop_interact():
	"""
	Is character.gd, we check for the presence of this method before calling it
	"""
	dialogue_label.text = ""
