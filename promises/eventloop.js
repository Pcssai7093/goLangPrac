// setTimeout(() => {
//     console.log('⏰ setTimeout');
  
//     process.nextTick(() => {
//       console.log('🌀 nextTick inside setTimeout');
//     });
  
//     Promise.resolve().then(() => {
//       console.log('💎 promise inside setTimeout');
//     });
//   }, 1000);


//   setTimeout(() => {
//     console.log('⏰ setTimeout2');
  
//     process.nextTick(() => {
//       console.log('🌀 nextTick inside setTimeout 2');
//     });
  
//     Promise.resolve().then(() => {
//       console.log('💎 promise inside setTimeout 2');
//     });
//   }, 500);
  
//   Promise.resolve().then(() => {
//     console.log('💎 promise');
//   });

//   process.nextTick(() => {
//     console.log('🌀 nextTick');
//   });
  
  

  let firstTask = new Promise(function(resolve, reject) {
    setTimeout(resolve, 500, 'Task One');
  });
  
  let secondTask;
  
  let thirdTask = new Promise(function(resolve, reject) {
    setTimeout(resolve, 1200, 'Task Three');
  });
  
  let fourthTask = new Promise(function(resolve, reject) {
    setTimeout(reject, 300, 'Task Four');
  });
  
  let fifthTask = new Promise(function(resolve, reject) {
    setTimeout(resolve, 1000, 'Task Two');
  });
  
  let combinedPromise = Promise.allSettled([firstTask, secondTask, thirdTask, fourthTask, fifthTask]);
  
  combinedPromise
    .then(function(data) {
      data.forEach(function(value) {
        console.log('Result:', value);
      });
    })
    .catch(function(error) {
      console.error('Error:', error);
    });



