import './appenv';
const fs = require('fs');
const INPUT = fs.readFileSync('./day09/input01.txt', 'utf-8').split('\n');


let tripBook = {};

INPUT.forEach(row => {

    let rowArray = row.split(" = ");
    let distance = parseInt(rowArray[1], 10);
    let startingPlace = rowArray[0].split(" to ")[0];
    let destination = rowArray[0].split(" to ")[1];

    if (!tripBook[startingPlace]){
        tripBook[startingPlace] = {};
    }
    if (!tripBook[destination]){
        tripBook[destination] = {};
    }

    

    tripBook.startingPlace.destination = distance
    tripBook.destination.startingPlace = distance

});