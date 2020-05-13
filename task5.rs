use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;
use std::collections::HashSet;

fn part1() -> std::io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let mut nice_count = 0;

    'outer: for line in reader.lines() {

        let mut vowel_count = 0;
        let mut prev_c = '#';
        let mut contains_double_letter = false;

        for c in line.unwrap().chars() {
            match c {
                'a' | 'e' | 'i' | 'o' | 'u' => vowel_count += 1,
                _ => () 
            }

            if prev_c == c {
                contains_double_letter = true;
            }

            match (prev_c, c) {
                ('a', 'b') | ('c', 'd') | ('p', 'q') | ('x', 'y') => continue 'outer,
                _ => ()
            }

            prev_c = c;
        }

        if vowel_count >= 3 && contains_double_letter {
            nice_count += 1;
        }
    }

    println!("Part 1: {}", nice_count);
    
    Ok(())
}

fn part2() -> std::io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let mut nice_count = 0;

    for line in reader.lines() {

        let mut prev_prev_c = '#';
        let mut prev_c = '#';

        let mut letter_pairs : HashSet<(char, char)> = HashSet::new();
        let mut last_letter_pair = ('#', '#');
        let mut contains_letter_pair_twice = false;
        let mut contains_double_letter_with_random_letter_inbetween = false;

        for c in line.unwrap().chars() {

            if prev_prev_c == c {
                contains_double_letter_with_random_letter_inbetween = true;
            }

            let curr_pair = (prev_c, c);

            // handle the overlapping rule, i.e. "aaa" does not count as 2 "aa" pairs
            if !(curr_pair == last_letter_pair) {
                if !letter_pairs.insert(curr_pair) {
                    contains_letter_pair_twice = true;
                }

                last_letter_pair = curr_pair;
            } else {
                // set last letter pair to default to allow valid strings like "aaaa"
                last_letter_pair = ('#', '#');
            }

            prev_prev_c = prev_c;
            prev_c = c;
        }

        if contains_letter_pair_twice && contains_double_letter_with_random_letter_inbetween {
            nice_count += 1;
        }
    }

    println!("Part 2: {}", nice_count);
    
    Ok(())
}

fn main() -> std::io::Result<()> {

    part1()?;
    part2()?;

    Ok(())
}