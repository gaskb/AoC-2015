
let input = [1,1,1,3,1,2,2,1,1,3];


//1st part
let iterations = 40;
//2nd part
//let iterations = 50;
for (let iteration = 0; iteration < iterations; iteration++){
    console.log("Start " , iteration, " iteration");
    let intermediate = read(input);
    input = write(intermediate);
    console.log(iteration, " iteration end");
}

console.log("result = ", input.length);


function read(input) {
    let elementReadings = [];
    for (let i = 0; i < input.length; i++) {
    
        let element = input[i];
    
        // console.log(elementReadings);
        // console.log(element);
        
        if(i==0){
            elementReadings.push({ element:element , count:1})
        }
        else{
            let lastElement = elementReadings[elementReadings.length - 1];
    
            if (lastElement.element === element) {
                lastElement.count++;
            }
            else{
                elementReadings.push({element:element, count:1})
            }
        }
      }
    
      console.log(elementReadings);
      return elementReadings;
}

function write(input) {
    let result = []
    for (let i = 0; i < input.length; i++) {
    
        let element = input[i];

        result.push(element.count);
        result.push(element.element);
        // for(let counter = 0; counter < element.count; counter++){
            
        // }
        
      }
    
      console.log(result);
      return result;
}
