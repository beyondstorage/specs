use std::collections::HashMap;

use anyhow::Result;
use serde::{Deserialize, Serialize};
use std::fs;

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Field {
    #[serde(rename = "type")]
    ty: String,
}

type Fields = HashMap<String, Field>;

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Info {
    #[serde(rename = "type")]
    ty: String,
    #[serde(default)]
    export: bool,
    #[serde(default)]
    description: String,
}

type Infos = HashMap<String, Info>;

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Pair {
    #[serde(rename = "type")]
    ty: String,
    #[serde(default)]
    description: String,
}

type Pairs = HashMap<String, Pair>;

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Operation {
    #[serde(default)]
    description: String,
    #[serde(default)]
    params: Vec<String>,
    #[serde(default)]
    pairs: Vec<String>,
    #[serde(default)]
    results: Vec<String>,
}

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Interface {
    #[serde(default)]
    description: String,
    op: HashMap<String, Operation>,
}

type Interfaces = HashMap<String, Interface>;

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Op {
    #[serde(default)]
    required: Vec<String>,
    #[serde(default)]
    optional: Vec<String>,
}

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Namespace {
    #[serde(default)]
    implement: Vec<String>,
    new: Op,
    op: HashMap<String, Op>,
}

#[derive(Deserialize, Serialize, Debug, Clone)]
struct Service {
    name: String,
    namespace: HashMap<String, Namespace>,
    #[serde(default)]
    paris: HashMap<String, Pair>,
    #[serde(default)]
    infos: HashMap<String, HashMap<String, HashMap<String, Info>>>,
}

pub fn parse_pairs() -> Result<super::Pairs> {
    let tp: Pairs = toml::from_slice(include_bytes!("definitions/pairs.toml"))?;

    let mut ps = super::Pairs::default();

    for (k, v) in tp.iter() {
        ps.push(super::Pair {
            name: k.clone(),
            ty: v.ty.clone(),
            description: v.description.clone(),
        })
    }

    ps.sort_by_key(|i| i.name.clone());

    Ok(ps)
}

pub fn parse_infos() -> Result<super::Infos> {
    let mut ps = super::Infos::default();

    // Parse object meta.
    let tp: Infos = toml::from_slice(include_bytes!("definitions/info_object_meta.toml"))?;

    for (k, v) in tp.iter() {
        ps.push(super::Info {
            scope: "object".to_string(),
            category: "meta".to_string(),
            name: k.clone(),
            ty: v.ty.clone(),
            export: v.export,
            description: v.description.clone(),
        })
    }

    // Parse storage meta.
    let tp: Infos = toml::from_slice(include_bytes!("definitions/info_storage_meta.toml"))?;

    for (k, v) in tp.iter() {
        ps.push(super::Info {
            scope: "storage".to_string(),
            category: "meta".to_string(),
            name: k.clone(),
            ty: v.ty.clone(),
            export: v.export,
            description: v.description.clone(),
        })
    }

    ps.sort_by_key(|i| i.name.clone());

    Ok(ps)
}

pub fn parse_operations() -> Result<super::Operations> {
    let mut ps = super::Operations::default();

    let interfaces: Interfaces = toml::from_slice(include_bytes!("definitions/operations.toml"))?;
    let fields: Fields = toml::from_slice(include_bytes!("definitions/fields.toml"))?;

    for (name, v) in fields.iter() {
        ps.fields.push(super::Field {
            name: name.clone(),
            ty: v.ty.clone(),
        })
    }

    for (name, v) in interfaces.iter() {
        let mut it = super::Interface {
            name: name.clone(),
            description: v.description.clone(),
            op: vec![],
        };

        for (name, v) in v.op.iter() {
            it.op.push(super::Operation {
                name: name.clone(),
                description: v.description.clone(),
                params: v.params.clone(),
                pairs: v.pairs.clone(),
                results: v.results.clone(),
            })
        }

        ps.interfaces.push(it);
    }

    Ok(ps)
}

pub fn parse_service(file_path: &str) -> Result<super::Service> {
    let mut ps = super::Service::default();

    let service: Service = toml::from_slice(&fs::read(file_path)?)?;

    ps.name = service.name.clone();

    // Parse pairs.
    for (name, v) in service.paris.iter() {
        ps.pairs.push(super::Pair {
            name: name.clone(),
            ty: v.ty.clone(),
            description: v.description.clone(),
        })
    }

    // Parse infos.
    for (scope, v) in service.infos.iter() {
        for (category, v) in v.iter() {
            for (name, v) in v.iter() {
                ps.infos.push(super::Info {
                    scope: scope.clone(),
                    category: category.clone(),
                    name: name.clone(),
                    ty: v.ty.clone(),
                    export: v.export,
                    description: v.description.clone(),
                })
            }
        }
    }

    // Parse namespace.
    for (name, v) in service.namespace.iter() {
        let mut n = super::Namespace {
            name: name.clone(),
            implement: v.implement.clone(),
            new: super::New {
                required: v.new.required.clone(),
                optional: v.new.optional.clone(),
            },
            op: vec![],
        };

        for (op_name, op) in v.op.iter() {
            n.op.push(super::Op {
                name: op_name.clone(),
                required: op.required.clone(),
                optional: op.optional.clone(),
            })
        }

        ps.namespaces.push(n);
    }

    Ok(ps)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse_pairs() -> Result<()> {
        let ps = parse_pairs()?;

        println!("{:?}", ps);

        assert!(ps.len() > 0);

        Ok(())
    }

    #[test]
    fn test_parse_infos() -> Result<()> {
        let ps = parse_infos()?;

        println!("{:?}", ps);

        assert!(ps.len() > 0);

        Ok(())
    }

    #[test]
    fn test_parse_operations() -> Result<()> {
        let ps = parse_operations()?;

        println!("{:?}", ps);

        assert!(ps.interfaces.len() > 0);
        assert!(ps.fields.len() > 0);

        Ok(())
    }

    #[test]
    fn test_parse_services() -> Result<()> {
        let ps = parse_service("testdata/service.toml")?;

        println!("{:?}", ps);

        assert!(ps.name.len() > 0);

        Ok(())
    }
}
