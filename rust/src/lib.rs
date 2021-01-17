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
    pub name: String,
    pub description: String,
    pub op: Vec<Operation>,
}

#[derive(Debug, Clone)]
pub struct Operation {
    pub name: String,
    pub description: String,
    pub params: Vec<String>,
    pub pairs: Vec<String>,
    pub results: Vec<String>,
}

#[derive(Debug, Clone, Default)]
pub struct Operations {
    pub interfaces: Vec<Interface>,
    pub fields: Vec<Field>,
}

#[derive(Debug, Clone)]
pub struct Field {
    pub name: String,
    pub ty: String,
}

#[derive(Debug, Clone, Default)]
pub struct Service {
    pub name: String,
    pub namespaces: Vec<Namespace>,
    pub pairs: Pairs,
    pub infos: Infos,
}

#[derive(Debug, Clone, Default)]
pub struct Namespace {
    pub name: String,
    pub implement: Vec<String>,
    pub new: New,
    pub op: Vec<Op>,
}

#[derive(Debug, Clone, Default)]
pub struct New {
    pub required: Vec<String>,
    pub optional: Vec<String>,
}

#[derive(Debug, Clone, Default)]
pub struct Op {
    pub name: String,
    pub required: Vec<String>,
    pub optional: Vec<String>,
}

pub type Pairs = Vec<Pair>;

#[derive(Debug, Clone)]
pub struct Pair {
    pub name: String,
    pub ty: String,
    pub description: String,
}

pub type Infos = Vec<Info>;

#[derive(Debug, Clone)]
pub struct Info {
    pub scope: String,
    pub category: String,
    pub name: String,
    pub ty: String,
    pub export: bool,
    pub description: String,
}

pub fn parse_services(file_path: &str) -> Result<Service> {
    parse::parse_service(file_path)
}
