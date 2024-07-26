import configparser
import json
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
        if address == None:
            address = ""
        port = config.getint("SERVER", "PORT")

        # Networking
        print(f'Server listening to "{address}:{port}"')
        self.socket = socket.socket()
        self.socket.bind((address, port))
        self.socket.listen(5)

        # Game logic
        self.player_pos = {"x": 0, "y": 0}
        self.player_speed = 10

    def run(self):
        print("Server running")
        conn, _ = self.socket.accept()

        while True:
            client_msg = self.receive_direction(conn)
            if not client_msg:
                break

            player_dir = json.loads(client_msg)
            self.update_position(player_dir)

            self.send_position(conn)

        conn.close()

    def update_position(self, player_dir):
        # Update position
        self.player_pos["x"] += player_dir["x"] * self.player_speed
        self.player_pos["y"] += player_dir["y"] * self.player_speed

        # Clamp
        self.player_pos["x"] = max(0, self.player_pos["x"])
        self.player_pos["x"] = min(self.W - 50, self.player_pos["x"])
        self.player_pos["y"] = max(0, self.player_pos["y"])
        self.player_pos["y"] = min(self.H - 50, self.player_pos["y"])

    def receive_direction(self, conn):
        client_msg = conn.recv(1024)
        if client_msg == b"":
            return False
        return client_msg

    def send_position(self, conn):
        conn.send(f"{json.dumps(self.player_pos)}".encode("utf-8"))
