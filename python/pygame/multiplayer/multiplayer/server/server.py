import configparser
import json
import select
import socket


class Server:
    def __init__(self):
        # Read config
        config = configparser.ConfigParser()
        with open("server.ini", "r") as configfile:
            config.read_file(configfile)

        self.H = config.getint("SERVER", "H")
        self.W = config.getint("SERVER", "W")

        address = config["SERVER"]["ADDRESS"]
        if address is None:
            address = ""
        port = config.getint("SERVER", "PORT")

        # Networking
        print(f'Server listening to "{address}:{port}"')
        self.socket = socket.socket()
        self.socket.bind((address, port))
        self.socket.listen(5)
        self.socks = [self.socket]

        # Game logic
        self.player_speed = 10
        self.players = {self.socket.fileno(): {"x": 0, "y": 0}}

    def run(self):
        print("Server running")

        while True:
            socks_to_read, _, _ = select.select(self.socks, [], [])

            for s in socks_to_read:
                if s == self.socket:
                    self.handle_new_client()
                else:
                    self.handle_client_message(s)

    def update_position(self, index, player_dir):
        # Update position
        self.players[index]["x"] += player_dir["x"] * self.player_speed
        self.players[index]["y"] += player_dir["y"] * self.player_speed

        # Clamp
        self.players[index]["x"] = max(0, self.players[index]["x"])
        self.players[index]["x"] = min(self.W - 50, self.players[index]["x"])
        self.players[index]["y"] = max(0, self.players[index]["y"])
        self.players[index]["y"] = min(self.H - 50, self.players[index]["y"])

    def receive_direction(self, conn):
        client_msg = conn.recv(1024)
        if client_msg == b"":
            return False
        return client_msg

    def send_position(self, conn):
        conn.send(f"{json.dumps(self.players[conn.fileno()])}".encode("utf-8"))

    def handle_new_client(self):
        conn, addr = self.socket.accept()
        self.log(conn, "connected")

        # Load direction from client
        message = conn.recv(1024)
        player_dir = json.loads(message)

        # Add new player, update pos, send pos
        self.players[conn.fileno()] = {"x": 0, "y": 0}
        self.update_position(conn.fileno(), player_dir)
        self.send_position(conn)

        self.socks.append(conn)

    def handle_client_message(self, s):
        message = s.recv(1024)

        if not message:
            self.disconnect_client(s)
            return

        # Load direction from client
        player_dir = json.loads(message)
        self.update_position(s.fileno(), player_dir)
        self.send_position(s)

    def log(self, sock, message):
        print(f"{sock.fileno()} : {message}")

    def disconnect_client(self, sock):
        del self.players[sock.fileno()]
        self.log(sock, "disconnected")
        sock.close()
        self.socks.remove(sock)
