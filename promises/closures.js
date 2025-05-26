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



