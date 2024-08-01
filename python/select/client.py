import socket
import time

def main():
    # Create & connect socket
    s = socket.socket()
    port = 9998
    address = '127.0.0.1'
    s.connect((address, port))

    try:
        while True:
            # Send data to server (newline terminated)
            to_send = b'Message from client\n'
            s.send(to_send)

            # Receive data from server
            received = s.recv(1024)
            print("From server", received)
            
            time.sleep(1)
    except:
        s.close()


if __name__ == "__main__":
    main()
