use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;
use std::collections::HashSet;

fn part1(input : &String) {
    let mut visited : HashSet<(i32, i32)> = HashSet::new();

    let mut position = (0, 0);

    for chr in input.chars() {
        match chr {
            '^' => position.1 += 1,
            'v' => position.1 -= 1,
            '<' => position.0 += 1,
            '>' => position.0 -= 1,
            _ => ()
        }

        visited.insert(position);
    }

    println!("Part 1: {}", visited.len());
}

fn part2(input : &String) {
    let mut visited : HashSet<(i32, i32)> = HashSet::new();

    let mut pos_santa = (0, 0);
    let mut pos_robo = (0, 0);

    let mut santa_turn = true;

    for chr in input.chars() {
        if santa_turn {
            match chr {
                '^' => pos_santa.1 += 1,
                'v' => pos_santa.1 -= 1,
                '<' => pos_santa.0 += 1,
                '>' => pos_santa.0 -= 1,
                _ => ()
            }
            visited.insert(pos_santa);
        } else {
            match chr {
                '^' => pos_robo.1 += 1,
                'v' => pos_robo.1 -= 1,
                '<' => pos_robo.0 += 1,
                '>' => pos_robo.0 -= 1,
                _ => ()
            }
            visited.insert(pos_robo);
        }

        santa_turn = !santa_turn;
    }

    println!("Part 2: {}", visited.len());
}

fn main() -> std::io::Result<()>{
    let file = File::open("input.txt")?;
    let mut reader = BufReader::new(file);
    let mut input = String::new();
    reader.read_to_string(&mut input)?;
    
    part1(&input);
    part2(&input);

    Ok(())
}