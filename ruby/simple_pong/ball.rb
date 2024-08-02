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

  def update(p1, p2)
    # Wall collisions
    if @shape.x - @shape.radius < 0 or WIDTH < @shape.x + @shape.radius
      @dir[0] *= -1
    end

    if @shape.y - @shape.radius < 0 or HEIGHT < @shape.y + @shape.radius
      @dir[1] *= -1
    end

    # Player collision
    if @shape.intersects?(p1.rect) or @shape.intersects?(p2.rect)
      @dir[0] *= -1
    end

    # Update pos
    @shape.x += @dir[0] * @speed
    @shape.y += @dir[1] * @speed
  end

end
