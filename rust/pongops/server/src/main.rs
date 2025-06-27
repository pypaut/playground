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

        println!("handle connection...");
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
        println!("Message from client: {}", message);

        let dir: Vec2d = serde_json::from_str(&message[..17]).unwrap();
        println!("{:?}", dir);

        // Update dir
        position.x += dir.x;
        position.y += dir.y;

        // Send position
        let serialized = serde_json::to_string(&dir).unwrap();
        let buffer = serialized.as_bytes();
        stream.write_all(buffer).unwrap();
    }
}
