extern crate grpc;
extern crate ping_rust_grpc;
extern crate clap;

use ping_rust_grpc::*;

use clap::{Arg, App};
use std::net::{ToSocketAddrs};

fn main() {
    let matches = App::new("ping client")
        .arg(Arg::with_name("address")
            .short("a")
            .long("address")
            .value_name("address")
            .help("grpc server address (IPv4)")
            .takes_value(true))
        .arg(Arg::with_name("MESSAGE")
            .help("the message to send in the ping")
            .index(1))
        .get_matches();

    let message = matches.value_of("message").unwrap_or("");
    let address = matches.value_of("address").unwrap_or("127.0.0.1:50051" );

    // Resolve any addresses to hostnames if necessary.
    let addresses = address.to_socket_addrs()
        .unwrap_or_else(|err| {
            eprintln!("failed to resolve addresss {:?}: {}", address, err);
            std::process::exit(1)
        });

    // If it resolves to multiple IP addresses, try until it succeeds. For example,
    // localhost will resolve to [::1]:50051 and  127.0.0.1:50051.car
    for addr in addresses {
        let mut req = PingRequest::new();
        req.set_message(message.to_string());

        println!("dialling {}...", &addr.to_string());
        println!("ping message: {:?}", req.message);
        let client = PingClient::new_plain(&addr.ip().to_string(), addr.port(), Default::default()).unwrap();
        let resp = client.ping(grpc::RequestOptions::new(), req);
        
        match resp.wait() {
            Ok((_, reply, _)) => {
                println!("ping reply: {:?}", reply.message);
                break;
            },
            Err(err) => eprintln!("received error {:?}", err),
        }
    }
}