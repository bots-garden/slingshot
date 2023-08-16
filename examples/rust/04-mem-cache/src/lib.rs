use extism_pdk::*;
use serde::{Serialize, Deserialize};
use thiserror::Error;


extern "C" {
    fn hostPrint(ptr: u64) -> u64;
}

pub fn print(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrint(memory_text.offset) };
}

extern "C" {
    fn hostMemorySet(offset: u64) -> u64;
}

#[derive(Serialize, Deserialize, Debug)]
struct MemArguments {
    pub key: String,
    pub value: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct StringResult {
    pub success: String,
    pub failure: String,
}

#[derive(Error, Debug)]
pub enum MemError {
    #[error("Store issue")]
    StoreFailure,
    #[error("Not found")]
    NotFound,
}

pub fn memory_set(key: String, value: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "key": "name",
	//    "value": "Bob Morane"
	// }
    let record = MemArguments {
        key: key,
        value: value,
    };
    let json_str: String = serde_json::to_string(&record).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostMemorySet(memory_json_str.offset) };

    // Get result from the shared memory
	// The host function (hostMemorySet) returns a JSON buffer:
	// {
    //   "success": "the value associated to the key",
	//   "failure": "error message if error, else empty"
	// }
    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
    
    let json_string:String = memory_result.to_string().unwrap();
    let result: StringResult = serde_json::from_str(&json_string).unwrap();

    if result.failure.is_empty()  {
        return Ok(result.success);
    } else {
        return Err(MemError::StoreFailure.into());
    }
}

extern "C" {
    fn hostMemoryGet(offset: u64) -> u64;
}

pub fn memory_get(key: String) -> Result<String, Error> {
	// Copy argument to memory
    let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
    memory_key.store(key);

	// Call the host function
    let offset: u64 = unsafe { hostMemoryGet(memory_key.offset) };

    // Get result (the value associated to the key) from shared memory
	// The host function (hostMemoryGet) returns a JSON buffer:
	// {
    //   "success": "the value associated to the key",
	//   "failure": "error message if error, else empty"
	// }
    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
    let json_string:String = memory_result.to_string().unwrap();
    let result: StringResult = serde_json::from_str(&json_string).unwrap();
    
    if result.failure.is_empty()  {
        return Ok(result.success);
    } else {
        return Err(MemError::NotFound.into());
    }
    
}


#[plugin_fn]
pub fn hello(_: String) -> FnResult<u64> {

    match memory_set("bob".to_string(), "Bob Morane".to_string()) {
        Ok(value)  => print("🦀 saved value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match memory_get("bob".to_string()) {
        Ok(value)  => print("🦀 value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match memory_get("sam".to_string()) {
        Ok(value)  => print("🦀 value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }
    
    Ok(0)
}
