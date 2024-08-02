require 'ruby2d'
require 'socket'



def main
  # Connect to server
  socket = TCPSocket.new '0.0.0.0', 8080
  puts "Connected to server"

  # Read init message
  server_msg = socket.recv 20
  puts "Server: \"#{server_msg}\""

  # Init game
  set width: 1000, height: 800
  s = Square.new
  s.color = 'red'
  show

  loop do
    # Server communication
    socket.sendmsg("some message")
    server_msg = socket.recv 20

    # on :key_down do |event|
    #   puts event.key
    # end
    # Handle input
    dir_x = 0
    dir_y = 0
    on :key_held do |event|
      case event.key
      when "w"
        dir_y -= 1
      when "a"
        dir_x -= 1
      when "s"
        dir_y += 1
      when "d"
        dir_x += 1
      end
    end

    s.x += dir_x
    s.y += dir_y

  end
rescue Errno::EPIPE, Errno::ECONNRESET, Interrupt
  puts "Disconnected"
ensure
  socket.close
end

main
