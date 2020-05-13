use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;

struct Box {
    l: u32,
    w: u32,
    h: u32,
}

fn part1 (boxes: &Vec<Box>) {
    let mut total = 0;
    for b in boxes {
        let &smallest = [b.l * b.w, b.l * b.h, b.w * b.h].iter().min().unwrap();
        let area = 2 * b.l * b.w + 2 * b.l * b.h + 2 * b.w * b.h;
        total += area + smallest;
    }

    println!("Part 1: {}", total);
}

fn part2 (boxes: &Vec<Box>) {
    let mut total = 0;
    for b in boxes {
        let mut sorted = vec![b.l, b.w, b.h];
        sorted.sort();
        total += 2 * sorted[0] + 2 * sorted[1] + b.l * b.w * b.h;
    }

    println!("Part 2: {}", total);
}

fn main() -> std::io::Result<()>{
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let mut boxes : Vec<Box> = vec![];
    
    for line in reader.lines() {
        let l = line.unwrap();
        let dims : Vec<u32> = l.split("x").map(|s| s.parse::<u32>().unwrap()).collect();
        boxes.push(Box { l: dims[0], w: dims[1], h: dims[2] });
    }

    part1(&boxes);
    part2(&boxes);

    Ok(())
}