extern crate grpc;
extern crate ping_rust_grpc;
extern crate clap;

use std::time::Duration;
use std::net::ToSocketAddrs;

use clap::{Arg, App};
use grpc::ClientConf;

use ping_rust_grpc::*;

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

        // By default, there's no connection timeout. Make sure we exit if we're
        // struggling to connect.
        let mut client_config = grpc::ClientConf::new();
        client_config.http.connection_timeout = Some(Duration::from_millis(1000));

        let client = PingClient::new_plain(&addr.ip().to_string(), addr.port(), client_config).unwrap();
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