use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;
use std::collections::HashMap;

enum BinaryOp {
    AND,
    OR,
    RS,
    LS
}

struct BinaryGate {
    lhs: String,
    rhs: String,
    op: BinaryOp
}

struct NotGate {
    input: String
}

struct Wire {
    input: String
}

enum Part {
    BG(BinaryGate),
    NG(NotGate),
    WIRE(Wire)
}

fn compute_output_signal_of_id(id: &str, parts_by_ids: &HashMap<String, Part>, outputs_by_ids: &mut HashMap<String, u16>) -> u16 {
    match id.parse::<u16>() {
        Ok(val) => val,
        Err(_) => {
            if let Some(&val) = outputs_by_ids.get(id) {
                return val;
            }

            let val = parts_by_ids.get(id).unwrap().compute_output_signal(parts_by_ids, outputs_by_ids);

            outputs_by_ids.insert(String::from(id), val);
            val
        }
    }
}

impl Part {
    fn compute_output_signal(&self, parts_by_ids: &HashMap<String, Part>, outputs_by_ids: &mut HashMap<String, u16>) -> u16 {
        match self {
            Part::BG(bin_gate) => {
                let lhs_val = compute_output_signal_of_id(&bin_gate.lhs, parts_by_ids, outputs_by_ids);
                let rhs_val = compute_output_signal_of_id(&bin_gate.rhs, parts_by_ids, outputs_by_ids);
                match bin_gate.op {
                    BinaryOp::AND => lhs_val & rhs_val,
                    BinaryOp::OR => lhs_val | rhs_val,
                    BinaryOp::RS => lhs_val >> rhs_val,
                    BinaryOp::LS => lhs_val << rhs_val
                }
            },
            Part::NG(not_gate) => {
                !compute_output_signal_of_id(&not_gate.input, parts_by_ids, outputs_by_ids)
            },
            Part::WIRE(w) => {
                compute_output_signal_of_id(&w.input, parts_by_ids, outputs_by_ids)
            }
        }
    }
}

fn main() -> std::io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);
        
    let mut parts_by_ids : HashMap<String, Part> = HashMap::new();

    for line in reader.lines() {
        let l = line.unwrap();
        let split : Vec<&str> = l.split(" ").collect();
        let part : Part = match split.len() {
            3 => Part::WIRE(Wire { input : String::from(split[0]) }),
            4 => Part::NG(NotGate { input : String::from(split[1]) }),
            5 => Part::BG(BinaryGate {
                lhs : String::from(split[0]),
                rhs : String::from(split[2]),
                op : match split[1] {
                    "AND" => BinaryOp::AND,
                    "OR" => BinaryOp::OR,
                    "RSHIFT" => BinaryOp::RS,
                    "LSHIFT" => BinaryOp::LS,
                    _ => panic!("Unsupported binary operator")
                }
            }),
            _ => panic!("Unexpected input")
        };

        parts_by_ids.insert(String::from(split[split.len() - 1]), part);
    }

    let mut outputs_by_ids_1 : HashMap<String, u16> = HashMap::new();
    let final_out_1 = compute_output_signal_of_id("a", &parts_by_ids, &mut outputs_by_ids_1);
    println!("Part 1: {}", final_out_1);

    let mut outputs_by_ids_2 : HashMap<String, u16> = HashMap::new();
    outputs_by_ids_2.insert(String::from("b"), final_out_1);
    let final_out_2 = compute_output_signal_of_id("a", &parts_by_ids, &mut outputs_by_ids_2);
    println!("Part 2: {}", final_out_2);

    Ok(())
}