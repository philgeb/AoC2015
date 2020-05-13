use md5;

const KEY : &str = "ckczppom";

fn md5_hex(input : &str) -> String {
    let digest = md5::compute(input);
    format!("{:x}", digest)
}

fn main() -> std::io::Result<()> {
    let mut count = 0;

    let mut_key : String = KEY.to_owned();
    loop {
        let combined_key = mut_key.clone() + &count.to_string();
        let hash = md5_hex(&combined_key);

        // for part 1: 5 zeros, for part 2: 6
        if hash.starts_with("000000") {
            println!("Number: {}", count);
            break;
        }

        count += 1;
    };

    Ok(())
}