let headers = ["name","mail",'number']
let records =  [
  ["Alice Johnson", "alice@example.com", "9876543210"],
  ["Bob Smith", "bob@example.com", "9123456780"],
  ["Charlie Ray", "charlie@example.com", "9988776655"],
  ["Diana Prince", "diana@example.com", "9090909090"],
  ["Ethan Hunt", "ethan@example.com", "9001122334"]
];
let transformedRecords = []
for(let record of records){
    let newRecord = record.reduce((acc,value,index)=>{
        acc[headers[index]] = value
        return acc
    },{})
    transformedRecords.push(newRecord)
}

console.log(transformedRecords)