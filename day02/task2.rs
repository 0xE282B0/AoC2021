// copied from https://doc.rust-lang.org/rust-by-example/std_misc/file/read_lines.html
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    // File hosts must exist in current path before this produces output
    if let Ok(lines) = read_lines("./data/input.txt") {
        // Consumes the iterator, returns an (Optional) String
        let mut forward = 0;
        let mut depth = 0;
        let mut aim = 0;

        for line in lines {
            if let Ok(mv) = line {
                if mv.starts_with("forward"){
                    let val = mv.chars().skip(7).take(mv.chars().count()-7).collect::<String>();
                    let x = val.trim().parse::<i32>().unwrap();
                    forward = forward + x;
                    depth = depth + aim * x;
                } else if mv.starts_with("down"){
                    let val = mv.chars().skip(4).take(mv.chars().count()-4).collect::<String>();
                    aim = aim + val.trim().parse::<i32>().unwrap();
                } else if mv.starts_with("up"){
                    let val = mv.chars().skip(2).take(mv.chars().count()-2).collect::<String>();
                    aim = aim - val.trim().parse::<i32>().unwrap();
                }
            }
        }
        println!("forward: {}", forward);      
        println!("depth: {}", depth);
        println!("mul {}", depth*forward);
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
