require 'socket'

# socket = Socket.new(Socket::AF_INET, Socket::SOCK_STREAM)
# socket.connect('0.0.0.0', 8080)

socket = TCPSocket.new '0.0.0.0', 8080

socket.sendmsg("hello server!" + "\n")
puts socket.recv 20

socket.close
