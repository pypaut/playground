require 'socket'

socket = TCPServer.new('0.0.0.0', 8080)
client = socket.accept

puts "New client! #{client}"

client.write("Hello from server")
client.close
