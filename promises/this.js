const user = {
    name: "Alice",
    sayHi() {
      console.log("Hi, I'm", this.name);
    }
  };
  
  user.sayHi(); // "Hi, I'm Alice"
  

  const car = {
    brand: "Tesla",
    showBrand() {
      console.log(this.brand);
    }
  };
  
  const show = car.showBrand.bind(car);
  let ; show2  = car.showBrand
  show2 = show2.bind(car)
  show();
  show2()
  