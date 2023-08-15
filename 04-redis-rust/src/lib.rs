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
    fn hostGetEnv(ptr: u64) -> u64;
}

pub fn get_env(name: String) -> String {
    // copy the name of the environment variable to the shared memory
    let mut variable_name: Memory = extism_pdk::Memory::new(name.len());
    variable_name.store(name);

    // call the host function
    let offset: u64 = unsafe { hostGetEnv(variable_name.offset) };

    // read the value of the result from the shared memory
    let variable_value: Memory = extism_pdk::Memory::find(offset).unwrap();

    // return the value
    return variable_value.to_string().unwrap()
}

#[derive(Serialize, Deserialize, Debug)]
struct RedisClientArguments {
    pub id: String,
    pub uri: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct RedisArguments {
    pub id: String,
    pub key: String,
    pub value: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct StringResult {
    pub success: String,
    pub failure: String,
}

#[derive(Error, Debug)]
pub enum RedisError {
    #[error("Redis Client issue")]
    ClientFailure,
    #[error("Store issue")]
    StoreFailure,
    #[error("Not found")]
    NotFound,
}

extern "C" {
    fn hostInitRedisClient(offset: u64) -> u64;
}

pub fn init_redis_client(redis_client_id: String, redis_uri: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "uri": "redis uri"
	// }
    let args = RedisClientArguments {
        id: redis_client_id,
        uri: redis_uri,
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostInitRedisClient(memory_json_str.offset) };

    // Get result from the shared memory
	// The host function (hostInitRedisClient) returns a JSON buffer:
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
        return Err(RedisError::ClientFailure.into());
    }
}

extern "C" {
    fn hostRedisSet(offset: u64) -> u64;
}

pub fn redis_set(redis_client_id: String, key: String, value: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name",
	//    "value": "Bob Morane"
	// }
    let args = RedisArguments {
        id: redis_client_id, 
        key: key,
        value: value,
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostRedisSet(memory_json_str.offset) };

    // Get result from the shared memory
	// The host function (hostRedisSet) returns a JSON buffer:
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
        return Err(RedisError::StoreFailure.into());
    }
}

extern "C" {
    fn hostRedisGet(offset: u64) -> u64;
}

pub fn redis_get(redis_client_id: String, key: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name",
	//    "value": ""
	// }
    let args = RedisArguments {
        id: redis_client_id, 
        key: key,
        value: String::new(),
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostRedisGet(memory_json_str.offset) };

    // Get result (the value associated to the key) from shared memory
	// The host function (hostRedisGet) returns a JSON buffer:
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
        return Err(RedisError::NotFound.into());
    }
    
}

extern "C" {
    fn hostRedisDel(offset: u64) -> u64;
}

pub fn redis_del(redis_client_id: String, key: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name",
	//    "value": ""
	// }
    let args = RedisArguments {
        id: redis_client_id, 
        key: key,
        value: String::new(),
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostRedisDel(memory_json_str.offset) };

    // Get result (the value associated to the key) from shared memory
	// The host function (hostRedisDel) returns a JSON buffer:
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
        return Err(RedisError::NotFound.into());
    }
    
}

extern "C" {
    fn hostRedisFilter(offset: u64) -> u64;
}

pub fn redis_filter(redis_client_id: String, key: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name",
	//    "value": ""
	// }
    let args = RedisArguments {
        id: redis_client_id, 
        key: key,
        value: String::new(),
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostRedisFilter(memory_json_str.offset) };

    // Get result (the value associated to the key) from shared memory
	// The host function (hostRedisDel) returns a JSON buffer:
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
        return Err(RedisError::NotFound.into()); //???
    }
    
}

#[plugin_fn]
pub fn hello(_: String) -> FnResult<u64> {

    let redis_uri : String = get_env("REDIS_URI".to_string());
    let redis_client : Result<String, Error> = init_redis_client("redisDb".to_string(), redis_uri);

    match redis_client {
        Ok(value) => print("🦀 redis client: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_set("redisDb".to_string(), "100".to_string(), "Huey".to_string()) {
        Ok(value)  => print("🦀 saved value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }
    match redis_set("redisDb".to_string(), "200".to_string(), "Dewey".to_string()) {
        Ok(value)  => print("🦀 saved value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }
    match redis_set("redisDb".to_string(), "300".to_string(), "Louie".to_string()) {
        Ok(value)  => print("🦀 saved value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_get("redisDb".to_string(), "100".to_string()) {
        Ok(value)  => print("🦀 value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_get("redisDb".to_string(), "200".to_string()) {
        Ok(value)  => print("🦀 value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_get("redisDb".to_string(), "300".to_string()) {
        Ok(value)  => print("🦀 value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_get("redisDb".to_string(), "400".to_string()) {
        Ok(value)  => print("🦀 value: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_del("redisDb".to_string(), "200".to_string()) {
        Ok(value)  => print("🦀 deleted key: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match redis_filter("redisDb".to_string(), "*00".to_string()) {
        Ok(value)  => print("🦀 keys: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    Ok(0)
}
