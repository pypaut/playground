extends Sprite2D


func _on_left_wall_body_entered(_body: Node2D) -> void:
	get_tree().quit()


func _on_right_wall_body_entered(_body: Node2D) -> void:
	get_tree().quit()
