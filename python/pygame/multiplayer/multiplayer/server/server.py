import configparser
import json
import socket


class Server:
    def __init__(self):
        self.socket = socket.socket()
        self.socket.bind(("", 12350))
        self.socket.listen(5)


        # Read config
        config = configparser.ConfigParser()
        with open("config.ini", "r") as configfile:
            config.read_file(configfile)

        self.H = config.getint("CLIENT", "H")
        self.W = config.getint("CLIENT", "W")

        # Game logic
        self.player_pos = {"x": 0, "y": 0}
        self.player_speed = 10

    def run(self):
        print("Server running")
        conn, _ = self.socket.accept()

        while True:
            # Receive direction
            client_msg = conn.recv(1024)
            if client_msg == b"":
                break

            player_dir = json.loads(client_msg)

            # Update position
            self.player_pos["x"] += player_dir["x"] * self.player_speed
            self.player_pos["y"] += player_dir["y"] * self.player_speed

            # Clamp
            self.player_pos["x"] = max(0, self.player_pos["x"])
            self.player_pos["x"] = min(self.W - 50, self.player_pos["x"])
            self.player_pos["y"] = max(0, self.player_pos["y"])
            self.player_pos["y"] = min(self.H - 50, self.player_pos["y"])

            # Send position
            conn.send(f"{json.dumps(self.player_pos)}".encode("utf-8"))

        conn.close()
