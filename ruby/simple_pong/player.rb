require './constants.rb'

class Player

  def initialize(x, color, up_key, down_key)
    @speed = 10
    @rect = Rectangle.new(x: x, y: 100, width: 25, height: 150, color: color)
    @up_key = up_key
    @down_key = down_key
  end

  def key_held(key)
    case key
    when @up_key
      @rect.y -= @speed
    when @down_key
      @rect.y += @speed
    end

    @rect.y = @rect.y.clamp(0, HEIGHT - 150)
  end

end
