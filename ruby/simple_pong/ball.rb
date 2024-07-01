require "matrix"

class Ball

  def initialize
    @shape = Circle.new(
      x: WIDTH/2,
      y: HEIGHT/2,
      radius: 10,
      sectors: 32,
      color: 'fuchsia',
      z: 10,
    )
    @dir = Vector[rand(-1.0..1.0), rand(-1.0..1.0)].normalize
    @speed = 10
  end

  def update
    # FIXME: update dir

    # Update pos
    @shape.x += @dir[0] * @speed
    @shape.y += @dir[1] * @speed
  end

end
