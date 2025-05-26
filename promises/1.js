let p1 = new Promise((resolve,reject)=>{
    resolve()
})

p1.then((res)=>{
    console.log(res)
    // return "dsfds"
})
.then((res)=>{
    console.log(res)
    return "dfdfdfd"
}).then((res)=>{
    console.log(res)
    // return "dfdfdfd"
    throw Error("my error")
}).catch((err)=>{
    console.log(err)
    return "hello"
}).then((res)=>{
    console.log(res)
}).catch((err)=>{
    console.log(err)
})