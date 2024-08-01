import sys
import select

from socket import socket, AF_INET, SOCK_STREAM


def log(sock, message):
    print(f"{sock.fileno()} : {message}")


def init_master_sock(host, port):
    master_sock = socket(AF_INET, SOCK_STREAM)
    master_sock.bind((host, port))
    master_sock.listen(3)
    return master_sock


def handle_new_client(s, socks):
    conn, addr = s.accept()
    socks.append(conn)
    log(conn, "connected")
    message = conn.recv(1024)
    conn.send(b'Hello client!')


def handle_client_message(s, socks):
    message = s.recv(1024)
    if not message:
        log(s, "disconnected")
        s.close()
        socks.remove(s)
        return

    s.send(b'Hello client!')


def main():
    master_sock = init_master_sock('127.0.0.1', 9998)
    socks = [master_sock]

    while True:
        socks_to_read, _, _ = select.select(
            socks, [], [],
        )

        for s in socks_to_read:
            if s == master_sock:
                handle_new_client(s, socks)
            else:
                handle_client_message(s, socks)


if __name__ == "__main__":
    main()
