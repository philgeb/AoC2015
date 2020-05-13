use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;
use regex::Regex;
use std::cmp::max;

enum Action {
    ON,
    OFF,
    TOGGLE
}

struct Instr {
    act: Action,
    pos1: (u32, u32),
    pos2: (u32, u32)
}

fn part1(instructions: &Vec<Instr>) {
    let mut grid = vec![vec![false; 1000]; 1000];

    for ins in instructions {
        for x in ins.pos1.0 as usize ..= ins.pos2.0 as usize {
            for y in ins.pos1.1 as usize ..= ins.pos2.1 as usize {
                grid[x][y] = match ins.act {
                    Action::ON => true,
                    Action::OFF => false,
                    Action::TOGGLE => !grid[x][y]
                }
            }
        }
    }

    let mut lights_on = 0;
    for x in 0..1000 {
        for y in 0..1000 {
            if grid[x][y] {
                lights_on += 1;
            }
        }
    }

    println!("Part 1: {}", lights_on);
}

fn part2(instructions: &Vec<Instr>) {
    let mut grid = vec![vec![0; 1000]; 1000];

    for ins in instructions {
        for x in ins.pos1.0 as usize ..= ins.pos2.0 as usize {
            for y in ins.pos1.1 as usize ..= ins.pos2.1 as usize {
                grid[x][y] = match ins.act {
                    Action::ON => grid[x][y] + 1,
                    Action::OFF => max(grid[x][y] - 1, 0),
                    Action::TOGGLE => grid[x][y] + 2
                }
            }
        }
    }

    let mut total_bright = 0;
    for x in 0..1000 {
        for y in 0..1000 {
            total_bright += grid[x][y];
        }
    }

    println!("Part 2: {}", total_bright);
}


fn main() -> std::io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);
    
    let re = Regex::new(r"(\d+),(\d+) through (\d+),(\d+)").unwrap();
    
    let mut instructions = vec![];
    for line in reader.lines() {
        let l = line.unwrap();

        let caps = re.captures(&l).unwrap();
        let pos1 = (caps[1].parse::<u32>().unwrap(), caps[2].parse::<u32>().unwrap());
        let pos2 = (caps[3].parse::<u32>().unwrap(), caps[4].parse::<u32>().unwrap());

        let act = {
            match l.as_bytes()[6] as char {
                'n' => Action::ON,
                'f' => Action::OFF,
                _ => Action::TOGGLE,
            }
        };

        instructions.push(
            Instr { 
                act: act,
                pos1: pos1,
                pos2: pos2
            });
    }

    part1(&instructions);
    part2(&instructions);

    Ok(())
}