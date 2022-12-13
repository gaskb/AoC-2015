import './appenv';
const fs = require('fs');
const INPUT = fs.readFileSync('./day08/input01.txt', 'utf-8').split('\n');
//const INPUT = fs.readFileSync('./day08/test01.txt', 'utf-8').split('\n');

const countChars = (row) => {
    return row.length;
};

const countEncodedChars = (row) => {
    // console.log("");
    // console.log(row);

    let normalChars = row;

    
    normalChars = normalChars.replaceAll(/\\/g,"\\\\");
    normalChars = normalChars.replaceAll(/"/g,"\\\"");
    
    normalChars = '"' + normalChars + '"';
    
    console.log(normalChars)

    return normalChars.length;
}

const ciao = () =>{
    console.log("Hello world")
}

let result = {}

let grandStarting = 0;
let grandEncoded = 0;

let counter = 0;

INPUT.forEach(row => {
    console.log(" ")

    let startingChar = countChars(row);
    let encodedChars = countEncodedChars(row)

    console.log("encodedCharsLength", encodedChars)

    result[counter] = {}
    result[counter]["totstartingCharal"] = startingChar;
    result[counter]["normal"] = encodedChars;


    grandEncoded += encodedChars;
    grandStarting += startingChar

    counter ++;

});

console.log("grandStarting = ", grandStarting)
console.log("grandEncoded = ", grandEncoded)
console.log("");
console.log ("my result = ", grandEncoded - grandStarting)

//console.log(result)


