use byte::*;
use byte::ctx::{Str, NULL};
use std::io::prelude::*;
use std::net::TcpStream;

fn main() -> std::io::Result<()> {
    let mut stream = TcpStream::connect("127.0.0.1:7878")?;
    println!("connected");

    let message: &[u8] = b"hello, world!\n\0dump";
    stream.write(message)?;
    let _ = stream.flush();
    println!("message sent");

    let mut buf: [u8;128] = [0; 128];
    stream.read(&mut buf)?;
    println!("message read");

    let offset = &mut 0;
    let str = buf.read_with::<&str>(offset, Str::Delimiter(NULL)).unwrap();

    println!("{}", str);
    Ok(())
} // the stream is closed here
