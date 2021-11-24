use quicli::prelude::*;
use serde::{Deserialize, Serialize};
use serde_json;
use structopt::StructOpt;

#[derive(Serialize, Deserialize)]
struct Api {
    version: String,
}

#[derive(Serialize, Deserialize)]
struct Switch {
    id: i64,
    name: String,
    site: String,
    rack: String,
    status: String,
    netbox_url: String,
    position: String,
    primary_ip4: String,
}

#[derive(Serialize, Deserialize)]
struct Switches {
    switches: Vec<Switch>,
}

#[derive(Serialize, Deserialize)]
struct Interface {
    id: i64,
    // We cannot directly call the field "type", as it is a reserved Rust word
    #[serde(rename = "type")]
    ty: String,
    name: String,
    device_name: String,
    device_id: i64,
    mac_address: String,
    enabled: bool,
}

#[derive(Serialize, Deserialize)]
struct Interfaces {
    interfaces: Vec<Interface>,
}

#[derive(Debug, StructOpt)]
struct Cli {
    #[structopt(long = "switch")]
    switch: bool,
}

fn main() -> CliResult {
    let api_url = "http://0.0.0.0:8080/api";
    let mut total_url = format!("{}", api_url);

    let args = Cli::from_args();
    if args.switch {
        total_url += "/switch";
        let response: Switches = reqwest::blocking::get(total_url)?.json()?;
        println!("{}", serde_json::to_string(&response)?);
    } else {
        let response: Api = reqwest::blocking::get(total_url)?.json()?;
        println!("{}", serde_json::to_string(&response)?);
    }

    Ok(())
}
