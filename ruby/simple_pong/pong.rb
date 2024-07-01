require 'ruby2d'

require './constants.rb'
require './player.rb'


set title: "Pong"
set width: WIDTH, height: HEIGHT

p1 = Player.new(100, 'red', "w", "s")
p2 = Player.new(850, 'blue', "i", "k")

on :key_held do |event|
  if event.key == "escape"
    close
  end

  p1.key_held(event.key)
  p2.key_held(event.key)

end

show
