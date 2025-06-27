use std::{
    io::{Read, Write},
    net::{TcpListener, TcpStream},
    str,
};

use serde_json;

use lib::Vec2d;

fn main() {
    serve();
}

fn serve() {
    println!("Running server");
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();

    for stream in listener.incoming() {
        let stream = stream.unwrap();

        handle_connection(stream);
    }
}

fn handle_connection(mut stream: TcpStream) {
    let mut position = Vec2d { x: 0., y: 0. };

    loop {
        // Receive direction
        let mut buffer: &mut [u8] = &mut [0; 512];
        stream.read(&mut buffer).unwrap();
        let message = str::from_utf8(buffer).unwrap();

        let dir: Vec2d = serde_json::from_str(&message[..17]).unwrap();

        // Update dir
        position.x += dir.x;
        position.y += dir.y;

        // Send position
        let serialized = serde_json::to_string(&dir).unwrap();
        let buffer = serialized.as_bytes();
        stream.write_all(buffer).unwrap();
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::thread::{sleep, spawn};
    use std::time::Duration;

    // TODO : a single test actually tries the connection
    // Other tests will require better refactorization for clear unit testing

    #[test]
    fn position_is_updated() {
        // Run server
        let _ = spawn( || {
            serve();
        });

        // Wait for server to listen
        sleep(Duration::from_millis(100));

        // Connect to server
        let mut stream = TcpStream::connect(("127.0.0.1", 7878)).unwrap();

        // Send direction
        let dir = Vec2d { x: 0., y: 1. };
        let serialized = serde_json::to_string(&dir).unwrap();
        let buffer = serialized.as_bytes();
        stream.write_all(buffer).unwrap();

        // Receive position
        let mut buffer: &mut [u8] = &mut [0; 512];
        stream.read(&mut buffer).unwrap();
        let buffer_str = str::from_utf8(&buffer).unwrap()[..17].to_string();

        // Deserialize to pos
        let pos: Vec2d = serde_json::from_str(&buffer_str).unwrap();

        // Test values
        let init_pos = Vec2d { x: 0., y: 0. };
        assert_eq!(pos.x, init_pos.x + dir.x);
        assert_eq!(pos.y, init_pos.y + dir.y);
    }
}