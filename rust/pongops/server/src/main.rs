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
        let buffer_str = str::from_utf8(buffer).unwrap();
        let buffer_str_trimmed = buffer_str.trim_matches(char::from(0));
        let dir: Vec2d = serde_json::from_str(buffer_str_trimmed).unwrap();

        // Update dir
        position.x += dir.x;
        position.y += dir.y;

        // Send position
        let serialized = serde_json::to_string(&position).unwrap();
        let buffer = serialized.as_bytes();
        stream.write_all(buffer).unwrap();
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::thread::{sleep, spawn};
    use std::time::Duration;

    #[test]
    fn test_server() {
        // Run server
        let _ = spawn( || {
            serve();
        });

        // Wait for server to listen
        sleep(Duration::from_millis(100));

        // Connect to server
        let mut stream = TcpStream::connect(("127.0.0.1", 7878)).unwrap();

        // Send direction and check new position
        let directions = [
            Vec2d { x: 0., y: 0. }, Vec2d { x: 0., y: 1. }, Vec2d { x: -1., y: -1. },
        ];
        let mut expected_pos = Vec2d { x: 0., y: 0. };

        for dir in directions.iter() {
            // Send direction
            let serialized = serde_json::to_string(dir).unwrap();
            let buffer = serialized.as_bytes();
            stream.write_all(buffer).unwrap();

            // Receive position
            let mut buffer: &mut [u8] = &mut [0; 512];
            stream.read(&mut buffer).unwrap();
            let buffer_str = str::from_utf8(&buffer).unwrap().to_string();
            let buffer_str_trimmed = buffer_str.trim_matches(char::from(0));

            // Deserialize to pos
            let pos: Vec2d = serde_json::from_str(&buffer_str_trimmed).unwrap();

            // Update expected position
            expected_pos.x += dir.x;
            expected_pos.y += dir.y;

            // Test values
            assert_eq!(pos.x, expected_pos.x);
            assert_eq!(pos.y, expected_pos.y);
        }
    }
}