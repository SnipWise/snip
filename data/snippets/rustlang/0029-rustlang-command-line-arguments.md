## Command Line Arguments
Parsing command line arguments with clap
```rust
use clap::{Arg, App};

fn main() {
    let matches = App::new("My App")
        .version("1.0")
        .author("Your Name")
        .about("Does awesome things")
        .arg(Arg::with_name("name")
            .short("n")
            .long("name")
            .value_name("NAME")
            .help("Sets a custom name")
            .takes_value(true)
            .default_value("World"))
        .arg(Arg::with_name("verbose")
            .short("v")
            .long("verbose")
            .help("Enable verbose output"))
        .get_matches();
    
    let name = matches.value_of("name").unwrap();
    let verbose = matches.is_present("verbose");
    
    if verbose {
        println!("Verbose mode enabled");
        println!("Name parameter: {}", name);
    }
    
    println!("Hello, {}!", name);
}
```