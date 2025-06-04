function throttler() {
  let lastCall;
  let delay = 1000;

  return async function heavyComp() {
      let currentTIme = new Date().getTime();
      console.log(lastCall,currentTIme) ;

    if (!lastCall) {
      lastCall = currentTIme - 2000;
    }
    // console.log("hell")
    console.log(currentTIme - lastCall, lastCall);
    if (currentTIme - lastCall > delay) {
      let promise = new Promise((resolve, reject) => {
        setTimeout(() => {
          resolve("computation done");
        }, 1000);
      });

      let result = await promise;
      lastCall = currentTIme;
      console.log(result);
    }
  };

  console.log(currentTIme);
  // if(currentTIme - )
}

// let throFunction = throttler();
// throFunction();

// setTimeout(throFunction, 1500);
let date = new Date()
console.log(date)
console.log(date.getUTCMinutes())
// console.log(date.getMinutes.)



// let promise = new Promise((resolve, reject) => {
  
// })

let obj = {
  "id": 101,
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "isActive": true,
  "address": {
    "street": "123 Maple Street",
    "city": "Springfield",
    "state": "IL",
    "postalCode": "62704"
  },
  "orders": [
    {
      "orderId": "A001",
      "date": "2025-05-20",
      "total": 59.99,
      "items": [
        {
          "productId": "P1001",
          "name": "Wireless Mouse",
          "quantity": 1,
          "price": 25.99
        },
        {
          "productId": "P1002",
          "name": "USB Keyboard",
          "quantity": 1,
          "price": 34.00
        }
      ]
    },
    {
      "orderId": "A002",
      "date": "2025-06-01",
      "total": 99.50,
      "items": [
        {
          "productId": "P2001",
          "name": "HD Monitor",
          "quantity": 1,
          "price": 99.50
        }
      ]
    }
  ]
}
console.log(JSON.stringify(obj,null,1))


