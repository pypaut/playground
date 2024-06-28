require 'ruby2d'


set title: "Pong"
set width: 1000, height: 800
player_speed = 10

p1 = Rectangle.new(
  x: 100, y: 100, width: 25, height: 150, color: 'red'
)
p2 = Rectangle.new(
  x: 850, y: 100, width: 25, height: 150, color: 'blue'
)

on :key_held do |event|
  case event.key
  when "w"
    p1.y -= player_speed
  when "s"
    p1.y += player_speed
  end
end

show
