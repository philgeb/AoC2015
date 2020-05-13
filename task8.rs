use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;

fn count_total_and_inmem(line : &String) -> (i32, i32) {
    let mut total_chrs = 0;
    let mut in_mem = 0;
    let mut chrs_iter = line.chars();

    while let Some(curr) = chrs_iter.next() {
        total_chrs += 1;
        in_mem += 1;
        if curr == '\\' {
            total_chrs += 1;
            // we know that the input files are all valid strings, can't have None here
            if chrs_iter.next().unwrap() == 'x' {
                total_chrs += 2;
                // consume the escaped hex value
                chrs_iter.next();
                chrs_iter.next();
            }
        }
    }

    // substract the " chars before and after the string
    in_mem -= 2;

    (total_chrs, in_mem)
}

fn encode(line : &String) -> String {
    let mut encoded = String::new();
    
    encoded.push('"');
    for chr in line.chars() {
        if chr == '"' || chr == '\\' {
            encoded.push('\\');
        }
        encoded.push(chr);
    }
    encoded.push('"');

    encoded
}

fn main() -> std::io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let lines : Vec<String> = reader.lines().map(|l| l.unwrap()).collect();

    // Part 1
    let mut total_chrs = 0;
    let mut in_mem = 0;
    for line in &lines {
        let (line_chrs, line_in_mem) = count_total_and_inmem(&line);
        total_chrs += line_chrs;
        in_mem += line_in_mem;
    }

    println!("Part 1: {}", total_chrs - in_mem);

    // Part2
    let mut total_chrs = 0;
    let mut total_enc_chrs = 0;
    for line in &lines {
        let encoded = encode(&line);
        let (line_chrs, _) = count_total_and_inmem(&line);
        let (line_enc_chrs, _) = count_total_and_inmem(&encoded);
        total_chrs += line_chrs;
        total_enc_chrs += line_enc_chrs;
    }

    println!("Part 2: {}", total_enc_chrs - total_chrs);

    Ok(())
}