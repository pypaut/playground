use std::{
    io::{Read, Write},
    net::{TcpListener, TcpStream},
    str,
};

fn main() {
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();

    for stream in listener.incoming() {
        let stream = stream.unwrap();

        println!("handle connection...");
        handle_connection(stream);
    }
}

fn handle_connection(mut stream: TcpStream) {
    let mut buffer: &mut [u8] = &mut [0; 512];
    stream.read(&mut buffer).unwrap();
    let message = str::from_utf8(buffer);

    println!("Message from client : {}", message.unwrap());

    let response = "HTTP/1.1 200 OK\r\n\r\n";

    println!("sending response...");
    stream.write_all(response.as_bytes()).unwrap();
}
