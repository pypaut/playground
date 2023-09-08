import socket


# Create & connect socket
s = socket.socket()
port = 12350
address = '127.0.0.1'
address = '192.168.149.119'
s.connect((address, port))

# Send data to server (newline terminated)
msg = b'Hi\n'
s.send(msg)

# Receive data from server
# msg = s.recv(1024))

# Close socket connection

s.close()
