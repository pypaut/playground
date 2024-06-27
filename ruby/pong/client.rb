require 'socket'


def main

  # Connect to server
  socket = TCPSocket.new '0.0.0.0', 8080
  puts "Connected to server"

  # Read init message
  server_msg = socket.recv 20
  puts "Server: \"#{server_msg}\""

  loop do
    # Send
    socket.sendmsg("some message")

    # Receive
    server_msg = socket.recv 20
  end

rescue Errno::EPIPE, Errno::ECONNRESET, Interrupt

  puts "Disconnected"

ensure

  socket.close

end

main
