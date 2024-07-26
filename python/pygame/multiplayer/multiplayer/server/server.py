import json
import socket

class Server:
    def __init__(self):
        self.socket = socket.socket()
        self.socket.bind(('', 12350))
        self.socket.listen(5)

    def run(self):
        print("Server running")
        conn, _ = self.socket.accept()

        while True:
            client_msg = conn.recv(1024)
            if client_msg == b'':
                break

            client_dict = json.loads(client_msg)
            print(client_dict)

        conn.close()
