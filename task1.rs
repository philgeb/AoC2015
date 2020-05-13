use std::path::Path;
use std::io::prelude::*;
use std::fs::File;

fn part1(input: &String) {
    let mut floor = 0;
    for chr in input.chars() {
        if chr == '(' {
            floor += 1;
        } else {
            floor -= 1;
        }
    }

    println!("Part 1: {}", floor);
}

fn part2(input: &String) {
    let mut floor = 0;
    let mut index = 1;
    for chr in input.chars() {
        // to mix it up, using "match" here instead of "if else"
        match chr {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => (),
        }

        if floor == -1 {
            println!("Part 2: {}", index);
            return;
        }

        index += 1;
    }
}

fn main() -> std::io::Result<()>{
    let path = Path::new("input.txt");
    let mut file = File::open(&path)?;
    let mut content = String::new();
    file.read_to_string(&mut content)?;

    part1(&content);
    part2(&content);

    Ok(())
}