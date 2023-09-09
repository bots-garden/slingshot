#![no_main]
use std::collections::HashMap;

use extism_pdk::*;
use json::Value;
//use serde::Serialize;
use serde::{Serialize, Deserialize};


#[derive(Serialize)]
struct Output {
    pub success: String,
    pub failure: String,
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
struct Response {
    pub text_body: String,
    pub json_body: Value, // 🤔
    pub headers : HashMap<String,String>,
    pub status_code: i64
}
/*
{textBody, jsonBody, headers, statusCode}
*/

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Request {
    pub body: String,
    pub base_url: String,
    pub headers : HashMap<String,String>,
    pub method: String
}

extern "C" {
    fn hostPrint(ptr: u64) -> u64;
}

pub fn print(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrint(memory_text.offset) };
}

extern "C" {
    fn hostGetMessage(ptr: u64) -> u64;
}

pub fn get_message(key: String) -> String {
    let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
    memory_key.store(key);
    let offset: u64 = unsafe { hostGetMessage(memory_key.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();

    return memory_result.to_string().unwrap()
}

extern "C" {
    fn hostMemorySet(offset: u64) -> u64;
}

#[derive(Serialize, Deserialize, Debug)]
struct MemRecord {
    pub key: String,
    pub value: String,
}

pub fn memory_set(key: String, value: String) {
    let record = MemRecord {
        key: key,
        value: value,
    };
    let json_str: String = serde_json::to_string(&record).unwrap();

    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);
    let offset: u64 = unsafe { hostMemorySet(memory_json_str.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();

    let msg: String = "🟣 from 🦀 guest: ".to_string() + &memory_result.to_string().unwrap();

    print(msg)
}

extern "C" {
    fn hostMemoryGet(offset: u64) -> u64;
}

#[derive(Serialize, Deserialize, Debug)]
struct StringResult {
    pub success: String,
    pub failure: String,
}

pub fn memory_get(key: String) -> String {

    let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
    memory_key.store(key);
    let offset: u64 = unsafe { hostMemoryGet(memory_key.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
    let json_string:String = memory_result.to_string().unwrap();
    let result: StringResult = serde_json::from_str(&json_string).unwrap();

    // TODO: return Result
    return result.success
    
}

#[plugin_fn]
pub fn handle(input: String) -> FnResult<Json<Response>> {
    // TODO test with bytes?

    print("🟣 this is the wasm handle function".to_string());

    let request : Request = serde_json::from_str(&input).unwrap();

    let val1: String = get_message("hello".to_string());
    print(val1);

    let msg: String = "🦀 Hello ".to_string() + &request.body;

    memory_set("bill".to_string(), "🦊 ballantines".to_string());

    print(memory_get("bill".to_string()));

    let mut headers: HashMap<String, String> = HashMap::new();
    headers.insert("Content-Type".to_string(), "text/plain; charset=utf-8".to_string());
    headers.insert("X-Slingshot-version".to_string(), "0.0.0".to_string());


    let response : Response = Response { 
        text_body: msg, 
        json_body: serde_json::from_str("{}")?, 
        headers , 
        status_code: 200 
    };

    //let output = Output { success: msg, failure: "no error".to_string() };
    
    Ok(Json(response))
}

#[plugin_fn]
pub fn _start(_: String) -> FnResult<String> {
    print("hello from _start".to_string());
    Ok("hello".to_string())
}