import socket


# Create socket & bind to port
s = socket.socket()
port = 12350
s.bind(('', port))
s.listen(5)

while True:
    # Connection to client
    c, addr = s.accept()

    # Receive data from client
    msg = c.recv(1024)
    print(f"Received {msg} from Client")

    # Send data to client
    # c.send(b'Thank you for connecting')

    # Close connection with client
    c.close()
