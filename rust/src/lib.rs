use anyhow::Result;
use lazy_static::lazy_static;

mod parse;

lazy_static! {
    pub static ref PARSED_PAIRS: Pairs = parse::parse_pairs().unwrap();
    pub static ref PARSED_INFOS: Infos = parse::parse_infos().unwrap();
    pub static ref PARSED_OPERATIONS: Operations = parse::parse_operations().unwrap();
}

#[derive(Debug, Clone)]
pub struct Interface {
    name: String,
    description: String,
    op: Vec<Operation>,
}

#[derive(Debug, Clone)]
pub struct Operation {
    name: String,
    description: String,
    params: Vec<String>,
    pairs: Vec<String>,
    results: Vec<String>,
}

#[derive(Debug, Clone, Default)]
pub struct Operations {
    interfaces: Vec<Interface>,
    fields: Vec<Field>,
}

#[derive(Debug, Clone)]
pub struct Field {
    name: String,
    ty: String,
}

#[derive(Debug, Clone, Default)]
pub struct Service {
    name: String,
    namespaces: Vec<Namespace>,
    pairs: Pairs,
    infos: Infos,
}

#[derive(Debug, Clone, Default)]
pub struct Namespace {
    name: String,
    implement: Vec<String>,
    new: New,
    op: Vec<Op>,
}

#[derive(Debug, Clone, Default)]
pub struct New {
    required: Vec<String>,
    optional: Vec<String>,
}

#[derive(Debug, Clone, Default)]
pub struct Op {
    name: String,
    required: Vec<String>,
    optional: Vec<String>,
}

pub type Pairs = Vec<Pair>;

#[derive(Debug, Clone)]
pub struct Pair {
    name: String,
    ty: String,
    description: String,
}

pub type Infos = Vec<Info>;

#[derive(Debug, Clone)]
pub struct Info {
    scope: String,
    category: String,
    name: String,
    ty: String,
    export: bool,
    description: String,
}

pub fn parse_services(file_path: &str) -> Result<Service> {
    parse::parse_service(file_path)
}
