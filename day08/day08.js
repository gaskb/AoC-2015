import './appenv';
const fs = require('fs');
const INPUT = fs.readFileSync('./day08/input01.txt', 'utf-8').split('\n');
//const INPUT = fs.readFileSync('./day08/test01.txt', 'utf-8').split('\n');




const result2 = INPUT.reduce((acc, line) => acc + (line.length - eval(line).length), 0);


const start = /^"/g;
const end = /"$/g;
const doubleSlash = /\\\\/g;
const escapeQuote = /\\"/g;
const encodedChar = /\\[xX][0-9a-fA-F][0-9a-fA-F]/g;

const countChars = (row) => {
    return row.length;
};

const countNormalChars = (row) => {
    // console.log("");
    // console.log(row);

    let normalChars = row;

    normalChars = normalChars.replaceAll(start,"");
    normalChars = normalChars.replaceAll(end,"");
    normalChars = normalChars.replaceAll(doubleSlash,"\\");
    normalChars = normalChars.replaceAll(escapeQuote,'"');
    normalChars = normalChars.replaceAll(encodedChar," ");

    //console.log("normalChars = ", normalChars, " -> ", normalChars.length)
    console.log(normalChars)

    return normalChars.length;
}

const ciao = () =>{
    console.log("Hello world")
}

let result = {}

let grandTotal = 0;
let grandSpecial = 0;
let grandNormal = 0;

let counter = 0;
INPUT.forEach(row => {
    console.log(" ")

    let totalChar = countChars(row);
    let normalChars = countNormalChars(row)

    console.log(eval(row), " - eval(row)")
    console.log("normalCharsLength", normalChars)
    console.log("eval(row).length = ", eval(row).length)

    result[counter] = {}
    result[counter]["total"] = totalChar;
    result[counter]["normal"] = normalChars;


    grandNormal += normalChars;
    grandTotal += totalChar

    counter ++;

});

console.log("grandTotal = ", grandTotal)
console.log("grandNormal = ", grandNormal)
console.log("");
console.log ("my result = ", grandTotal - grandNormal)

console.log("1224 is too low")
console.log("1446 is too high")

console.log(result2);

//console.log(result)


