use std::io::{Read, Write};
use std::net::TcpStream;

use serde_json;

use lib::Vec2d;

fn main() {
    // Connect to server
    let mut stream = TcpStream::connect("127.0.0.1:7878").unwrap();

    // Send direction
    let dir = Vec2d { x: 0., y: 1. };
    let serialized = serde_json::to_string(&dir).unwrap();
    let buffer = serialized.as_bytes();
    stream.write_all(buffer).unwrap();

    // Receive position
    let mut buffer: &mut [u8] = &mut [0; 512];
    stream.read(&mut buffer).unwrap();
    let buffer_str = str::from_utf8(&buffer).unwrap()[..17].to_string();

    // Deserialize to dir
    let pos: Vec2d = serde_json::from_str(&buffer_str).unwrap();
    println!("{:?}", pos);
}
