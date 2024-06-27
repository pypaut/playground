require 'socket'


# Functions

def handle_connection(player_socket, players, player_i)
  players[player_i] = 1
  client_nb = player_i + 1
  print "Player #{client_nb} connected\n"
  player_socket.sendmsg("You are player #{client_nb}")

  while msg = player_socket.recv(50)
    if msg == ""
      break
    end
    player_socket.sendmsg("Ok")
  end
rescue Errno::ECONNRESET
  # Catch client disconnect error
ensure
  player_socket.close
  print "Player #{client_nb} disconnected\n"
  players[player_i] = 0
end


# Main

players = Array.new(2, 0)
server_socket = TCPServer.new 8080
puts "Server launched"

loop do
  player_socket = server_socket.accept
  if players.include? 0
    player_i = players.find_index(0)
    Thread.new { handle_connection(player_socket, players, player_i) }
  else
    player_socket.sendmsg("Sorry, lobby is full.")
    player_socket.close
    # TODO : login queue system
    # TODO : spec mode
  end
end
